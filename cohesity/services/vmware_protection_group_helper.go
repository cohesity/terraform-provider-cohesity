package services

import (
	"github.com/cohesity/go-sdk/v2/client/policy"
	"github.com/cohesity/go-sdk/v2/client/storage_domain"
	modelsV2 "github.com/cohesity/go-sdk/v2/models"
)

func (c *CohesityClientV2) GetProtectionPolicy(policyName string) ([]*modelsV2.ProtectionPolicyResponse, error) {
	resp, err := c.client.Policy.GetProtectionPolicies(policy.NewGetProtectionPoliciesParams().WithPolicyNames([]string{policyName}), c.bearerToken)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Policies, nil
}

func (c *CohesityClientV2) GetStorageDomains(domainName string) ([]*modelsV2.StorageDomain, error) {
	resp, err := c.client.StorageDomain.GetStorageDomains(storage_domain.NewGetStorageDomainsParams().WithNames([]string{domainName}), c.bearerToken)
	if err != nil {
		return nil, err
	}
	return resp.Payload.StorageDomains, nil
}
