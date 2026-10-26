package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func convertStringToASCIISlice(str string) []int64 {
	result := make([]int64, len(str))
	for i, char := range str {
		result[i] = int64(uint8(char))
	}
	return result
}
func normalizePrivateKey(clientPrivateKey string) string {
	// Replace all encoded equal signs with actual '='
	clientPrivateKey = strings.ReplaceAll(clientPrivateKey, "\\u003d", "=")
	
	// Replace all string representations of '\n' with actual new lines
	clientPrivateKey = strings.ReplaceAll(clientPrivateKey, "\\n", "\n")
	
	return clientPrivateKey
}
func convertASCIISliceToString(asciiSlice []uint8) string {
	return string(asciiSlice)
}
type RegisterEntityResult struct {

	// If the entity was successfully registered, it's EntityProto is returned to
	// the user.
	// Note: This might not be required for magneto v2, but will be populated
	// none-the-less for backwards compatibility reasons.
	Entity PrivateEntityProto `json:"entity,omitempty"`
}
type EntityArgWrapper struct {
	Entity PrivateEntityProto `json:"entity,omitempty"`
	EntityInfo EntityAccessInfo `json:"entityInfo,omitempty"`

}
type EntityAccessInfo struct {

	// Credentials that will be used to login to the environment.
	Credentials PrivateCredentials `json:"credentials,omitempty"`


	// The type of environment behind the endpoint.
	Type int32 `json:"type,omitempty"`

}
type PrivateCredentials struct {
	CloudCredentials NexusCloudCredentials `json:"cloudCredentials,omitempty"`
}
type NexusCloudCredentials struct {
	GcpCredentials PrivateGcpCredentials `json:"gcpCredentials,omitempty"`
}
type PrivateGcpCredentials struct {

	// Client email address associated with the service account.
	ClientEmailAddress string `json:"clientEmailAddress,omitempty"`


	// Client private key generated at the time of creating the service
	// account. This is used for authenticating ourselves with Google cloud.
	// This field only stores the private key in bytes and doesn't encrypt it.
	EncryptedClientPrivateKey []int64 `json:"encryptedClientPrivateKey"`

	// Id of the project associated with Google cloud account.
	ProjectID string `json:"projectId,omitempty"`

}
type EntityHierarchyWrapper struct {

	// EntityHierarchy is the protobuf for the registered entity along with its
	// children hierarchy.
	EntityHierarchy *EntityHierarchyProto `json:"entityHierarchy,omitempty"`
}
type EntityHierarchyProto struct {
	Entity PrivateEntityProto `json:"entity,omitempty"`
}
type PrivateEntityProto struct {
	GcpEntity PrivateGcpEntity `json:"gcpEntity,omitempty"`
	ID int64 `json:"id,omitempty"`
	Type int32 `json:"type,omitempty"`
}
type PrivateGcpEntity struct {

	// Populated in entities of type kSubnet if the subnet is part of a Shared
	// VPC. This contains the ID of the host project the subnet belongs to.
	// Populated in entities of type kProject if the project is a service project
	// in a Shared VPC setup. This contains the ID of the host project it is
	// attached to.
	HostProjectID string `json:"hostProjectId,omitempty"`

	// The host environment type. This is set for entities of type
	// kVirtualMachine.
	HostType int32 `json:"hostType,omitempty"`

	// The client email of GCP service account is used as owner id.
	OwnerID string `json:"ownerId,omitempty"`

	// Private ip address. This is set for entities of type kVirtualMachine.
	PrivateIPAddress string `json:"privateIpAddress,omitempty"`

	// For entities of type kVirtualMachine this contains the id of the project
	// the virtual machine belongs to. For entities of type kSubnet, this
	// contains the id of project the subnet is available to. For entities of the
	// type kVPCConnector this contains the id of the project the VPC connector
	// belongs to.
	ProjectID string `json:"projectId,omitempty"`

	Id *int64 `json:"id,omitempty"`

	// For entities of type kVirtualMachine this contains the region
	// the virtual machine belongs to.
	// For the kVPCConnectorEntity this contains the region the VPC connector
	// belongs to.
	Region string `json:"region,omitempty"`

	// The type of entity this proto refers to.
	Type *int32 `json:"type,omitempty"`

	// In the kSubnet entity this field contains the VPC network the subnet
	// belongs to.
	// In the kSubnet, kVirtualMachine and kVPCConnector entities this field
	// contains the name of VPC network the entity resource belongs to.
	VpcNetwork string `json:"vpcNetwork,omitempty"`

	// Subnetwork to be used for deploying proxy VMs. Set only for the service
	// account entity
	VpcSubnetwork string `json:"vpcSubnetwork,omitempty"`

	// For entities of type kVirtualMachine this contains the zone the
	// virtual machine belongs to.
	Zone string `json:"zone,omitempty"`

	CommonInfo EntityCommonInfo `json:"commonInfo,omitempty"`

}
type EntityCommonInfo struct {
	ID string `json:"id,omitempty"`
}
func AddGCPProtectionSource(clusterVip, bearerToken, clientEmailAddress, clientPrivateKey, projectId, VPCNetwork, VPCSubnetwork string) (int64, error) {
	url := fmt.Sprintf("https://%s/irisservices/api/v1/backupsources", clusterVip)
	gcpType := int32(0)
	// Request body structure.
	requestBody:= EntityArgWrapper{
		Entity: PrivateEntityProto{
			Type: int32(20),
			GcpEntity: PrivateGcpEntity{
				CommonInfo: EntityCommonInfo{
					ID: clientEmailAddress,
				},
				Type: &gcpType,
				OwnerID: clientEmailAddress,
				ProjectID: projectId,
				VpcNetwork: VPCNetwork,
				VpcSubnetwork: VPCSubnetwork,
			},
		},
		EntityInfo: EntityAccessInfo{
			Type: int32(20),
			Credentials: PrivateCredentials{
				CloudCredentials: NexusCloudCredentials{
					GcpCredentials: PrivateGcpCredentials{
						ClientEmailAddress: clientEmailAddress,
						ProjectID: projectId,
						EncryptedClientPrivateKey: convertStringToASCIISlice(normalizePrivateKey(clientPrivateKey)),
					},
				},
			},
		},
	}
	

	// Marshal request body to JSON.
	requestBodyJSON, err := json.Marshal(requestBody)
	if err != nil {
		log.Printf("failed to marshal request body: %v", err)
		return -1, err
	}

	// Create a custom HTTP client with insecure TLS config.
	client := initHttpClientInstance()

	// Create an HTTP request.
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestBodyJSON))
	if err != nil {
		log.Printf("failed to create http request: %v", err)
		return -1, err
	}

	// Add headers.
	req.Header.Set("Authorization", "Bearer "+ bearerToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Make the API call.
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("failed to make http request: %v", err)
		return -1, err
	}
	defer resp.Body.Close()
	// Check for HTTP errors.
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Printf("http call failed with status code %d: %s ", resp.StatusCode, string(body))
		return -1, fmt.Errorf("failed to add GCP protection source: %s\nstatuscode: %v \n\n \n body: %v", string(body),resp.StatusCode,string(requestBodyJSON))
	}

	// Parse the response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("failed to read response body: %v", err)
		return -1, err
	}

	// Define response structure.
	var response RegisterEntityResult


	// Unmarshal the response body.
	err = json.Unmarshal(body, &response)
	if err != nil {
		// log.Printf("[INFO] failed to unmarshal response:%v %v,\n response:%v",resp.StatusCode, err, body)
		return -1, fmt.Errorf("failed to unmarshal response:%v", err)
	}

	return response.Entity.ID, nil
}


func GetGcpProtectionSource(clusterVip, bearerToken, id string) (*EntityHierarchyProto, error) {
	// Convert string to int64 to validate the ID
	_, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %v", err)
	}

	url := fmt.Sprintf("https://%s/irisservices/api/v1/backupsources?entityId=%s&onlyReturnOneLevel=true", clusterVip, id)

	// Create a custom HTTP client with insecure TLS config.
	client := initHttpClientInstance()

	// Create an HTTP request.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("failed to create http request: %v", err)
		return nil, err
	}

	// Add headers.
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bearerToken))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Make the API call.
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("failed to make http request: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Check for HTTP errors.
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Printf("http call failed with status code %d: %s", resp.StatusCode, string(body))
		return nil, fmt.Errorf("failed to get GCP protection source: %s", string(body))
	}
	log.Printf("[INFO]status: %v", resp.StatusCode)
	// Parse the response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("failed to read response body: %v", err)
		return nil, err
	}

	// Unmarshal the response body into a generic map.
	var response EntityHierarchyWrapper
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Printf("failed to unmarshal response: %d ,: %v",resp.StatusCode, err)
		return nil, err
	}

	return response.EntityHierarchy, nil
}

func DeleteProtectionSource(clusterVip, bearerToken, id string) error {
	// Convert string to int64 to validate the ID
	_, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid id: %v", err)
	}

	// Construct the URL.
	url := fmt.Sprintf("https://%s/irisservices/api/v1/protectionsources/%s", clusterVip, id)

	// Create a custom HTTP client with insecure TLS config.
	client := initHttpClientInstance()

	// Create an HTTP DELETE request.
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Printf("failed to create http request: %v", err)
		return err
	}

	// Add headers.
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bearerToken))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")


	// Make the API call.
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("failed to make http request: %v", err)
		return err
	}
	defer resp.Body.Close()

	// Check for HTTP errors.
	if resp.StatusCode != http.StatusNoContent {
		body, _ := io.ReadAll(resp.Body)
		log.Printf("http call failed with status code %d: %s", resp.StatusCode, string(body))
		return fmt.Errorf("failed to delete protection source: %s", string(body))
	}

	return nil
}

func UpdateGcpProtectionSource(clusterVip, bearerToken, id,clientEmailAddress, clientPrivateKey, projectId, VPCNetwork, VPCSubnetwork string) error {
	url := fmt.Sprintf("https://%s/irisservices/api/v1/backupsources/%s", clusterVip, id)
	gcpType := int32(0)
	entityId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid id: %v", err)
	}
	// Request body structure.
	requestBody:= EntityArgWrapper{
		Entity: PrivateEntityProto{
			Type: int32(20),
			ID: entityId,
			GcpEntity: PrivateGcpEntity{
				CommonInfo: EntityCommonInfo{
					ID: clientEmailAddress,
				},
				Type: &gcpType,
				OwnerID: clientEmailAddress,
				ProjectID: projectId,
				VpcNetwork: VPCNetwork,
				VpcSubnetwork: VPCSubnetwork,
			},
		},
		EntityInfo: EntityAccessInfo{
			Type: int32(20),
			Credentials: PrivateCredentials{
				CloudCredentials: NexusCloudCredentials{
					GcpCredentials: PrivateGcpCredentials{
						ClientEmailAddress: clientEmailAddress,
						ProjectID: projectId,
						EncryptedClientPrivateKey: convertStringToASCIISlice(normalizePrivateKey(clientPrivateKey)),
					},
				},
			},
		},
	}
	

	// Marshal request body to JSON.
	requestBodyJSON, err := json.Marshal(requestBody)
	if err != nil {
		log.Printf("[ERROR] failed to marshal request body: %v", err)
		return err
	}

	// Create a custom HTTP client with insecure TLS config.
	client := initHttpClientInstance()

	// Create an HTTP request.
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(requestBodyJSON))
	if err != nil {
		log.Printf("[ERROR] failed to create http request: %v", err)
		return err
	}

	// Add headers.
	req.Header.Set("Authorization", "Bearer "+ bearerToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Make the API call.
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("[ERROR] failed to make http request: %v", err)
		return err
	}
	defer resp.Body.Close()
	// Check for HTTP errors.
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Printf("[ERROR] http call failed with status code %d: %s ", resp.StatusCode, string(body))
		return fmt.Errorf("failed to add GCP protection source: %s\nstatuscode: %v \n\n \n body: %v", string(body),resp.StatusCode,string(requestBodyJSON))
	}


	return nil
}
