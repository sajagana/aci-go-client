package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/v2/models"
)

func (sm *ServiceManager) CreateRelationFromCloudLDevToCloudCtx(cloud_l4_l7_device string, tenant string, description string, cloudRsLDevToCtxAttr models.RelationFromCloudLDevToCloudCtxAttributes) (*models.RelationFromCloudLDevToCloudCtx, error) {
	parentDn := fmt.Sprintf(models.ParentDnCloudRsLDevToCtx, tenant, cloud_l4_l7_device)
	cloudRsLDevToCtx := models.NewRelationFromCloudLDevToCloudCtx(models.RnCloudRsLDevToCtx, parentDn, description, cloudRsLDevToCtxAttr)
	err := sm.Save(cloudRsLDevToCtx)
	return cloudRsLDevToCtx, err
}

func (sm *ServiceManager) ReadRelationFromCloudLDevToCloudCtx(cloud_l4_l7_device string, tenant string) (*models.RelationFromCloudLDevToCloudCtx, error) {
	parentDn := fmt.Sprintf(models.ParentDnCloudRsLDevToCtx, tenant, cloud_l4_l7_device)
	dn := fmt.Sprintf("%s/%s", parentDn, models.RnCloudRsLDevToCtx)
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}
	cloudRsLDevToCtx := models.RelationFromCloudLDevToCloudCtxFromContainer(cont)
	return cloudRsLDevToCtx, nil
}

func (sm *ServiceManager) DeleteRelationFromCloudLDevToCloudCtx(cloud_l4_l7_device string, tenant string) error {
	parentDn := fmt.Sprintf(models.ParentDnCloudRsLDevToCtx, tenant, cloud_l4_l7_device)
	dn := fmt.Sprintf("%s/%s", parentDn, models.RnCloudRsLDevToCtx)
	return sm.DeleteByDn(dn, models.CloudRsLDevToCtxClassName)
}

func (sm *ServiceManager) UpdateRelationFromCloudLDevToCloudCtx(cloud_l4_l7_device string, tenant string, description string, cloudRsLDevToCtxAttr models.RelationFromCloudLDevToCloudCtxAttributes) (*models.RelationFromCloudLDevToCloudCtx, error) {
	parentDn := fmt.Sprintf(models.ParentDnCloudRsLDevToCtx, tenant, cloud_l4_l7_device)
	cloudRsLDevToCtx := models.NewRelationFromCloudLDevToCloudCtx(models.RnCloudRsLDevToCtx, parentDn, description, cloudRsLDevToCtxAttr)
	cloudRsLDevToCtx.Status = "modified"
	err := sm.Save(cloudRsLDevToCtx)
	return cloudRsLDevToCtx, err
}

func (sm *ServiceManager) ListRelationFromCloudLDevToCloudCtx(cloud_l4_l7_device string, tenant string) ([]*models.RelationFromCloudLDevToCloudCtx, error) {
	parentDn := fmt.Sprintf(models.ParentDnCloudRsLDevToCtx, tenant, cloud_l4_l7_device)
	dnUrl := fmt.Sprintf("%s/%s/%s.json", models.BaseurlStr, parentDn, models.CloudRsLDevToCtxClassName)
	cont, err := sm.GetViaURL(dnUrl)
	list := models.RelationFromCloudLDevToCloudCtxListFromContainer(cont)
	return list, err
}
