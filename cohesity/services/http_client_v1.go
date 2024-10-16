package services

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	apiClientV1 "github.com/cohesity/go-sdk/v1/client"
	"github.com/cohesity/go-sdk/v1/client/access_tokens"
	"github.com/cohesity/go-sdk/v1/client/principals"
	"github.com/cohesity/go-sdk/v1/client/protection_sources"
	modelsV1 "github.com/cohesity/go-sdk/v1/models"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/terraform-providers/terraform-provider-cohesity/cohesity/utils"
)

type TokenResponse struct {
	Token string `json:"accessToken"`
}
type CohesityClientV1 struct {
	client      *apiClientV1.CohesityInternalRESTAPI
	bearerToken runtime.ClientAuthInfoWriter
}

// initHttpClientInstance Creates a http client instance.
func initHttpClientInstance() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
}

func NewCohesityClientV1(config utils.Config) (*CohesityClientV1, error) {
	client := GetClientV1(config.ClusterVIP)
	bearerToken, err := GetAccessToken(config.ClusterVIP, config.ClusterUsername, config.ClusterPassword, config.ClusterDomain)
	if err != nil {
		return nil, err
	}
	return &CohesityClientV1{
		client:      client,
		bearerToken: httptransport.BearerToken(bearerToken),
	}, nil
}
func GetClientV1(clusterVip string) *apiClientV1.CohesityInternalRESTAPI {
	return apiClientV1.New(httptransport.NewWithClient(clusterVip, "/v2", nil, initHttpClientInstance()), strfmt.Default)
}
func GetAccessTokenV1(client *apiClientV1.CohesityInternalRESTAPI, username, password, domain string) (runtime.ClientAuthInfoWriter, error) {
	body := &modelsV1.AccessTokenCredential{
		Username: &username,
		Password: &password,
		Domain:   &domain,
	}
	accessTokenResponse, err := client.AccessTokens.GenerateAccessToken(access_tokens.NewGenerateAccessTokenParams().WithBody(body), nil)
	if err != nil {
		return nil, err
	}
	bearerToken := httptransport.BearerToken(*accessTokenResponse.Payload.AccessToken)
	return bearerToken, nil
}

// GetAccessToken generates an API access token for the provided user.
// Using this so as to ensure we don't get asked to change password.
func GetAccessToken(clusterVip, username, password, domain string) (string, error) {
	url := fmt.Sprintf("https://%s/irisservices/api/v1/public/accessTokens", clusterVip)

	// Request body.
	requestBody, err := json.Marshal(map[string]interface{}{
		"username":                username,
		"password":                password,
		"domain":                  domain,
		"skipForcePasswordChange": true,
	})
	if err != nil {
		log.Printf("failed to marshal json :%v", err)
		return "", err
	}

	// Create a custom HTTP client with insecure TLS config
	client := initHttpClientInstance()

	// Make the API call.
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Printf("failed to make http call. error %s", err.Error())
		return "", err
	}
	defer resp.Body.Close()

	// Parse the http response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("err:%v", err)
		return "", err
	}

	var tokenResponse TokenResponse
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		fmt.Printf("failed to unmarshal response into struct:%v", err)
		return "", err
	}

	return tokenResponse.Token, nil
}

// UpdateLinuxPassword sends a PUT request to update the Linux password.
func (c *CohesityClientV1) UpdateLinuxPassword(supportPassword string) error {
	body := &modelsV1.UpdateLinuxPasswordReqParams{
		LinuxUsername: toStrPtr("support"),
		LinuxPassword: &supportPassword,
	}
	_, err := c.client.Principals.UpdateLinuxCredentials(principals.NewUpdateLinuxCredentialsParams().WithBody(body), c.bearerToken)
	if err != nil {
		return err
	}
	return nil
}

// EnableSudoAccess sends a PUT request to enable sudo access.
func (c *CohesityClientV1) EnableSudoAccess() error {
	enable := true
	body := &modelsV1.LinuxSupportUserSudoAccessReqParams{
		SudoAccessEnable: &enable,
	}
	_, err := c.client.Principals.LinuxSupportUserSudoAccess(principals.NewLinuxSupportUserSudoAccessParams().WithBody(body), c.bearerToken)
	if err != nil {
		return err
	}
	return nil
}
func (c *CohesityClientV1) DeleteProtectionSource(id string) error {
	// Convert string to int64
	idInt64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid id: %v", err)
	}

	_, err = c.client.ProtectionSources.UnregisterProtectionSource(protection_sources.NewUnregisterProtectionSourceParams().WithID(idInt64), c.bearerToken)
	return err
}
