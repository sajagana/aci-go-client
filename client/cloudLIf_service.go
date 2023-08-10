package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/v2/models"
)

func (sm *ServiceManager) CreateCloudL4L7LogicalInterface(name string, cloud_l4_l7_device string, tenant string, cloudLIfAttr models.CloudL4L7LogicalInterfaceAttributes) (*models.CloudL4L7LogicalInterface, error) {
	rn := fmt.Sprintf(models.RnCloudLIf, name)
	parentDn := fmt.Sprintf(models.ParentDnCloudLIf, tenant, cloud_l4_l7_device)
	cloudLIf := models.NewCloudL4L7LogicalInterface(rn, parentDn, cloudLIfAttr)
	err := sm.Save(cloudLIf)
	return cloudLIf, err
}

func (sm *ServiceManager) ReadCloudL4L7LogicalInterface(name string, cloud_l4_l7_device string, tenant string) (*models.CloudL4L7LogicalInterface, error) {
	rn := fmt.Sprintf(models.RnCloudLIf, name)
	parentDn := fmt.Sprintf(models.ParentDnCloudLIf, tenant, cloud_l4_l7_device)
	dn := fmt.Sprintf("%s/%s", parentDn, rn)
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}
	cloudLIf := models.CloudL4L7LogicalInterfaceFromContainer(cont)
	return cloudLIf, nil
}

func (sm *ServiceManager) DeleteCloudL4L7LogicalInterface(name string, cloud_l4_l7_device string, tenant string) error {
	rn := fmt.Sprintf(models.RnCloudLIf, name)
	parentDn := fmt.Sprintf(models.ParentDnCloudLIf, tenant, cloud_l4_l7_device)
	dn := fmt.Sprintf("%s/%s", parentDn, rn)
	return sm.DeleteByDn(dn, models.CloudLIfClassName)
}

func (sm *ServiceManager) UpdateCloudL4L7LogicalInterface(name string, cloud_l4_l7_device string, tenant string, cloudLIfAttr models.CloudL4L7LogicalInterfaceAttributes) (*models.CloudL4L7LogicalInterface, error) {
	rn := fmt.Sprintf(models.RnCloudLIf, name)
	parentDn := fmt.Sprintf(models.ParentDnCloudLIf, tenant, cloud_l4_l7_device)
	cloudLIf := models.NewCloudL4L7LogicalInterface(rn, parentDn, cloudLIfAttr)
	cloudLIf.Status = "modified"
	err := sm.Save(cloudLIf)
	return cloudLIf, err
}

func (sm *ServiceManager) ListCloudL4L7LogicalInterface(cloud_l4_l7_device string, tenant string) ([]*models.CloudL4L7LogicalInterface, error) {
	parentDn := fmt.Sprintf(models.ParentDnCloudLIf, tenant, cloud_l4_l7_device)
	dnUrl := fmt.Sprintf("%s/%s/%s.json", models.BaseurlStr, parentDn, models.CloudLIfClassName)
	cont, err := sm.GetViaURL(dnUrl)
	list := models.CloudL4L7LogicalInterfaceListFromContainer(cont)
	return list, err
}

// func (sm *ServiceManager) CreateRelationcloudRsLIfToCloudSubnet(parentDn, annotation, tDn string) error {
// 	dn := fmt.Sprintf("%s/rsLIfToCloudSubnet", parentDn)
// 	containerJSON := []byte(fmt.Sprintf(`{
// 		"%s": {
// 			"attributes": {
// 				"dn": "%s",
// 				"annotation": "%s",
// 				"tDn": "%s"
// 			}
// 		}
// 	}`, "cloudRsLIfToCloudSubnet", dn, annotation, tDn))

// 	jsonPayload, err := container.ParseJSON(containerJSON)
// 	if err != nil {
// 		return err
// 	}
// 	req, err := sm.client.MakeRestRequest("POST", fmt.Sprintf("%s.json", sm.MOURL), jsonPayload, true)
// 	if err != nil {
// 		return err
// 	}
// 	cont, _, err := sm.client.Do(req)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Printf("%+v", cont)
// 	return nil
// }

// func (sm *ServiceManager) DeleteRelationcloudRsLIfToCloudSubnet(parentDn string) error {
// 	dn := fmt.Sprintf("%s/rsLIfToCloudSubnet", parentDn)
// 	return sm.DeleteByDn(dn, "cloudRsLIfToCloudSubnet")
// }

// func (sm *ServiceManager) ReadRelationcloudRsLIfToCloudSubnet(parentDn string) (interface{}, error) {
// 	dnUrl := fmt.Sprintf("%s/%s/%s.json", models.BaseurlStr, parentDn, "cloudRsLIfToCloudSubnet")
// 	cont, err := sm.GetViaURL(dnUrl)
// 	contList := models.ListFromContainer(cont, "cloudRsLIfToCloudSubnet")

// 	if len(contList) > 0 {
// 		paramMap := make(map[string]string)
// 		paramMap["tDn"] = models.G(contList[0], "tDn")
// 		return paramMap, err
// 	} else {
// 		return nil, err
// 	}
// }
