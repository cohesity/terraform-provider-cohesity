package services

import (
	"fmt"
	"net/url"
	"strings"

	apiClientV2 "github.com/cohesity/go-sdk/v2/client"
	"github.com/cohesity/go-sdk/v2/client/access_token"
	"github.com/cohesity/go-sdk/v2/client/platform"
	modelsV2 "github.com/cohesity/go-sdk/v2/models"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/terraform-providers/terraform-provider-cohesity/cohesity/utils"
)



type CohesityClientV2 struct {
	client      *apiClientV2.CohesityRESTAPI
	bearerToken runtime.ClientAuthInfoWriter
}

func NewCohesityClientV2(config utils.Config) (*CohesityClientV2, error) {
	client := GetClientV2(config.ClusterVIP)
	bearerTokenString, err := GetAccessToken(config)
	if err != nil {
		return nil, err
	}
	return &CohesityClientV2{
		client:      client,
		bearerToken: httptransport.BearerToken(bearerTokenString),
	}, nil
}
func GetClientV2(clusterVip string) *apiClientV2.CohesityRESTAPI {
	return apiClientV2.New(httptransport.NewWithClient(clusterVip, "/v2", nil, initHttpClientInstance()), strfmt.Default)
}
func (c *CohesityClientV2) GetClusterInfo() (*modelsV2.Cluster,error) {
	resp, err := c.client.Platform.GetCluster(platform.NewGetClusterParams(),c.bearerToken)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}
func (c *CohesityClientV2) UpdateClusterInfo(toUpdateClusterInfo *modelsV2.Cluster) (*modelsV2.Cluster,error) {
	resp, err := c.client.Platform.UpdateCluster(platform.NewUpdateClusterParams().WithBody(toUpdateClusterInfo),c.bearerToken)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
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


// NormalizeHTTPSAddress ensures the VIP has an https:// prefix and proper formatting
func NormalizeHTTPSAddress(vip string) (string, error) {
	original := vip

	// If scheme is missing, prepend "//" to parse host and port correctly
	if !strings.Contains(vip, "://") {
		vip = "//" + vip
	}

	parsed, err := url.Parse(vip)
	if err != nil {
		return "", fmt.Errorf("failed to parse VIP %q: %w", original, err)
	}

	host := parsed.Hostname()
	port := parsed.Port()

	// Default port is 443 for https
	address := host
	if port != "" && port != "443" {
		address = fmt.Sprintf("%s:%s", host, port)
	}

	return fmt.Sprintf("https://%s", address), nil
}