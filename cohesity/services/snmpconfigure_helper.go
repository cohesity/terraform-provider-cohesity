package services

import (
	snmpConfig "github.com/cohesity/go-sdk/v2/client/snmp_config"
	modelsV2 "github.com/cohesity/go-sdk/v2/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-providers/terraform-provider-cohesity/cohesity/utils"
)

type SnmpConfig = modelsV2.SnmpConfig

func BuildSnmpUserMap(user *modelsV2.SnmpUser) []interface{} {
	userMap := make(map[string]interface{})
	userMap["user_name"] = user.UserName

	if user.AuthProtocol != nil {
		if *user.AuthProtocol == "kAuthMD5" {
			userMap["auth_protocol"] = "MD5"
		} else if *user.AuthProtocol == "kAuthSHA" {
			userMap["auth_protocol"] = "SHA"
		}
	}

	if user.AuthPassword != nil {
		userMap["auth_password"] = user.AuthPassword
	}

	if user.PrivProtocol != nil {
		if *user.PrivProtocol == "kPrivAES" {
			userMap["priv_protocol"] = "AES"
		} else if *user.PrivProtocol == "kPrivDES" {
			userMap["priv_protocol"] = "DES"
		}
	}

	if user.PrivPassword != nil {
		userMap["priv_password"] = user.PrivPassword
	}

	return []interface{}{userMap}
}

func BuildSnmpUserModel(data interface{}, snmpVersion string) *modelsV2.SnmpUser {
	m, ok := data.(map[string]interface{})
	if !ok {
		return nil
	}

	user := &modelsV2.SnmpUser{}
	var userName string

	if snmpVersion == "V2C" {
		userName = "cohesityV2Public"
		user.AuthProtocol = utils.StringPtr("kAuthMD5")
		user.PrivProtocol = utils.StringPtr("kPrivDES")
		user.SecurityLevel = utils.StringPtr("kAuthPriv")
	} else {
		userName = "cohesityV3Public"
		authProtocol := "None"
		privProtocol := "None"

		if val, ok := m["auth_protocol"].(string); ok && val != "" {
			if val == "MD5" {
				authProtocol = "kAuthMD5"
			} else if val == "SHA" {
				authProtocol = "kAuthSHA"
			}
		}

		if authProtocol != "None" {
			user.AuthProtocol = utils.StringPtr(authProtocol)
			if password, ok := m["auth_password"].(string); ok && password != "" {
				user.AuthPassword = utils.StringPtr(password)
			}

			if priv, ok := m["priv_protocol"].(string); ok && priv != "" {
				if priv == "AES" {
					privProtocol = "kPrivAES"
				} else if priv == "DES" {
					privProtocol = "kPrivDES"
				}
			}

			if privProtocol != "None" {
				user.PrivProtocol = utils.StringPtr(privProtocol)
				if password, ok := m["priv_password"].(string); ok && password != "" {
					user.PrivPassword = utils.StringPtr(password)
				}
			}
		}

		switch {
		case authProtocol == "None":
			user.SecurityLevel = utils.StringPtr("kNoAuthNoPriv")
		case privProtocol == "None":
			user.SecurityLevel = utils.StringPtr("kAuthNoPriv")
		default:
			user.SecurityLevel = utils.StringPtr("kAuthPriv")
		}

		user.UserName = utils.StringPtr(userName)
	}

	if val, ok := m["user_name"].(string); ok && val != "" {
		userName = val
	}
	user.UserName = utils.StringPtr(userName)

	return user
}

func BuildSnmpConfigParam(d *schema.ResourceData, enable bool) *modelsV2.SnmpConfig {
	snmpConfigParam := &modelsV2.SnmpConfig{
		AgentPort: utils.Int32Ptr(int32(d.Get("agent_port").(int))),
		TrapPort:  utils.Int32Ptr(int32(d.Get("trap_port").(int))),
		Server:    utils.StringPtr(d.Get("server").(string)),
	}

	if enable {
		snmpConfigParam.Operation = utils.StringPtr("kOperationEnable")
	} else {
		snmpConfigParam.Operation = utils.StringPtr("kOperationDisable")
	}

	if d.Get("snmp_version").(string) == "V2C" {
		snmpConfigParam.Version = utils.StringPtr("kSnmpV2")
	} else {
		snmpConfigParam.Version = utils.StringPtr("kSnmpV3")
	}

	if v, ok := d.GetOk("read_user"); ok && len(v.([]interface{})) > 0 {
		snmpConfigParam.ReadUser = BuildSnmpUserModel(v.([]interface{})[0], d.Get("snmp_version").(string))
	}

	if v, ok := d.GetOk("write_user"); ok && len(v.([]interface{})) > 0 {
		snmpConfigParam.WriteUser = BuildSnmpUserModel(v.([]interface{})[0], d.Get("snmp_version").(string))
	}

	var trapUserData interface{}
	if v, ok := d.GetOk("trap_user"); ok && len(v.([]interface{})) > 0 {
		trapUserData = v.([]interface{})[0]
	} else {
		trapUserData = map[string]interface{}{}
	}
	snmpConfigParam.TrapUser = BuildSnmpUserModel(trapUserData, d.Get("snmp_version").(string))

	snmpConfigParam.SystemInfo = &modelsV2.SnmpSysInfo{}

	return snmpConfigParam
}

func (c *CohesityClientV2) GetSnmpConfig() (*SnmpConfig, error) {
	resp, err := c.client.SnmpConfig.GetSnmpConfig(snmpConfig.NewGetSnmpConfigParams(), c.bearerToken)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (c *CohesityClientV2) UpdateSnmpConfig(snmpConfigParam *modelsV2.SnmpConfig) (*SnmpConfig, error) {
	params := snmpConfig.NewUpdateSnmpConfigParams().WithBody(snmpConfigParam)

	resp, err := c.client.SnmpConfig.UpdateSnmpConfig(params, c.bearerToken)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}
