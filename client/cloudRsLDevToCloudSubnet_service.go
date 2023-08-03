package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/v2/models"
)

func (sm *ServiceManager) CreateRelationFromCloudLDevToCloudSubnet(tDn string, parentDn string, cloudRsLDevToCloudSubnetAttr models.RelationFromCloudLDevToCloudSubnetAttributes) (*models.RelationFromCloudLDevToCloudSubnet, error) {
	rn := fmt.Sprintf(models.RnCloudRsLDevToCloudSubnet, tDn)
	cloudRsLDevToCloudSubnet := models.NewRelationFromCloudLDevToCloudSubnet(rn, parentDn, cloudRsLDevToCloudSubnetAttr)
	err := sm.Save(cloudRsLDevToCloudSubnet)
	return cloudRsLDevToCloudSubnet, err
}

func (sm *ServiceManager) ReadRelationFromCloudLDevToCloudSubnet(tDn string, parentDn string) (*models.RelationFromCloudLDevToCloudSubnet, error) {
	rn := fmt.Sprintf(models.RnCloudRsLDevToCloudSubnet, tDn)
	dn := fmt.Sprintf("%s/%s", parentDn, rn)
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}
	cloudRsLDevToCloudSubnet := models.RelationFromCloudLDevToCloudSubnetFromContainer(cont)
	return cloudRsLDevToCloudSubnet, nil
}

func (sm *ServiceManager) DeleteRelationFromCloudLDevToCloudSubnet(tDn string, parentDn string) error {
	rn := fmt.Sprintf(models.RnCloudRsLDevToCloudSubnet, tDn)
	dn := fmt.Sprintf("%s/%s", parentDn, rn)
	return sm.DeleteByDn(dn, models.CloudRsLDevToCloudSubnetClassName)
}

func (sm *ServiceManager) UpdateRelationFromCloudLDevToCloudSubnet(tDn string, parentDn string, cloudRsLDevToCloudSubnetAttr models.RelationFromCloudLDevToCloudSubnetAttributes) (*models.RelationFromCloudLDevToCloudSubnet, error) {
	rn := fmt.Sprintf(models.RnCloudRsLDevToCloudSubnet, tDn)
	cloudRsLDevToCloudSubnet := models.NewRelationFromCloudLDevToCloudSubnet(rn, parentDn, cloudRsLDevToCloudSubnetAttr)
	cloudRsLDevToCloudSubnet.Status = "modified"
	err := sm.Save(cloudRsLDevToCloudSubnet)
	return cloudRsLDevToCloudSubnet, err
}

func (sm *ServiceManager) ListRelationFromCloudLDevToCloudSubnet(parentDn string) ([]*models.RelationFromCloudLDevToCloudSubnet, error) {
	dnUrl := fmt.Sprintf("%s/%s/%s.json", models.BaseurlStr, parentDn, models.CloudRsLDevToCloudSubnetClassName)
	cont, err := sm.GetViaURL(dnUrl)
	list := models.RelationFromCloudLDevToCloudSubnetListFromContainer(cont)
	return list, err
}
