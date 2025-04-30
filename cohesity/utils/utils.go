package utils

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// WaitTimeToSeconds is used to convert operation time to seconds
const WaitTimeToSeconds int = 60

func EscapeSpecialSymbols(input string) string {
	specialSymbols := `\.$^*+?()[]{}|`

	var escaped strings.Builder
	for _, ch := range input {
		if strings.ContainsRune(specialSymbols, ch) {
			escaped.WriteRune('\\')
		}
		escaped.WriteRune(ch)
	}

	return escaped.String()
}

//----------------------------------------------------------------------------------------

func ListDifference(oldList, newList []interface{}) (removed []interface{}, added []interface{}) {
	oldMap := make(map[interface{}]bool)
	newMap := make(map[interface{}]bool)

	// Populate maps
	for _, old := range oldList {
		oldMap[old] = true
	}
	for _, newIP := range newList {
		newMap[newIP] = true
	}

	// Determine removed items
	for _, old := range oldList {
		if !newMap[old] {
			removed = append(removed, old)
		}
	}

	// Determine added items
	for _, newIP := range newList {
		if !oldMap[newIP] {
			added = append(added, newIP)
		}
	}

	return removed, added
}
func SuppressNodeIPsDiff(k, old, new string, d *schema.ResourceData) bool {
	oldNodeIPsInterface, newNodeIPsInterface := d.GetChange("node_ips")
	if oldNodeIPsInterface == nil {
        oldNodeIPsInterface = []interface{}{}
    }
    if newNodeIPsInterface == nil {
        newNodeIPsInterface = []interface{}{}
    }
	// Assert the types to []interface{}
	oldNodeIPs, ok1 := oldNodeIPsInterface.([]interface{})
	newNodeIPs, ok2 := newNodeIPsInterface.([]interface{})
	// Check if type assertion succeeded
	if !ok1 || !ok2 {
		log.Fatalf("Failed to assert types of oldNodeIPs or newNodeIPs")
	}
	log.Printf("[INFO] ips %v to %v", oldNodeIPs, newNodeIPs)
	// Find removed and added nodes
	removedNodes, addedNodes := ListDifference(oldNodeIPs, newNodeIPs)
	if len(removedNodes) == 0 && len(addedNodes) == 0 {
		return true
	}
	return false
}
func SuppressNetworkNameDiff(k, old, new string, d *schema.ResourceData) bool {
	return strings.HasSuffix(old, "networks/"+new)
}

func Int32Ptr(i int32) *int32 {
	return &i
}

func StringPtr(s string) *string {
	return &s
}

func BoolPtr(b bool) *bool {
	return &b
}
