package cohesity

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/terraform-providers/terraform-provider-cohesity/cohesity/services"
	"github.com/terraform-providers/terraform-provider-cohesity/cohesity/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

type CohesityClusterConfig struct {
	Name                   string           `json:"name"`
	NodeIps                []string         `json:"nodeIps"`
	Hostname               string           `json:"hostname"`
	SubnetGateway          string           `json:"subnetGateway"`
	SubnetMask             string           `json:"subnetMask"`
	DnsServerIps           []string         `json:"dnsServerIps"`
	DomainNames            string           `json:"domainNames"`
	NtpServers             string           `json:"ntpServers"`
	MetadataFaultTolerance int              `json:"metadataFaultTolerance"`
	AddApps                bool             `json:"addApps"`
	AppsSubnet             string           `json:"appsSubnet"`
	AppsSubnetMask         string           `json:"appsSubnetMask"`
	NodeIPMap              map[int][]string `json:"nodeIPMap"`
}

// ---------------------------------------------------------------------------------
func resourceCohesityNGCECluster() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCohesityNGCEClusterCreate,
		ReadContext:   resourceCohesityNGCEClusterRead,
		UpdateContext: resourceCohesityNGCEClusterUpdate,
		DeleteContext: resourceCohesityNGCEClusterDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"node_ips": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subnet_gateway": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsIPAddress,
			},
			"subnet_mask": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsIPAddress,
			},
			"dns_server_ips": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.ValidateCommaSeparatedIPList,
			},
			"ntp_servers": {
				Type:     schema.TypeString,
				Required: true,
			},
			"domain_names": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"metadata_fault_tolerance": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntAtLeast(0),
			},
			"add_apps": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"apps_subnet": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: utils.ValidateOptionalIP,
			},
			"apps_subnet_mask": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: utils.ValidateOptionalIP,
			},
			"node_ip_map": {
				Type: schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"node_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"node_ips": {
							Type: schema.TypeList,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Computed: true,
						},
					},
				},
				Computed: true,
			},
		},
	}
}

//---------------------------------------------------------------------------------

// setClusterConfigParamsFromSchema sets the schema vars into a goStruct.
func setClusterConfigParamsFromSchema(d *schema.ResourceData) CohesityClusterConfig {
	// Convert []interface into []string
	convertToStringArr := func(inputs []interface{}) []string {
		inputArr := make([]string, 0)
		for _, input := range inputs {
			inputStr := strings.TrimSpace(input.(string))
			inputArr = append(inputArr, inputStr)
		}

		return inputArr
	}

	// parse the node ip map.
	hashMap := make(map[int][]string)
	nodeIpMap := d.Get("node_ip_map").([]interface{})
	for _, v := range nodeIpMap {
		// Each item in nodeIpMap is a map representing the nested resource
		nodeMap := v.(map[string]interface{})

		// Reading node_id
		nodeID := nodeMap["node_id"].(int)

		// Reading node_ips (which is a list of strings)
		nodeIPs := nodeMap["node_ips"].([]interface{})
		ipList := make([]string, len(nodeIPs))
		for i, ip := range nodeIPs {
			ipList[i] = ip.(string)
		}

		// Add to hashmap.
		hashMap[nodeID] = ipList
	}

	return CohesityClusterConfig{
		Name:                   d.Get("name").(string),
		NodeIps:                convertToStringArr(d.Get("node_ips").([]interface{})),
		Hostname:               d.Get("hostname").(string),
		SubnetGateway:          d.Get("subnet_gateway").(string),
		SubnetMask:             d.Get("subnet_mask").(string),
		DnsServerIps:           strings.Split(d.Get("dns_server_ips").(string), ","),
		DomainNames:            d.Get("domain_names").(string),
		NtpServers:             d.Get("ntp_servers").(string),
		MetadataFaultTolerance: d.Get("metadata_fault_tolerance").(int),
		AddApps:                d.Get("add_apps").(bool),
		AppsSubnet:             d.Get("apps_subnet").(string),
		AppsSubnetMask:         d.Get("apps_subnet_mask").(string),
		NodeIPMap:              hashMap,
	}
}

//---------------------------------------------------------------------------------

// resourceCohesityNGCEClusterCreate Create a new Cluster
func resourceCohesityNGCEClusterCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[INFO] creating cluster")
	// Parse the input params.
	clusterParams := setClusterConfigParamsFromSchema(d)

	config := m.(Config)
	supportPassword := config.SupportPassword
	if err := ValidateSupportPassword(supportPassword); err != nil {
		return diag.FromErr(err)
	}

	// Validate the IPs.
	if err := utils.AreValidIPs(clusterParams.NodeIps); err != nil {
		return diag.FromErr(err)
	}

	// Validate App subnet fields.
	if clusterParams.AddApps {
		if err := ValidateAppSubnetFields(clusterParams.AppsSubnet, clusterParams.AppsSubnetMask); err != nil {
			return diag.FromErr(err)
		}
	}

	// Execute the cluster create script with the config params.
	// The script returns the cluster status as JSON after the cluster has been successfully created.
	output, err := execCreateClusterCmd(clusterParams)
	if err != nil {
		return diag.FromErr(err)
	}

	// Parse the cluster status output after cluster has been created.
	clusterStatusJsonResp := extractJSONFromCreateClusterCmdOutput(output)
	clusterInfo, err := parseClusterStatusJSONResponse(clusterStatusJsonResp)
	if err != nil {
		return diag.FromErr(err)
	}

	// Check if all the nodes have been added successfully or not.
	_, _, areAllNodesPartOfCluster := getNodeInformationFromClusterStatus(*clusterInfo)

	if !areAllNodesPartOfCluster {
		return diag.FromErr(fmt.Errorf("all nodes are not part of the cluster. " +
			"something went wrong during create cluster"))
	}

	// Enable support user login.
	log.Printf("[INFO] enabling support user...")
	err = enableSupportUser(clusterParams.Hostname, supportPassword)
	if err != nil {
		log.Printf("failed to enable support user on the cluster %s", clusterParams.Hostname)
		return diag.FromErr(err)
	}

	log.Printf("[INFO] cluster created successfully. Cluster Id %d", clusterInfo.ClusterID)

	// Set the id of the resource
	d.SetId(strconv.FormatInt(clusterInfo.ClusterID, 10))

	return resourceCohesityNGCEClusterRead(ctx, d, m)
}

//---------------------------------------------------------------------------------

// resourceCohesityNGCEClusterRead fetches the current cluster state.
func resourceCohesityNGCEClusterRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Print("Reading current state of cluster resource")
	// Parse the input params.
	clusterParams := setClusterConfigParamsFromSchema(d)

	config := m.(Config)
	clusterUsername := config.ClusterUsername
	clusterPassword := config.ClusterPassword
	escapedClusterPassword := utils.EscapeSpecialSymbols(clusterPassword)
	supportPassword := config.SupportPassword

	log.Printf("[INFO] Cluster host is %s", clusterParams.Hostname)

	// SSH into the cluster.
	client, err := services.NewSSHClient(clusterParams.Hostname, "support", supportPassword, time.Hour)
	if err != nil {
		log.Printf("failed to init ssh client on host %s", clusterParams.Hostname)
		return diag.FromErr(err)
	}

	// Fetch the cluster status through making an API call.
	token, err := services.GetAccessToken(clusterParams.Hostname, clusterUsername, clusterPassword)
	if err != nil {
		log.Printf("failed to get access token. check credentials %s", err.Error())
		return diag.FromErr(err)
	}

	// Fetch the clusters details /v2/clusters.
	clusterInfo, err := services.GetClusterInfo(clusterParams.Hostname, token)
	if err != nil {
		log.Printf("failed to fetch the cluster info. error %s", err.Error())
		return diag.FromErr(err)
	}

	// Fetch the cluster status
	clusterStatus, err := fetchClusterStatus(clusterUsername, escapedClusterPassword, client)
	if err != nil {
		log.Printf("failed to fetch the cluster status. error %s", err.Error())
		return diag.FromErr(err)
	}

	log.Printf("[INFO] cluster status: %+v", *clusterStatus)
	nodeIpsMap, nodeIps, _ := getNodeInformationFromClusterStatus(*clusterStatus)

	// Update the information in the struct.
	clusterParams.Name = clusterInfo.Name
	clusterParams.NodeIps = nodeIps
	clusterParams.Hostname = clusterInfo.NetworkConfig.VipHostName
	clusterParams.SubnetGateway = clusterInfo.NetworkConfig.ManualNetworkConfig.Gateway
	clusterParams.SubnetMask = clusterInfo.NetworkConfig.ManualNetworkConfig.SubnetMask
	clusterParams.DnsServerIps = clusterInfo.NetworkConfig.ManualNetworkConfig.DnsServers
	clusterParams.NtpServers = strings.Join(clusterInfo.NetworkConfig.NtpServers, ",")
	clusterParams.DomainNames = strings.Join(clusterInfo.NetworkConfig.DomainNames, ",")

	// Set the node id map.
	if err := d.Set("node_ip_map", nodeIpsMap); err != nil {
		return diag.FromErr(err)
	}

	// Save the attributes in the state
	d.Set("name", clusterParams.Name)
	d.Set("node_ips", clusterParams.NodeIps)
	d.Set("hostname", clusterParams.Hostname)
	d.Set("subnet_gateway", clusterParams.SubnetGateway)
	d.Set("subnet_mask", clusterParams.SubnetMask)
	d.Set("dns_server_ips", strings.Join(clusterParams.DnsServerIps, ","))
	d.Set("ntp_servers", clusterParams.NtpServers)
	d.Set("domain_names", clusterParams.DomainNames)
	d.Set("metadata_fault_tolerance", clusterParams.MetadataFaultTolerance)
	d.Set("add_apps", clusterParams.AddApps)

	if clusterParams.AddApps {
		d.Set("apps_subnet", clusterParams.AppsSubnet)
		d.Set("apps_subnet_mask", clusterParams.AppsSubnetMask)
	}

	return nil
}

//---------------------------------------------------------------------------------

// Update the Cluster
func resourceCohesityNGCEClusterUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("Updating cluster resource request..")

	// Parse the input params.
	clusterParams := setClusterConfigParamsFromSchema(d)

	config := m.(Config)
	clusterPassword := config.ClusterPassword
	escapedClusterPassword := utils.EscapeSpecialSymbols(clusterPassword)
	supportPassword := config.SupportPassword
	clusterUsername := config.ClusterUsername

	if d.HasChange("node_ips") {
		// There are changes in the node ip. Add or remove node op req.
		log.Printf("Node Ips have changed.")

		oldNodeIPsInterface, newNodeIPsInterface := d.GetChange("node_ips")
		// Assert the types to []interface{}
		oldNodeIPs, ok1 := oldNodeIPsInterface.([]interface{})
		newNodeIPs, ok2 := newNodeIPsInterface.([]interface{})
		// Check if type assertion succeeded
		if !ok1 || !ok2 {
			log.Fatalf("Failed to assert types of oldNodeIPs or newNodeIPs")
		}
		log.Printf("[INFO] ips %v to %v", oldNodeIPs, newNodeIPs)
		// Find removed and added nodes
		removedNodes, addedNodes := difference(oldNodeIPs, newNodeIPs)
		log.Printf("[INFO] Cluster node IPs updated to: %v", newNodeIPs)
		nodeIPMap := d.Get("node_ip_map").([]interface{})

		// Init SSH client.
		client, err := services.NewSSHClient(clusterParams.Hostname, "support", supportPassword, time.Hour)
		if err != nil {
			log.Printf("failed to init ssh client using support user on host %s", clusterParams.Hostname)
			d.Set("node_ips", oldNodeIPs)
			return diag.FromErr(err)
		}
		defer client.Close()

		// If nodes are to be added.
		if len(addedNodes) > 0 {
			var nodeIPsStrings []string
			for _, ip := range addedNodes {
				// Assert the type of interface{} to string
				if ipStr, ok := ip.(string); ok {
					nodeIPsStrings = append(nodeIPsStrings, ipStr)
				}
			}
			nodeIPsString := strings.Join(nodeIPsStrings, ",")
			err := execAddNodesCmd(clusterUsername, escapedClusterPassword, nodeIPsString, client)
			if err != nil {
				log.Printf("failed to add nodes. error %s", err.Error())
				return diag.FromErr(err)
			}
		}

		// Remove nodes.
		for _, node := range removedNodes {
			log.Printf("[INFO] nodeMap: %v, node: %v", nodeIPMap, node)
			nodeID, ok := getNodeIDByIP(nodeIPMap, node.(string))
			if !ok {
				continue
			}
			// Nodes are removed one at a time.
			err := execRemoveNodeCmd(clusterUsername, escapedClusterPassword, nodeID, client)
			if err != nil {
				log.Printf("failed to remove node id %d. error %s", nodeID, err.Error())
				return diag.FromErr(err)
			}

		}
	}

	// If there are changes in other params.
	if d.HasChangesExcept("node_ips") {
		// Fetch the cluster status through making an API call.
		token, err := services.GetAccessToken(clusterParams.Hostname, clusterUsername, clusterPassword)
		if err != nil {
			log.Printf("failed to get access token. check credentials %s", err.Error())
			return diag.FromErr(err)
		}

		// Fetch the clusters details /v2/clusters.
		currentClusterInfo, err := services.GetClusterInfo(clusterParams.Hostname, token)
		if err != nil {
			log.Printf("failed to fetch the cluster info. error %s", err.Error())
			return diag.FromErr(err)
		}

		log.Printf("cluster update params %+v", clusterParams)
		// Update the vars.
		currentClusterInfo.Name = clusterParams.Name
		currentClusterInfo.NetworkConfig.VipHostName = clusterParams.Hostname
		currentClusterInfo.NetworkConfig.ManualNetworkConfig.Gateway = clusterParams.SubnetGateway
		currentClusterInfo.NetworkConfig.ManualNetworkConfig.SubnetMask = clusterParams.SubnetMask
		currentClusterInfo.NetworkConfig.ManualNetworkConfig.DnsServers = clusterParams.DnsServerIps
		currentClusterInfo.NetworkConfig.NtpServers = strings.Split(clusterParams.NtpServers, ",")
		currentClusterInfo.NetworkConfig.DomainNames = strings.Split(clusterParams.DomainNames, ",")

		// Make API call to update the cluster info.
		_, err = services.UpdateClusterInfo(clusterParams.Hostname, *currentClusterInfo, token)
		if err != nil {
			log.Printf("failed to update the cohesity cluster info. err %s", err.Error())
			return diag.FromErr(err)
		}
	}

	return resourceCohesityNGCEClusterRead(ctx, d, m)
}

//---------------------------------------------------------------------------------

// Delete the Cluster
func resourceCohesityNGCEClusterDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clusterID := d.Id()
	log.Printf("[INFO] Deleting Cluster: %s", clusterID)

	d.SetId("")

	return nil
}

//---------------------------------------------------------------------------------
