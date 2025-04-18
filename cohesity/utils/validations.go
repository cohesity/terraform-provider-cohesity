package utils

import (
	"fmt"
	"net"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// ValidateOptionalIP validates if the ip field is either empty or
func ValidateOptionalIP(val interface{}, key string) (warns []string, errs []error) {
	if val == "" {
		return
	}

	return validation.IsIPAddress(val, key)
}

// ValidateCommaSeparatedIPList ensures that a comma-separated string contains valid IP addresses
func ValidateCommaSeparatedIPList(val interface{}, key string) (warns []string, errs []error) {
	ipList := strings.Split(val.(string), ",")
	for _, ip := range ipList {
		ip = strings.TrimSpace(ip)
		if !IsValidIP(ip) {
			errs = append(errs, fmt.Errorf("%q contains an invalid IP address: %s", key, ip))
		}
	}
	return
}

// IsValidIP checks if the IP is a valid one.
func IsValidIP(ip string) bool {
	// trim the value.
	ip = strings.TrimSpace(ip)

	parsedIP := net.ParseIP(ip)
	return parsedIP != nil
}

// AreValidIPs checks if multiple ips are valid
func AreValidIPs(nodeIps []string) error {
	for _, ip := range nodeIps {
		if !IsValidIP(ip) {
			return fmt.Errorf("%q is not a valid IP address", ip)
		}
	}

	return nil
}

func IsValidNetappSourceType(val interface{}, key string) (warns []string, errs []error) {
	v := val.(string)
	if v != "kCluster" && v != "kVserver" {
		errs = append(errs, fmt.Errorf("%q must be either 'kCluster' or 'kVserver', got: %s", key, v))
	}
	return
}
