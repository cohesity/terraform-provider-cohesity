package services

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/terraform-providers/terraform-provider-cohesity/cohesity/utils"
)

type TokenResponse struct {
	Token string `json:"accessToken"`
}

// GetAccessToken generates an API access token for the provided user.
func GetAccessToken(config utils.Config) (string, error) {
	url := fmt.Sprintf("https://%s/irisservices/api/v1/public/accessTokens", config.ClusterVIP)

	// Request body.
	requestBody, err := json.Marshal(map[string]interface{}{
		"username":                config.ClusterUsername,
		"password":                config.ClusterPassword,
		"domain":                  config.ClusterDomain,
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
		return "", fmt.Errorf("failed to make http call. error %s", err.Error())
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
		return "", fmt.Errorf("failed to unmarshal response into struct:%v", err)

	}

	return tokenResponse.Token, nil
}

// UpdateLinuxPassword sends a PUT request to update the Linux password.
func UpdateLinuxPassword(config utils.Config, token string) error {
	url := fmt.Sprintf("https://%s/irisservices/api/v1/public/users/linuxPassword", config.ClusterVIP)
	requestBody := fmt.Sprintf("{\"linuxUsername\" : \"support\", \"linuxPassword\" : \"%s\"}", config.SupportPassword)

	return sendPutRequest(url, token, requestBody)
}

// EnableSudoAccess sends a PUT request to enable sudo access.
func EnableSudoAccess(target, token string) error {
	url := fmt.Sprintf("https://%s/irisservices/api/v1/public/users/linuxSupportUserSudoAccess", target)
	requestBody := `{"sudoAccessEnable" : true}`

	return sendPutRequest(url, token, requestBody)
}

// sendPutRequest sends a PUT request to the specified URL with the given token and request body.
func sendPutRequest(url, token, requestBody string) error {
	// Create a custom HTTP client with insecure TLS config
	client := initHttpClientInstance()
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer([]byte(requestBody)))
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("error: %s", string(body))
	}
	fmt.Println(string(body))
	fmt.Printf("\nStatus Code: %d\n", resp.StatusCode)
	fmt.Printf("Content-Type: %s\n", resp.Header.Get("Content-Type"))
	return nil
}

// initHttpClientInstance Creates a http client instance.
func initHttpClientInstance() *http.Client {
	return &http.Client{
		Timeout: 0,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
}
