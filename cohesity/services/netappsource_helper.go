package services

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-openapi/swag"

	"github.com/cohesity/go-sdk/v2/client/source"
	modelsV2 "github.com/cohesity/go-sdk/v2/models"
)

type sourceNetapp struct {
	modelsV2.CommonSourceRegistrationRequestParams
	NetappParams *modelsV2.NetappRegistrationParams `json:"netappParams,omitempty"`
}

func NewSourceNetapp(username, password, endpoint, sourceType string, backupSMBVolumes bool) *sourceNetapp {
	return &sourceNetapp{
		CommonSourceRegistrationRequestParams: modelsV2.CommonSourceRegistrationRequestParams{
			Environment: toStrPtr("kNetapp"),
		},
		NetappParams: &modelsV2.NetappRegistrationParams{
			Credentials: &modelsV2.Credentials{
				Username: &username,
				Password: &password,
			},
			Endpoint:         &endpoint,
			SourceType:       &sourceType,
			BackUpSMBVolumes: &backupSMBVolumes,
		},
	}
}

// Option functions for optional parameters
func (m *sourceNetapp) WithSMBCredentials(smb_username, smb_password string) *sourceNetapp {
	m.NetappParams.SmbCredentials.Username = &smb_username
	m.NetappParams.SmbCredentials.Password = &smb_password
	return m
}

func (m *sourceNetapp) WithAllowedIPAddresses(allowedIpAddresses []string) *sourceNetapp {
	if m.NetappParams.FilterIPConfig == nil {
		m.NetappParams.FilterIPConfig = &modelsV2.FilterIPConfig{}
	}
	m.NetappParams.FilterIPConfig.AllowedIPAddresses = allowedIpAddresses
	return m
}
func (m *sourceNetapp) WithDeniedIPAddresses(deniedIpAddresses []string) *sourceNetapp {
	if m.NetappParams.FilterIPConfig == nil {
		m.NetappParams.FilterIPConfig = &modelsV2.FilterIPConfig{}
	}
	m.NetappParams.FilterIPConfig.DeniedIPAddresses = deniedIpAddresses
	return m
}

// MarshalBinary interface implementation
func (m *sourceNetapp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *sourceNetapp) UnmarshalBinary(b []byte) error {
	var res sourceNetapp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

func (c *CohesityClientV2) AddNetappProtectionSource(netAppSource *sourceNetapp) (int64, error) {
	body := &modelsV2.SourceRegistrationRequestParams{}
	marshalledData, err := netAppSource.MarshalBinary()
	if err != nil {
		return -1, err
	}
	err = body.UnmarshalBinary(marshalledData)
	if err != nil {
		return -1, err
	}

	resp, err := c.client.Source.RegisterProtectionSource(source.NewRegisterProtectionSourceParams().WithBody(body).WithTimeout(300*time.Second), c.bearerToken)
	if err != nil {
		return -1, err
	}

	return *resp.Payload.SourceID, nil
}

func (c *CohesityClientV2) GetNetAppSource(id string) (*sourceNetapp, error) {
	// Convert string to int64
	idInt64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %v", err)
	}

	resp, err := c.client.Source.GetProtectionSourceRegistration(source.NewGetProtectionSourceRegistrationParams().WithID(idInt64), c.bearerToken)
	if err != nil {
		return nil, err
	}
	marshalledData, err := resp.Payload.MarshalBinary()
	if err != nil {
		return nil, err
	}
	sourceResp := &sourceNetapp{}
	err = sourceResp.UnmarshalBinary(marshalledData)
	if err != nil {
		return nil, err
	}
	return sourceResp, nil
}

func (c *CohesityClientV2) UpdateNetAppSource(netAppSource *sourceNetapp,id string) (int64, error) {
	body := &modelsV2.SourceRegistrationUpdateRequestParams{}
	marshalledData, err := netAppSource.MarshalBinary()
	if err != nil {
		return -1, err
	}
	err = body.UnmarshalBinary(marshalledData)
	if err != nil {
		return -1, err
	}
	idInt64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return -1, fmt.Errorf("invalid id: %v", err)
	}
	resp, err := c.client.Source.UpdateProtectionSourceRegistration(source.NewUpdateProtectionSourceRegistrationParams().WithBody(body).WithTimeout(300*time.Second).WithID(idInt64), c.bearerToken)
	if err != nil {
		return -1, fmt.Errorf("err: %v \n body:%v", err,convertASCIISliceToString(marshalledData))
	}
	return *resp.Payload.SourceID, nil

}
