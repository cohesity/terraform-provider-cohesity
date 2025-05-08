package services

import (
	smtpConfig "github.com/cohesity/go-sdk/v2/client/platform"
	modelsV2 "github.com/cohesity/go-sdk/v2/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-providers/terraform-provider-cohesity/cohesity/utils"
)

type SMTPConfiguration = modelsV2.SMTPConfiguration

func BuildSmtpConfigParam(d *schema.ResourceData, enable bool) *modelsV2.UpdateSMTPParams {
	smtpConfigParam := &modelsV2.UpdateSMTPParams{
		SMTPConfiguration: modelsV2.SMTPConfiguration{
			Hostname:           d.Get("smtp_server").(string),
			Port:               int32(d.Get("port").(int)),
			// SenderEmailAddress: utils.StringPtr(d.Get("sender_email_address").(string)),
			Username:           utils.StringPtr(d.Get("username").(string)),
			UseSSL:             utils.BoolPtr(d.Get("use_ssl").(bool)),
		},
		Password: utils.StringPtr(d.Get("password").(string)),
	}

	return smtpConfigParam
}

func (c *CohesityClientV2) GetSmtpConfig() (*SMTPConfiguration, error) {
	resp, err := c.client.Platform.GetSMTPConfiguration(smtpConfig.NewGetSMTPConfigurationParams(), c.bearerToken)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (c *CohesityClientV2) EnableSmtpConfig(smtpConfigParam *modelsV2.UpdateSMTPParams) (*SMTPConfiguration, error) {
	params := smtpConfig.NewUpdateSMTPConfigurationParams().WithBody(smtpConfigParam)

	resp, err := c.client.Platform.UpdateSMTPConfiguration(params, c.bearerToken)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *CohesityClientV2) DisableSmtpConfig(smtpConfigParam *modelsV2.UpdateSMTPParams) error {
	_, err := c.client.Platform.ClearSMTPConfiguration(smtpConfig.NewClearSMTPConfigurationParams(), c.bearerToken)
	if err != nil {
		return err
	}

	return nil
}
