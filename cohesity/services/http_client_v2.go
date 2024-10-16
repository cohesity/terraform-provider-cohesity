package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	apiClientV2 "github.com/cohesity/go-sdk/v2/client"
	"github.com/cohesity/go-sdk/v2/client/access_token"
	modelsV2 "github.com/cohesity/go-sdk/v2/models"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	v2 "github.com/terraform-providers/terraform-provider-cohesity/cohesity/models/v2"
	"github.com/terraform-providers/terraform-provider-cohesity/cohesity/utils"
)

// GetClusterInfo calls the v2 /clusters API and returns the API response.
func GetClusterInfo(clusterVip, accessToken string) (*v2.Cluster, error) {
	url := fmt.Sprintf("https://%s/v2/clusters", clusterVip)

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// Set the Authorization header
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("Accept", "application/json")

	// Make the GET request using the default HTTP client
	client := initHttpClientInstance()
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make GET request: %v", err)
	}
	defer resp.Body.Close()

	/// Check if the status code is 200 OK
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response code: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// Unmarshal the JSON into the Cluster struct
	var cluster v2.Cluster
	err = json.Unmarshal(body, &cluster)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON response: %v", err)
	}

	return &cluster, nil
}

// UpdateClusterInfo calls the PUT v2 /clusters API and to update the cluster information.
func UpdateClusterInfo(clusterVip string, toUpdateClusterInfo v2.Cluster, accessToken string) (*v2.Cluster, error) {
	url := fmt.Sprintf("https://%s/v2/clusters", clusterVip)

	// Request body.
	requestBody, err := json.Marshal(toUpdateClusterInfo)
	if err != nil {
		log.Printf("failed to marshal req body into json :%v", err)
		return nil, err
	}

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// Set the Authorization header
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("Accept", "application/json")

	// Make the GET request using the default HTTP client
	client := initHttpClientInstance()
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make GET request: %v", err)
	}
	defer resp.Body.Close()

	/// Check if the status code is 200 OK
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response code: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// Unmarshal the JSON into the Cluster struct
	var cluster v2.Cluster
	err = json.Unmarshal(body, &cluster)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON response: %v", err)
	}

	return &cluster, nil
}

type CohesityClientV2 struct {
	client      *apiClientV2.CohesityRESTAPI
	bearerToken runtime.ClientAuthInfoWriter
}

func NewCohesityClientV2(config utils.Config) (*CohesityClientV2, error) {
	client := GetClientV2(config.ClusterVIP)
	bearerToken, err := GetAccessTokenV2(client, config.ClusterUsername, config.ClusterPassword, config.ClusterDomain)
	if err != nil {
		return nil, err
	}
	return &CohesityClientV2{
		client:      client,
		bearerToken: bearerToken,
	}, nil
}
func GetClientV2(clusterVip string) *apiClientV2.CohesityRESTAPI {
	return apiClientV2.New(httptransport.NewWithClient(clusterVip, "/v2", nil, initHttpClientInstance()), strfmt.Default)
}

func GetAccessTokenV2(client *apiClientV2.CohesityRESTAPI, username, password, domain string) (runtime.ClientAuthInfoWriter, error) {
	body := &modelsV2.CreateAccessTokenRequestParams{
		Username: &username,
		Password: &password,
		Domain:   &domain,
	}
	accessTokenResponse, err := client.AccessToken.CreateAccessToken(access_token.NewCreateAccessTokenParams().WithBody(body), nil)
	if err != nil {
		return nil, err
	}
	bearerToken := httptransport.BearerToken(*accessTokenResponse.Payload.AccessToken)
	return bearerToken, nil
}

func toStrPtr(variable string) *string {
	return &variable
}
