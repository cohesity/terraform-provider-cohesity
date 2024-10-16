package services

import (
	"fmt"
	"strconv"

	"github.com/cohesity/go-sdk/v1/client/backup_sources"
	modelsV1 "github.com/cohesity/go-sdk/v1/models"
)

func convertStringToASCIISlice(str string) []uint8 {
	result := make([]uint8, len(str))
	for i, char := range str {
		result[i] = uint8(char)
	}
	return result
}
func (c *CohesityClientV1) AddGCPProtectionSource(clientEmailAddress, clientPrivateKey, projectId, VPCNetwork, VPCSubnetwork string) (int64, error) {
	entityType := int32(20)
	gcpType := int32(0)
	body := &modelsV1.EntityArgWrapper{
		Entity: &modelsV1.PrivateEntityProto{
			Type: &entityType,
			GcpEntity: &modelsV1.PrivateGcpEntity{
				Type: &gcpType,
				CommonInfo: &modelsV1.EntityCommonInfo{
					ID: &clientEmailAddress,
				},
				OwnerID:       &clientEmailAddress,
				ProjectID:     &projectId,
				VpcNetwork:    &VPCNetwork,
				VpcSubnetwork: &VPCSubnetwork,
			},
		},
		EntityInfo: &modelsV1.EntityAccessInfo{
			Type: &entityType,
			Credentials: &modelsV1.PrivateCredentials{
				CloudCredentials: &modelsV1.NexusCloudCredentials{
					GcpCredentials: &modelsV1.PrivateGcpCredentials{
						ClientEmailAddress:        &clientEmailAddress,
						EncryptedClientPrivateKey: convertStringToASCIISlice(clientPrivateKey),
						ProjectID:                 &projectId,
					},
				},
			},
		},
	}
	resp, err := c.client.BackupSources.RegisterBackupSource(backup_sources.NewRegisterBackupSourceParams().WithBody(body), c.bearerToken)
	if err != nil {
		return -1, err
	}
	return *resp.Payload.Entity.ID, nil
}
func (c *CohesityClientV1) GetGcpProtectionSource(id string) (*modelsV1.EntityHierarchyProto, error) {
	// Convert string to int64
	idInt64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %v", err)
	}
	returnOnlyOneLevel := true
	resp, err := c.client.BackupSources.ListBackupSources(backup_sources.NewListBackupSourcesParams().WithEntityID(&idInt64).WithOnlyReturnOneLevel(&returnOnlyOneLevel), c.bearerToken)

	if err != nil {
		return nil, err
	}
	return resp.Payload.EntityHierarchy, nil
}
