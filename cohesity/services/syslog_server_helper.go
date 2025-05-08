package services

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/cohesity/go-sdk/v2/client/syslog"
	modelsV2 "github.com/cohesity/go-sdk/v2/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-providers/terraform-provider-cohesity/cohesity/utils"
)

func BuildSyslogServerParam(d *schema.ResourceData, isNew bool) *modelsV2.SyslogServer {
	syslogServer := &modelsV2.SyslogServer{
		Name:          utils.StringPtr(d.Get("name").(string)),
		IP:            utils.StringPtr(d.Get("ip_address").(string)),
		Port:          utils.Int32Ptr(int32(d.Get("port").(int))),
		Protocol:      utils.StringPtr(d.Get("protocol").(string)),
		Enabled:       utils.BoolPtr(d.Get("is_enabled").(bool)),
		CaCertificate: utils.StringPtr(d.Get("ca_certificate").(string)),
	}
	programNameList, err := utils.InterfaceSliceToStringSlice(d.Get("program_names").([]interface{}))
	if err != nil {
		return nil
	}
	syslogServer.ProgramNameList = programNameList
	if isNew {
		syslogServer.IsTLSEnabled = utils.BoolPtr(true)
	} else {
		idInt64, _ := strconv.ParseInt(d.Id(), 10, 32)
		syslogServer.ID = utils.Int32Ptr(int32(idInt64))
	}
	return syslogServer
}
func GetSyslogServer(url, token, id string) (*modelsV2.SyslogServer, error) {
	resp, code, err := GetWithModel(fmt.Sprintf("%s/v2/syslog/%s", url, id), token)
	if err != nil {
		return nil, err
	}
	if code != 202 {
		return nil, fmt.Errorf("the request failed with code %d, resp: %+v", code, string(resp))
	}
	var response modelsV2.SyslogServer
	err = json.Unmarshal(resp, &response)
	if err != nil {
		log.Printf("failed to unmarshal response: %d ,: %v", resp, err)
		return nil, err
	}
	return &response, err
}

func CreateSyslogServer(url, token string, syslogServer *modelsV2.SyslogServer) (*modelsV2.SyslogServer, error) {
	resp, code, err := PostWithModel(fmt.Sprintf("%s/v2/syslog", url), token, syslogServer)
	if err != nil {
		return nil, err
	}
	if code != 202 {
		return nil, fmt.Errorf("the request failed with code %d, resp: %+v", code, string(resp))
	}
	var response modelsV2.SyslogServer
	err = json.Unmarshal(resp, &response)
	if err != nil {
		log.Printf("failed to unmarshal response: %d ,: %v", resp, err)
		return nil, err
	}
	return &response, err
}

func UpdateSyslogServer(url, token, id string, syslogServer *modelsV2.SyslogServer) (*modelsV2.SyslogServer, error) {
	resp, code, err := PutWithModel(fmt.Sprintf("%s/v2/syslog/%s", url, id), token, syslogServer)
	if err != nil {
		return nil, err
	}
	if code != 202 {
		return nil, fmt.Errorf("the request failed with code %d, resp: %+v, for id: %s", code, string(resp), id)
	}
	var response modelsV2.SyslogServer
	err = json.Unmarshal(resp, &response)
	if err != nil {
		log.Printf("failed to unmarshal response: %d ,: %v", resp, err)
		return nil, err
	}
	return &response, err
}

func (c *CohesityClientV2) DeleteSyslogServer(id string) error {
	idInt64, _ := strconv.ParseInt(id, 10, 32)
	_, err := c.client.Syslog.RemoveSyslogServer(syslog.NewRemoveSyslogServerParams().WithID(int32(idInt64)), c.bearerToken)
	return err
}
