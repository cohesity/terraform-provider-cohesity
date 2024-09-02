package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	v2 "github.com/terraform-providers/terraform-provider-cohesity/cohesity/models/v2"
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
