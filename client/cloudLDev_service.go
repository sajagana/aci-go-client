package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/v2/container"
	"github.com/ciscoecosystem/aci-go-client/v2/models"
)

func (sm *ServiceManager) CreateCloudL4L7Device(name string, tenant string, description string, cloudLDevAttr models.CloudL4L7DeviceAttributes) (*models.CloudL4L7Device, error) {
	rn := fmt.Sprintf(models.RnCloudLDev, name)
	parentDn := fmt.Sprintf(models.ParentDnCloudLDev, tenant)
	cloudLDev := models.NewCloudL4L7Device(rn, parentDn, description, cloudLDevAttr)
	err := sm.Save(cloudLDev)
	return cloudLDev, err
}

func (sm *ServiceManager) ReadCloudL4L7Device(name string, tenant string) (*models.CloudL4L7Device, error) {
	rn := fmt.Sprintf(models.RnCloudLDev, name)
	parentDn := fmt.Sprintf(models.ParentDnCloudLDev, tenant)
	dn := fmt.Sprintf("%s/%s", parentDn, rn)
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}
	cloudLDev := models.CloudL4L7DeviceFromContainer(cont)
	return cloudLDev, nil
}

func (sm *ServiceManager) DeleteCloudL4L7Device(name string, tenant string) error {
	rn := fmt.Sprintf(models.RnCloudLDev, name)
	parentDn := fmt.Sprintf(models.ParentDnCloudLDev, tenant)
	dn := fmt.Sprintf("%s/%s", parentDn, rn)
	return sm.DeleteByDn(dn, models.CloudLDevClassName)
}

func (sm *ServiceManager) UpdateCloudL4L7Device(name string, tenant string, description string, cloudLDevAttr models.CloudL4L7DeviceAttributes) (*models.CloudL4L7Device, error) {
	rn := fmt.Sprintf(models.RnCloudLDev, name)
	parentDn := fmt.Sprintf(models.ParentDnCloudLDev, tenant)
	cloudLDev := models.NewCloudL4L7Device(rn, parentDn, description, cloudLDevAttr)
	cloudLDev.Status = "modified"
	err := sm.Save(cloudLDev)
	return cloudLDev, err
}

func (sm *ServiceManager) ListCloudL4L7Device(tenant string) ([]*models.CloudL4L7Device, error) {
	parentDn := fmt.Sprintf(models.ParentDnCloudLDev, tenant)
	dnUrl := fmt.Sprintf("%s/%s/%s.json", models.BaseurlStr, parentDn, models.CloudLDevClassName)
	cont, err := sm.GetViaURL(dnUrl)
	list := models.CloudL4L7DeviceListFromContainer(cont)
	return list, err
}

// cloudRsLDevToCtx - Bind VRF with Cloud Device
func (sm *ServiceManager) CreateOrUpdateRelationcloudRsLDevToCtx(parentDn, annotation, tDn string) error {
	dn := fmt.Sprintf("%s/rsLDevToCtx", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s",
				"annotation": "%s",
				"tDn": "%s"
			}
		}
	}`, "cloudRsLDevToCtx", dn, annotation, tDn))

	jsonPayload, err := container.ParseJSON(containerJSON)
	if err != nil {
		return err
	}
	req, err := sm.client.MakeRestRequest("POST", fmt.Sprintf("%s.json", sm.MOURL), jsonPayload, true)
	if err != nil {
		return err
	}
	cont, _, err := sm.client.Do(req)
	if err != nil {
		return err
	}
	fmt.Printf("%+v", cont)
	return nil
}

func (sm *ServiceManager) DeleteRelationcloudRsLDevToCtx(parentDn string) error {
	dn := fmt.Sprintf("%s/rsLDevToCtx", parentDn)
	return sm.DeleteByDn(dn, "cloudRsLDevToCtx")
}

func (sm *ServiceManager) ReadRelationcloudRsLDevToCtx(parentDn string) (interface{}, error) {
	dnUrl := fmt.Sprintf("%s/%s/%s.json", models.BaseurlStr, parentDn, "cloudRsLDevToCtx")
	cont, err := sm.GetViaURL(dnUrl)
	contList := models.ListFromContainer(cont, "cloudRsLDevToCtx")

	if len(contList) > 0 {
		paramMap := make(map[string]string)
		paramMap["tDn"] = models.G(contList[0], "tDn")
		return paramMap, err
	} else {
		return nil, err
	}
}

// func (sm *ServiceManager) CreateRelationcloudRsLDevToCloudSubnet(parentDn, annotation, tDn string) error {
// 	dn := fmt.Sprintf("%s/rsLDevToCloudSubnet-[%s]", parentDn, tDn)
// 	containerJSON := []byte(fmt.Sprintf(`{
// 		"%s": {
// 			"attributes": {
// 				"dn": "%s",
// 				"annotation": "%s",
// 				"tDn": "%s"
// 			}
// 		}
// 	}`, "cloudRsLDevToCloudSubnet", dn, annotation, tDn))

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

// func (sm *ServiceManager) DeleteRelationcloudRsLDevToCloudSubnet(parentDn, tDn string) error {
// 	dn := fmt.Sprintf("%s/rsLDevToCloudSubnet-[%s]", parentDn, tDn)
// 	return sm.DeleteByDn(dn, "cloudRsLDevToCloudSubnet")
// }

// func (sm *ServiceManager) ReadRelationcloudRsLDevToCloudSubnet(parentDn string) (interface{}, error) {
// 	dnUrl := fmt.Sprintf("%s/%s/%s.json", models.BaseurlStr, parentDn, "cloudRsLDevToCloudSubnet")
// 	cont, err := sm.GetViaURL(dnUrl)
// 	contList := models.ListFromContainer(cont, "cloudRsLDevToCloudSubnet")

// 	st := &schema.Set{
// 		F: schema.HashString,
// 	}
// 	for _, contItem := range contList {
// 		paramMap := make(map[string]string)
// 		paramMap["tDn"] = models.G(contItem, "tDn")
// 		st.Add(paramMap)
// 	}
// 	return st, err
// }

// func (sm *ServiceManager) CreateRelationcloudRsLDevToComputePol(parentDn, annotation, tDn string) error {
// 	dn := fmt.Sprintf("%s/rsLDevToComputePol", parentDn)
// 	containerJSON := []byte(fmt.Sprintf(`{
// 		"%s": {
// 			"attributes": {
// 				"dn": "%s",
// 				"annotation": "%s",
// 				"tDn": "%s"
// 			}
// 		}
// 	}`, "cloudRsLDevToComputePol", dn, annotation, tDn))

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

// func (sm *ServiceManager) DeleteRelationcloudRsLDevToComputePol(parentDn string) error {
// 	dn := fmt.Sprintf("%s/rsLDevToComputePol", parentDn)
// 	return sm.DeleteByDn(dn, "cloudRsLDevToComputePol")
// }

// func (sm *ServiceManager) ReadRelationcloudRsLDevToComputePol(parentDn string) (interface{}, error) {
// 	dnUrl := fmt.Sprintf("%s/%s/%s.json", models.BaseurlStr, parentDn, "cloudRsLDevToComputePol")
// 	cont, err := sm.GetViaURL(dnUrl)
// 	contList := models.ListFromContainer(cont, "cloudRsLDevToComputePol")

// 	if len(contList) > 0 {
// 		paramMap := make(map[string]string)
// 		paramMap["tDn"] = models.G(contList[0], "tDn")
// 		return paramMap, err
// 	} else {
// 		return nil, err
// 	}
// }

// func (sm *ServiceManager) CreateRelationcloudRsLDevToMgmtPol(parentDn, annotation, tDn string) error {
// 	dn := fmt.Sprintf("%s/rsLDevToMgmtPol", parentDn)
// 	containerJSON := []byte(fmt.Sprintf(`{
// 		"%s": {
// 			"attributes": {
// 				"dn": "%s",
// 				"annotation": "%s",
// 				"tDn": "%s"
// 			}
// 		}
// 	}`, "cloudRsLDevToMgmtPol", dn, annotation, tDn))

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

// func (sm *ServiceManager) DeleteRelationcloudRsLDevToMgmtPol(parentDn string) error {
// 	dn := fmt.Sprintf("%s/rsLDevToMgmtPol", parentDn)
// 	return sm.DeleteByDn(dn, "cloudRsLDevToMgmtPol")
// }

// func (sm *ServiceManager) ReadRelationcloudRsLDevToMgmtPol(parentDn string) (interface{}, error) {
// 	dnUrl := fmt.Sprintf("%s/%s/%s.json", models.BaseurlStr, parentDn, "cloudRsLDevToMgmtPol")
// 	cont, err := sm.GetViaURL(dnUrl)
// 	contList := models.ListFromContainer(cont, "cloudRsLDevToMgmtPol")

// 	if len(contList) > 0 {
// 		paramMap := make(map[string]string)
// 		paramMap["tDn"] = models.G(contList[0], "tDn")
// 		return paramMap, err
// 	} else {
// 		return nil, err
// 	}
// }

// func (sm *ServiceManager) CreateRelationvnsRsALDevToDevMgr(parentDn, annotation, tnVnsDevMgrName string) error {
// 	dn := fmt.Sprintf("%s/rsaLDevToDevMgr", parentDn)
// 	containerJSON := []byte(fmt.Sprintf(`{
// 		"%s": {
// 			"attributes": {
// 				"dn": "%s",
// 				"annotation": "%s",
// 				"tnVnsDevMgrName": "%s"
// 			}
// 		}
// 	}`, "vnsRsALDevToDevMgr", dn, annotation, tnVnsDevMgrName))

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

// func (sm *ServiceManager) DeleteRelationvnsRsALDevToDevMgr(parentDn string) error {
// 	dn := fmt.Sprintf("%s/rsaLDevToDevMgr", parentDn)
// 	return sm.DeleteByDn(dn, "vnsRsALDevToDevMgr")
// }

// func (sm *ServiceManager) ReadRelationvnsRsALDevToDevMgr(parentDn string) (interface{}, error) {
// 	dnUrl := fmt.Sprintf("%s/%s/%s.json", models.BaseurlStr, parentDn, "vnsRsALDevToDevMgr")
// 	cont, err := sm.GetViaURL(dnUrl)
// 	contList := models.ListFromContainer(cont, "vnsRsALDevToDevMgr")

// 	if len(contList) > 0 {
// 		paramMap := make(map[string]string)
// 		paramMap["tnVnsDevMgrName"] = models.G(contList[0], "tnVnsDevMgrName")
// 		paramMap["tDn"] = models.G(contList[0], "tDn")
// 		return paramMap, err
// 	} else {
// 		return nil, err
// 	}
// }

// func (sm *ServiceManager) CreateRelationvnsRsALDevToDomP(parentDn, annotation, switchingMode string, tDn string) error {
// 	dn := fmt.Sprintf("%s/rsALDevToDomP", parentDn)
// 	containerJSON := []byte(fmt.Sprintf(`{
// 		"%s": {
// 			"attributes": {
// 				"dn": "%s",
// 				"annotation": "%s",
// 				"tDn": "%s"
// 			}
// 		}
// 	}`, "vnsRsALDevToDomP", dn, annotation, tDn))

// 	attributes := map[string]interface{}{
// 		"switchingMode": switchingMode,
// 	}
// 	var output map[string]interface{}
// 	err_output := json.Unmarshal([]byte(containerJSON), &output)
// 	if err_output != nil {
// 		return err_output
// 	}
// 	for _, mo := range output {
// 		if mo_map, ok := mo.(map[string]interface{}); ok {
// 			for _, mo_attributes := range mo_map {
// 				if mo_attributes_map, ok := mo_attributes.(map[string]interface{}); ok {
// 					for key, value := range attributes {
// 						if value != "" {
// 							mo_attributes_map[key] = value
// 						}
// 					}
// 				}
// 			}
// 		}

// 	}
// 	input, out_err := json.Marshal(output)
// 	if out_err != nil {
// 		return out_err
// 	}
// 	jsonPayload, err := container.ParseJSON(input)
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

// func (sm *ServiceManager) DeleteRelationvnsRsALDevToDomP(parentDn string) error {
// 	dn := fmt.Sprintf("%s/rsALDevToDomP", parentDn)
// 	return sm.DeleteByDn(dn, "vnsRsALDevToDomP")
// }

// func (sm *ServiceManager) ReadRelationvnsRsALDevToDomP(parentDn string) (interface{}, error) {
// 	dnUrl := fmt.Sprintf("%s/%s/%s.json", models.BaseurlStr, parentDn, "vnsRsALDevToDomP")
// 	cont, err := sm.GetViaURL(dnUrl)
// 	contList := models.ListFromContainer(cont, "vnsRsALDevToDomP")

// 	if len(contList) > 0 {
// 		paramMap := make(map[string]string)
// 		paramMap["tDn"] = models.G(contList[0], "tDn")
// 		paramMap["switchingMode"] = models.G(contList[0], "switchingMode")
// 		return paramMap, err
// 	} else {
// 		return nil, err
// 	}
// }

// func (sm *ServiceManager) CreateRelationvnsRsALDevToPhysDomP(parentDn, annotation, tDn string) error {
// 	dn := fmt.Sprintf("%s/rsALDevToPhysDomP", parentDn)
// 	containerJSON := []byte(fmt.Sprintf(`{
// 		"%s": {
// 			"attributes": {
// 				"dn": "%s",
// 				"annotation": "%s",
// 				"tDn": "%s"
// 			}
// 		}
// 	}`, "vnsRsALDevToPhysDomP", dn, annotation, tDn))

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

// func (sm *ServiceManager) DeleteRelationvnsRsALDevToPhysDomP(parentDn string) error {
// 	dn := fmt.Sprintf("%s/rsALDevToPhysDomP", parentDn)
// 	return sm.DeleteByDn(dn, "vnsRsALDevToPhysDomP")
// }

// func (sm *ServiceManager) ReadRelationvnsRsALDevToPhysDomP(parentDn string) (interface{}, error) {
// 	dnUrl := fmt.Sprintf("%s/%s/%s.json", models.BaseurlStr, parentDn, "vnsRsALDevToPhysDomP")
// 	cont, err := sm.GetViaURL(dnUrl)
// 	contList := models.ListFromContainer(cont, "vnsRsALDevToPhysDomP")

// 	if len(contList) > 0 {
// 		paramMap := make(map[string]string)
// 		paramMap["tDn"] = models.G(contList[0], "tDn")
// 		return paramMap, err
// 	} else {
// 		return nil, err
// 	}
// }

// func (sm *ServiceManager) CreateRelationvnsRsALDevToVlanInstP(parentDn, annotation, tDn string) error {
// 	dn := fmt.Sprintf("%s/rsALDevToVlanInstP-[%s]", parentDn, tDn)
// 	containerJSON := []byte(fmt.Sprintf(`{
// 		"%s": {
// 			"attributes": {
// 				"dn": "%s",
// 				"annotation": "%s",
// 				"tDn": "%s"
// 			}
// 		}
// 	}`, "vnsRsALDevToVlanInstP", dn, annotation, tDn))

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

// func (sm *ServiceManager) DeleteRelationvnsRsALDevToVlanInstP(parentDn, tDn string) error {
// 	dn := fmt.Sprintf("%s/rsALDevToVlanInstP-[%s]", parentDn, tDn)
// 	return sm.DeleteByDn(dn, "vnsRsALDevToVlanInstP")
// }

// func (sm *ServiceManager) ReadRelationvnsRsALDevToVlanInstP(parentDn string) (interface{}, error) {
// 	dnUrl := fmt.Sprintf("%s/%s/%s.json", models.BaseurlStr, parentDn, "vnsRsALDevToVlanInstP")
// 	cont, err := sm.GetViaURL(dnUrl)
// 	contList := models.ListFromContainer(cont, "vnsRsALDevToVlanInstP")

// 	st := &schema.Set{
// 		F: schema.HashString,
// 	}
// 	for _, contItem := range contList {
// 		paramMap := make(map[string]string)
// 		paramMap["tDn"] = models.G(contItem, "tDn")
// 		st.Add(paramMap)
// 	}
// 	return st, err
// }

// func (sm *ServiceManager) CreateRelationvnsRsALDevToVxlanInstP(parentDn, annotation, tDn string) error {
// 	dn := fmt.Sprintf("%s/rsALDevToVxlanInstP", parentDn)
// 	containerJSON := []byte(fmt.Sprintf(`{
// 		"%s": {
// 			"attributes": {
// 				"dn": "%s",
// 				"annotation": "%s",
// 				"tDn": "%s"
// 			}
// 		}
// 	}`, "vnsRsALDevToVxlanInstP", dn, annotation, tDn))

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

// func (sm *ServiceManager) DeleteRelationvnsRsALDevToVxlanInstP(parentDn string) error {
// 	dn := fmt.Sprintf("%s/rsALDevToVxlanInstP", parentDn)
// 	return sm.DeleteByDn(dn, "vnsRsALDevToVxlanInstP")
// }

// func (sm *ServiceManager) ReadRelationvnsRsALDevToVxlanInstP(parentDn string) (interface{}, error) {
// 	dnUrl := fmt.Sprintf("%s/%s/%s.json", models.BaseurlStr, parentDn, "vnsRsALDevToVxlanInstP")
// 	cont, err := sm.GetViaURL(dnUrl)
// 	contList := models.ListFromContainer(cont, "vnsRsALDevToVxlanInstP")

// 	if len(contList) > 0 {
// 		paramMap := make(map[string]string)
// 		paramMap["tDn"] = models.G(contList[0], "tDn")
// 		return paramMap, err
// 	} else {
// 		return nil, err
// 	}
// }

// func (sm *ServiceManager) CreateRelationvnsRsDevEpg(parentDn, annotation, tDn string) error {
// 	dn := fmt.Sprintf("%s/rsdevEpg", parentDn)
// 	containerJSON := []byte(fmt.Sprintf(`{
// 		"%s": {
// 			"attributes": {
// 				"dn": "%s",
// 				"annotation": "%s",
// 				"tDn": "%s"
// 			}
// 		}
// 	}`, "vnsRsDevEpg", dn, annotation, tDn))

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

// func (sm *ServiceManager) DeleteRelationvnsRsDevEpg(parentDn string) error {
// 	dn := fmt.Sprintf("%s/rsdevEpg", parentDn)
// 	return sm.DeleteByDn(dn, "vnsRsDevEpg")
// }

// func (sm *ServiceManager) ReadRelationvnsRsDevEpg(parentDn string) (interface{}, error) {
// 	dnUrl := fmt.Sprintf("%s/%s/%s.json", models.BaseurlStr, parentDn, "vnsRsDevEpg")
// 	cont, err := sm.GetViaURL(dnUrl)
// 	contList := models.ListFromContainer(cont, "vnsRsDevEpg")

// 	if len(contList) > 0 {
// 		paramMap := make(map[string]string)
// 		paramMap["tDn"] = models.G(contList[0], "tDn")
// 		return paramMap, err
// 	} else {
// 		return nil, err
// 	}
// }

// func (sm *ServiceManager) CreateRelationvnsRsMDevAtt(parentDn, annotation, tDn string) error {
// 	dn := fmt.Sprintf("%s/rsmDevAtt", parentDn)
// 	containerJSON := []byte(fmt.Sprintf(`{
// 		"%s": {
// 			"attributes": {
// 				"dn": "%s",
// 				"annotation": "%s",
// 				"tDn": "%s"
// 			}
// 		}
// 	}`, "vnsRsMDevAtt", dn, annotation, tDn))

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

// func (sm *ServiceManager) DeleteRelationvnsRsMDevAtt(parentDn string) error {
// 	dn := fmt.Sprintf("%s/rsmDevAtt", parentDn)
// 	return sm.DeleteByDn(dn, "vnsRsMDevAtt")
// }

// func (sm *ServiceManager) ReadRelationvnsRsMDevAtt(parentDn string) (interface{}, error) {
// 	dnUrl := fmt.Sprintf("%s/%s/%s.json", models.BaseurlStr, parentDn, "vnsRsMDevAtt")
// 	cont, err := sm.GetViaURL(dnUrl)
// 	contList := models.ListFromContainer(cont, "vnsRsMDevAtt")

// 	if len(contList) > 0 {
// 		paramMap := make(map[string]string)
// 		paramMap["tDn"] = models.G(contList[0], "tDn")
// 		return paramMap, err
// 	} else {
// 		return nil, err
// 	}
// }

// func (sm *ServiceManager) CreateRelationvnsRsSvcMgmtEpg(parentDn, annotation, tDn string) error {
// 	dn := fmt.Sprintf("%s/rssvcMgmtEpg", parentDn)
// 	containerJSON := []byte(fmt.Sprintf(`{
// 		"%s": {
// 			"attributes": {
// 				"dn": "%s",
// 				"annotation": "%s",
// 				"tDn": "%s"
// 			}
// 		}
// 	}`, "vnsRsSvcMgmtEpg", dn, annotation, tDn))

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

// func (sm *ServiceManager) DeleteRelationvnsRsSvcMgmtEpg(parentDn string) error {
// 	dn := fmt.Sprintf("%s/rssvcMgmtEpg", parentDn)
// 	return sm.DeleteByDn(dn, "vnsRsSvcMgmtEpg")
// }

// func (sm *ServiceManager) ReadRelationvnsRsSvcMgmtEpg(parentDn string) (interface{}, error) {
// 	dnUrl := fmt.Sprintf("%s/%s/%s.json", models.BaseurlStr, parentDn, "vnsRsSvcMgmtEpg")
// 	cont, err := sm.GetViaURL(dnUrl)
// 	contList := models.ListFromContainer(cont, "vnsRsSvcMgmtEpg")

// 	if len(contList) > 0 {
// 		paramMap := make(map[string]string)
// 		paramMap["tDn"] = models.G(contList[0], "tDn")
// 		return paramMap, err
// 	} else {
// 		return nil, err
// 	}
// }
