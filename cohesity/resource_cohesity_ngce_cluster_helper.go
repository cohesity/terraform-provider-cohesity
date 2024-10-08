package cohesity

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"
	"unicode"

	"github.com/terraform-providers/terraform-provider-cohesity/cohesity/services"
	"github.com/terraform-providers/terraform-provider-cohesity/cohesity/utils"

	"golang.org/x/crypto/ssh"
)

const DefaultCohesityPswd = "Cohe$1ty"
const CreateClusterScriptLocation = "/home/cohesity/cluster_creation.sh"

//----------------------------------------------------------------------------------------

type NodeAPIResponse struct {
	Error   string `json:"Error"`
	Message string `json:"message"`
}

//----------------------------------------------------------------------------------------

// ClusterInfo Cluster status model.
type ClusterInfo struct {
	NodeStatus    []NodeStatus  `json:"nodeStatus"`
	ClusterID     int64         `json:"clusterId"`
	IncarnationID int64         `json:"incarnationId"`
	ClusterName   string        `json:"clusterName"`
	BulletinState BulletinState `json:"bulletinState"`
}

//----------------------------------------------------------------------------------------

type BulletinState struct {
	CurrentOperation string `json:"currentOperation"`
}

//----------------------------------------------------------------------------------------

type NodeStatus struct {
	NodeId           int64    `json:"nodeId"`
	NodeIps          []string `json:"nodeIps"`
	PartOfCluster    bool     `json:"partOfCluster"`
	MarkedForRemoval bool     `json:"markedForRemoval"`
}

//----------------------------------------------------------------------------------------

// NodeIPMap represents the structure of node IPs and IDs
type NodeIPMap struct {
	NodeID  int64    `json:"node_id"`
	NodeIPs []string `json:"node_ips"`
}

//----------------------------------------------------------------------------------------

/**
 * enableSupportUser attempts to update the Linux password and enable sudo access.
 * It retries up to 10 times at 30-second intervals and returns an error if unsuccessful.
 */
func enableSupportUser(clusterHostname, supportPassword string) error {
	for i := 0; i < 10; i++ {
		log.Printf("enabling support user access on the cluster %s", clusterHostname)

		// Sleep for 30 seconds before the next attempt
		if i > 0 { // No need to sleep for the first run
			fmt.Println("Waiting for 30 seconds...")
			time.Sleep(30 * time.Second)
		}

		// Get the token
		token, err := services.GetAccessToken(clusterHostname, "admin", "admin")
		if err != nil || token == "" {
			log.Println("Failed to retrieve token")
			continue
		}

		// Update Linux password
		err = services.UpdateLinuxPassword(clusterHostname, supportPassword, token)
		if err != nil {
			log.Printf("Failed to update Linux password: %v\n", err)
			continue
		} else {
			log.Println("Linux password updated successfully.")
		}

		// Enable sudo access
		err = services.EnableSudoAccess(clusterHostname, token)
		if err != nil {
			log.Printf("Failed to enable sudo access: %v\n", err)
			continue
		} else {
			log.Println("Sudo access enabled successfully.")
		}

		// If both operations succeed, return nil
		return nil

	}

	// If all retries fail, return an error
	return fmt.Errorf("enabling support user failed after 10 attempts")
}

//----------------------------------------------------------------------------------------

// execCreateClusterCmd executes command on a remote host to create cohesity cluster.
func execCreateClusterCmd(clusterParams CohesityClusterConfig) (string, error) {
	log.Printf("[INFO] Creating cluster with the config %+v", clusterParams)
	log.Printf("[INFO] connecting to node through hostname %s..", clusterParams.Hostname)

	// Create an SSH client
	client, err := services.NewSSHClient(clusterParams.Hostname, "cohesity", DefaultCohesityPswd, time.Hour)
	if err != nil {
		log.Printf("Failed to init SSH client instance.")
		return "", err
	}
	defer client.Close()

	log.Printf("[INFO] moving cluster create clusterCreationScript to one of the cluster node.")

	// Script content.
	clusterCreationScript := getClusterCreateScriptContent()
	// Upload the script onto the cluster.
	err = services.UploadScriptToVM(client, CreateClusterScriptLocation, clusterCreationScript)
	if err != nil {
		return "", fmt.Errorf("failed to upload clusterCreationScript to VM: %s", err)
	}
	log.Printf("[INFO] successfully moved the cluster creation script to %s", CreateClusterScriptLocation)

	// Execute the script.
	cmd := fmt.Sprintf("%s '%s' '%s' '%s' '%s' '%s' '%s' '%s' '%s' '%d' '%t' '%s' '%s'",
		CreateClusterScriptLocation,
		clusterParams.Name,
		strings.Join(clusterParams.NodeIps, ","),
		clusterParams.Hostname,
		clusterParams.SubnetGateway,
		clusterParams.SubnetMask,
		strings.Join(clusterParams.DnsServerIps, ","),
		clusterParams.NtpServers,
		clusterParams.DomainNames,
		clusterParams.MetadataFaultTolerance,
		clusterParams.AddApps,
		clusterParams.AppsSubnet,
		clusterParams.AppsSubnetMask)
	log.Printf("[INFO] Executing clusterCreationScript on VM: %s", cmd)
	output, err := services.ExecuteSSHCommand(client, cmd)
	if err != nil {
		log.Printf("Failed to execute the cluster create script. Error %s", err.Error())
		return "", err
	}

	return output, err
}

//----------------------------------------------------------------------------------------

// extractJSONFromCreateClusterCmdOutput function extracts JSON-like data from script output.
func extractJSONFromCreateClusterCmdOutput(output string) string {
	// Look for the line starting with "Parsing the final status output..."
	marker := "Parsing the final status output to extract Node IP information..."
	jsonStartIdx := strings.Index(output, marker)
	if jsonStartIdx == -1 {
		log.Fatalf("Failed to find the marker in the output")
	}

	// Extract the part of the string after the marker
	jsonOutput := output[jsonStartIdx+len(marker):]

	// Remove any non-JSON lines
	jsonOutput = strings.TrimSpace(jsonOutput)

	return jsonOutput
}

// getNodeIDByIP takes a slice of node_ip_map as []interface{} and returns node ID for a given IP
func getNodeIDByIP(nodeIPMapInterface []interface{}, searchIP string) (int, bool) {
	for _, item := range nodeIPMapInterface {
		nodeIPMapItem, ok := item.(map[string]interface{})
		if !ok {
			log.Fatalf("Failed to assert item to map[string]interface{}")
		}

		nodeID, _ := nodeIPMapItem["node_id"].(int)
		nodeIPsInterface, _ := nodeIPMapItem["node_ips"].([]interface{})

		for _, ip := range nodeIPsInterface {
			ipStr, _ := ip.(string)
			log.Printf("[INFO] searching for: %s checking: %s, nodeID: %d", searchIP, ipStr, nodeID)
			if ipStr == searchIP {
				return nodeID, true
			}
		}
	}
	return 0, false
}

//----------------------------------------------------------------------------------------

// execRemoveNodeCmd removes a node from the Cohesity cluster.
func execRemoveNodeCmd(user, password string, nodeID int, client *ssh.Client) error {
	log.Printf("[INFO] removing node id %d from the cluster", nodeID)

	// iris cli instruction.
	cmd := fmt.Sprintf("/home/support/bin/iris_cli --output=prettyjson --username=%s --password=%s"+
		" --skip_password_prompt=true --skip_force_password_change=true node rm id=%d; status=$?; sleep 30; exit"+
		" $status", user, password, nodeID)
	output, err := services.ExecuteSSHCommand(client, cmd)
	log.Printf("[INFO] cli output: %s", output)

	if err != nil {
		// If the execute instruction fails.
		output = strings.TrimSpace(output)

		var cliOutputResp NodeAPIResponse
		unmarshalErr := json.Unmarshal([]byte(output), &cliOutputResp)
		if unmarshalErr != nil {
			log.Printf("failed to unmarshal the output into API model. Error %s", unmarshalErr.Error())
			return unmarshalErr
		}

		// Define the regex pattern
		// It is possible that the node is down/deleted. In that case, fresh instruction needs to be exec based on
		// a new flag arg.
		// Compile the regex. If the node is already marked for removal, no action is required.
		re := regexp.MustCompile(`^Node having id: \d+ is marked for removal$`)
		if re.MatchString(cliOutputResp.Error) {
			log.Printf("[INFO] node is already marked for removal")
			return nil
		}

		// If not is not reachable...
		pattern := `^Node having id: \d+ is not reachable\. If the entity is down, specify is-offline=true option\.$`
		re = regexp.MustCompile(pattern)
		// Check if the string matches the pattern
		if re.MatchString(cliOutputResp.Error) {
			cmd := fmt.Sprintf("/home/support/bin/iris_cli --output=prettyjson --username=%s --password=%s"+
				" --skip_password_prompt=true --skip_force_password_change=true node rm id=%d is-offline=true"+
				"; status=$?; sleep 30; exit $status", user, password, nodeID)
			_, err = services.ExecuteSSHCommand(client, cmd)
			if err != nil {
				log.Printf("failed to remove node. error %s", err.Error())
				return err
			}

			log.Printf("[INFO] node is marked for removal")
			return nil
		}

		return fmt.Errorf("could not delete node. error: %s", cliOutputResp.Error)
	}

	// No errors. Node marked for removal.
	log.Printf("[INFO] node is marked for removal")
	return nil
}

//----------------------------------------------------------------------------------------

// execAddNodesCmd add nodes to a cluster.
func execAddNodesCmd(user, password, nodeIPsCsvString string, client *ssh.Client) error {
	log.Printf("adding node ips %s to the cluster", nodeIPsCsvString)

	var isCmdExecutedSuccessfully bool
	for i := 0; i < 50; i++ {
		log.Printf("attempting to add node to cluster. attempt: %d", i)
		isCmdExecutedSuccessfully = false

		// CLI cmd to join cloud cluster with new node ips.
		cmd := fmt.Sprintf("/home/support/bin/iris_cli --output=prettyjson --username=%s --password=%s"+
			" --skip_password_prompt=true --skip_force_password_change=true cluster cloud-join node-ips=%s; status=$"+
			"?; sleep 30; exit $status", user, password, nodeIPsCsvString)
		output, err := services.ExecuteSSHCommand(client, cmd)
		if err != nil {
			output = strings.TrimSpace(output)
			log.Printf("[INFO] attempt to add node to cluster failed: output %s, error %s", output, err.Error())

			// Map the error response into model.
			var outputResponse NodeAPIResponse
			marshalErr := json.Unmarshal([]byte(output), &outputResponse)
			if marshalErr != nil {
				log.Printf("failed to marshal output into response model. error %s", marshalErr)
				return marshalErr
			}
			log.Printf("[INFO] output response: %+v", outputResponse)

			// See if the error string suggests that the node ip is not reachable.
			re := regexp.MustCompile(`^Could not reach node IP (\d{1,3}\.){3}\d{1,3}: Request timed out\.$`)
			if re.MatchString(outputResponse.Error) {
				// Continue to try and add nodes.
				continue
			}

			return fmt.Errorf("cannot add node. error received :%s", outputResponse.Error)
		}

		// No error at this point.
		isCmdExecutedSuccessfully = true
		break
	}

	if isCmdExecutedSuccessfully {
		return waitUntilNodeAddition(user, password, client, nodeIPsCsvString)
	}

	return fmt.Errorf("failed to add node into cluster after max attempts")
}

//----------------------------------------------------------------------------------------

// waitUntilNodeAddition wait for addition of a node into the cluster.
func waitUntilNodeAddition(user, password string, client *ssh.Client, nodeIPsString string) error {
	// Wait for max 30 minx.
	for i := 0; i < 30; i++ {
		log.Printf("verifying if node addition was successful. attempt count %d", i)
		// Sleep for a minute.
		time.Sleep(time.Minute)

		clusterStatus, err := fetchClusterStatus(user, password, client)
		if err != nil {
			log.Printf("failed to fetch the current cluster status. %s", err.Error())
			return err
		}

		nodeIps := strings.Split(nodeIPsString, ",")
		if !checkIfNodesAreAddedToCluster(nodeIps, clusterStatus) {
			continue
		} else {

			return nil
		}
	}

	return fmt.Errorf("failed to determine node addition was successful after max attempts")
}

//----------------------------------------------------------------------------------------

// checkIfNodesAreAddedToCluster checks if all the input node ids are part of the cluster.
func checkIfNodesAreAddedToCluster(nodeIps []string, clusterStatus *ClusterInfo) bool {
	// Fetch the node information.
	_, nodeIpsInCluster, _ := getNodeInformationFromClusterStatus(*clusterStatus)

	// Check if the added node is present and part of the cluster.
	for _, nodeIp := range nodeIps {
		if !isInArray(nodeIp, nodeIpsInCluster) {
			log.Printf("Node Ip %s not part of the cluster", nodeIp)
			return false
		}
	}
	return true
}

//----------------------------------------------------------------------------------------

// isInArray checks if an array element is present.
func isInArray(ele string, haystack []string) bool {
	for _, a := range haystack {
		if a == ele {
			return true
		}
	}

	return false
}

//----------------------------------------------------------------------------------------

// fetchClusterStatus fetches the current cluster status.
func fetchClusterStatus(user, password string, client *ssh.Client) (*ClusterInfo, error) {
	log.Printf("[INFO] password: %s", password)
	cmd := fmt.Sprintf("/home/support/bin/iris_cli --output=prettyjson --username=%s --password=%s"+
		" --skip_password_prompt=true --skip_force_password_change=true cluster status && sleep 30", user, password)
	output, err := services.ExecuteSSHCommand(client, cmd)
	if err != nil {
		log.Printf("failed to exec cluster status cmd. error %s", err.Error())
		return nil, err
	}

	return parseClusterStatusJSONResponse(strings.TrimSpace(output))
}

//----------------------------------------------------------------------------------------

// parseClusterStatusJSONResponse function parses the cluster status JSON output into ClusterInfo struct.
func parseClusterStatusJSONResponse(jsonOutput string) (*ClusterInfo, error) {
	var clusterStatus ClusterInfo

	// Unmarshal the JSON data
	err := json.Unmarshal([]byte(jsonOutput), &clusterStatus)
	if err != nil {
		log.Printf("error unmarshalling cluster status json into struct %v", err)
		return nil, fmt.Errorf("error unmarshalling cluster status json into struct: %v", err)
	}

	return &clusterStatus, nil
}

//----------------------------------------------------------------------------------------

// getNodeInformationFromClusterStatus analyzes the clusterInfo and returns the node ips and other information
func getNodeInformationFromClusterStatus(clusterStatus ClusterInfo) ([]map[string]interface{}, []string, bool) {
	// Is true if all nodes are marked as part of the cluster.
	areAllNodesPartOfCluster := true
	var NodeIpsMapList []map[string]interface{}

	var allActiveNodeIps []string // List of all node ips.
	// Create a mapping of node IDs to IPs
	nodeIdToIPsMap := make(map[int64][]string)

	// Iterate over the nodes data and build map of nodeId -> Node Ip
	for _, node := range clusterStatus.NodeStatus {
		// Check if the node is part of cluster and not marked for removal.
		if len(node.NodeIps) > 0 && !node.MarkedForRemoval && node.PartOfCluster {
			nodeIdToIPsMap[node.NodeId] = node.NodeIps
			allActiveNodeIps = append(allActiveNodeIps, node.NodeIps...) // Append all IP addresses to the slice
		}

		if !node.PartOfCluster {
			areAllNodesPartOfCluster = false
		}
	}

	// Prepare the nodeIpMap.
	for nodeID, nodeIPs := range nodeIdToIPsMap {
		nodeIpsMap := map[string]interface{}{
			"node_id":  nodeID,
			"node_ips": nodeIPs,
		}
		log.Printf("[INFO] %d, %s", nodeID, nodeIPs)
		NodeIpsMapList = append(NodeIpsMapList, nodeIpsMap)
	}

	log.Printf("[INFO] allActiveNodeIps %v", allActiveNodeIps)
	log.Printf("[INFO] areAllNodesPartOfCluster %v", areAllNodesPartOfCluster)
	log.Printf("[INFO] NodeIpsMapList %v", NodeIpsMapList)

	return NodeIpsMapList, allActiveNodeIps, areAllNodesPartOfCluster
}

//----------------------------------------------------------------------------------------

// Returns the cluster creation script content.
func getClusterCreateScriptContent() string {
	clusterCreationScript := `#!/bin/bash

				# Variables
				CLUSTER_NAME="$1"
				NODE_IPS="$2"
				HOSTNAME="$3"
				SUBNET_GATEWAY="$4"
				SUBNET_MASK="$5"
				DNS_SERVER_IPS="$6"
				NTP_SERVERS="$7"
				DOMAIN_NAMES="$8"
				METADATA_FAULT_TOLERANCE="$9"
				ADD_APPS="${10}"
				APPS_SUBNET="${11}"
				APPS_SUBNET_MASK="${12}"

				# Initial cluster creation command
				cmd="iris_cli --output=prettyjson --username=admin --password=admin --skip_password_prompt=true --skip_force_password_change=true cluster cloud-create cluster-size=nextGen enable-software-encryption=true name=${CLUSTER_NAME} node-ips=${NODE_IPS} hostname=${HOSTNAME} subnet-gateway=${SUBNET_GATEWAY} subnet-mask=${SUBNET_MASK} dns-server-ips=${DNS_SERVER_IPS} ntp-servers=${NTP_SERVERS} domain-names=${DOMAIN_NAMES} metadata-fault-tolerance=${METADATA_FAULT_TOLERANCE}"
				if [ "${ADD_APPS}" = "true" ]; then
					cmd="${cmd} apps-subnet=${APPS_SUBNET} apps-subnet-mask=${APPS_SUBNET_MASK}"
				fi
				echo "${cmd}"

				# Run the cluster creation command and check for the desired output
				for i in {1..100}; do
					output=$(${cmd})
					echo "${output}"
					if echo "${output}" | grep -q "clusterId"; then
						echo "Cluster create command issued successfully."
						break
					else
						echo "Cluster create command not yet issued, retrying in 30s..."
						sleep 30
					fi
					if i==100; then
						exit 1
					fi
				done
				sleep 120

				# Wait for the cluster creation to complete and capture the final status output
				final_status_output=""
				for i in {1..100}; do
					sleep 30
					status_output=$(iris_cli --username=admin --password=admin --skip_password_prompt=true --skip_force_password_change=true cluster status)
					echo "${status_output}"
					if echo "${status_output}" | grep -Eq "CLUSTER ACTIVE OPERATION\s*:\s*$"; then
						echo "Cluster creation completed successfully."
						final_status_output=$(iris_cli --output=prettyjson --username=admin --password=admin --skip_password_prompt=true --skip_force_password_change=true cluster status)
						# Extract Node: x IP: [y] pairs from the final status output and print them
						echo "Parsing the final status output to extract Node IP information..."
						echo "${final_status_output}"
						break
					else
						echo "Cluster creation still in progress..."
					fi
					if i==100; then
						exit 1
					fi
				done

				sleep 60
				exit 0`

	return clusterCreationScript
}

//----------------------------------------------------------------------------------------

// ValidateAppSubnetFields validates the app subnets input.
func ValidateAppSubnetFields(appSubnet, appSubnetMask string) error {
	if appSubnet == "" || appSubnetMask == "" {
		return fmt.Errorf("apps_subnet and apps_subnet_mask are required when add_apps is set to true")
	}

	if !utils.IsValidIP(appSubnet) {
		return fmt.Errorf("invalid apps_subnet specified. %s", appSubnet)
	}

	if !utils.IsValidIP(appSubnetMask) {
		return fmt.Errorf("invalid apps_subnet_mask specified. %s", appSubnetMask)
	}
	return nil
}

// ValidateSupportPassword validates the support password input.
func ValidateSupportPassword(password string) error {
	if len(password) < 14 {
		return fmt.Errorf("support password must be at least 14 characters long")
	}

	var checks int
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			checks |= 1
		case unicode.IsLower(char):
			checks |= 2
		case unicode.IsDigit(char):
			checks |= 4
		case strings.ContainsRune("@#$%^&*()_+-=[]{}|;:,.<>?", char):
			checks |= 8
		}
	}

	if checks != 15 {
		return fmt.Errorf("support password must contain at least one uppercase letter, one lowercase letter, one number, and one symbol")
	}

	if strings.ContainsAny(password, "!`\"\\") {
		return fmt.Errorf("!, `, \", and \\ are not allowed in the support password")
	}

	return nil
}
