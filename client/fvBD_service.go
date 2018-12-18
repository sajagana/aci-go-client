package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/ciscoecosystem/aci-go-client/container"

)









func (sm *ServiceManager) CreateBridgeDomain(name string ,tenant string  ,description string, fvBDattr models.BridgeDomainAttributes) (*models.BridgeDomain, error) {	
	rn := fmt.Sprintf("BD-%s",name)
	parentDn := fmt.Sprintf("uni/tn-%s", tenant )
	fvBD := models.NewBridgeDomain(rn, parentDn, description, fvBDattr)
	err := sm.Save(fvBD)
	return fvBD, err
}

func (sm *ServiceManager) ReadBridgeDomain(name string ,tenant string ) (*models.BridgeDomain, error) {
	dn := fmt.Sprintf("uni/tn-%s/BD-%s", tenant ,name )    
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	fvBD := models.BridgeDomainFromContainer(cont)
	return fvBD, nil
}

func (sm *ServiceManager) DeleteBridgeDomain(name string ,tenant string ) error {
	dn := fmt.Sprintf("uni/tn-%s/BD-%s", tenant ,name )
	return sm.DeleteByDn(dn, models.FvbdClassName)
}

func (sm *ServiceManager) UpdateBridgeDomain(name string ,tenant string  ,description string, fvBDattr models.BridgeDomainAttributes) (*models.BridgeDomain, error) {
	rn := fmt.Sprintf("BD-%s",name)
	parentDn := fmt.Sprintf("uni/tn-%s", tenant )
	fvBD := models.NewBridgeDomain(rn, parentDn, description, fvBDattr)

    fvBD.Status = "modified"
	err := sm.Save(fvBD)
	return fvBD, err

}

func (sm *ServiceManager) ListBridgeDomain(tenant string ) ([]*models.BridgeDomain, error) {

	baseurlStr := "/api/node/class"	
	dnUrl := fmt.Sprintf("%s/uni/tn-%s/fvBD.json", baseurlStr , tenant )
    
    cont, err := sm.GetViaURL(dnUrl)
	list := models.BridgeDomainListFromContainer(cont)

	return list, err
}

func (sm *ServiceManager) CreateRelationfvRsBDToProfileFromBridgeDomain( parentDn, tnRtctrlProfileName string) error {
	dn := fmt.Sprintf("%s/rsBDToProfile", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnRtctrlProfileName": "%s"
								
			}
		}
	}`, "fvRsBDToProfile", dn,tnRtctrlProfileName))

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

func (sm *ServiceManager) DeleteRelationfvRsBDToProfileFromBridgeDomain(parentDn string) error{
	dn := fmt.Sprintf("%s/rsBDToProfile", parentDn)
	return sm.DeleteByDn(dn , "fvRsBDToProfile")
}
func (sm *ServiceManager) CreateRelationfvRsBDToRelayPFromBridgeDomain( parentDn, tnDhcpRelayPName string) error {
	dn := fmt.Sprintf("%s/rsBDToRelayP", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnDhcpRelayPName": "%s"
								
			}
		}
	}`, "fvRsBDToRelayP", dn,tnDhcpRelayPName))

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

func (sm *ServiceManager) DeleteRelationfvRsBDToRelayPFromBridgeDomain(parentDn string) error{
	dn := fmt.Sprintf("%s/rsBDToRelayP", parentDn)
	return sm.DeleteByDn(dn , "fvRsBDToRelayP")
}
func (sm *ServiceManager) CreateRelationfvRsABDPolMonPolFromBridgeDomain( parentDn, tnMonEPGPolName string) error {
	dn := fmt.Sprintf("%s/rsABDPolMonPol", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnMonEPGPolName": "%s"
								
			}
		}
	}`, "fvRsABDPolMonPol", dn,tnMonEPGPolName))

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

func (sm *ServiceManager) DeleteRelationfvRsABDPolMonPolFromBridgeDomain(parentDn string) error{
	dn := fmt.Sprintf("%s/rsABDPolMonPol", parentDn)
	return sm.DeleteByDn(dn , "fvRsABDPolMonPol")
}
func (sm *ServiceManager) CreateRelationfvRsBDToNdPFromBridgeDomain( parentDn, tnNdIfPolName string) error {
	dn := fmt.Sprintf("%s/rsBDToNdP", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnNdIfPolName": "%s"
								
			}
		}
	}`, "fvRsBDToNdP", dn,tnNdIfPolName))

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
func (sm *ServiceManager) CreateRelationfvRsBdFloodToFromBridgeDomain( parentDn, tDn string) error {
	dn := fmt.Sprintf("%s/rsbdFloodTo-[%s]", parentDn, tDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s"				
			}
		}
	}`, "fvRsBdFloodTo", dn))

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

func (sm *ServiceManager) DeleteRelationfvRsBdFloodToFromBridgeDomain(parentDn , tDn string) error{
	dn := fmt.Sprintf("%s/rsbdFloodTo-[%s]", parentDn, tDn)
	return sm.DeleteByDn(dn , "fvRsBdFloodTo")
}
func (sm *ServiceManager) CreateRelationfvRsBDToFhsFromBridgeDomain( parentDn, tnFhsBDPolName string) error {
	dn := fmt.Sprintf("%s/rsBDToFhs", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnFhsBDPolName": "%s"
								
			}
		}
	}`, "fvRsBDToFhs", dn,tnFhsBDPolName))

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

func (sm *ServiceManager) DeleteRelationfvRsBDToFhsFromBridgeDomain(parentDn string) error{
	dn := fmt.Sprintf("%s/rsBDToFhs", parentDn)
	return sm.DeleteByDn(dn , "fvRsBDToFhs")
}
func (sm *ServiceManager) CreateRelationfvRsCtxFromBridgeDomain( parentDn, tnFvCtxName string) error {
	dn := fmt.Sprintf("%s/rsctx", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnFvCtxName": "%s"
								
			}
		}
	}`, "fvRsCtx", dn,tnFvCtxName))

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
func (sm *ServiceManager) CreateRelationfvRsBDToNetflowMonitorPolFromBridgeDomain( parentDn, tnNetflowMonitorPolName,fltType string) error {
	dn := fmt.Sprintf("%s/rsBDToNetflowMonitorPol-[%s]-%s", parentDn, tnNetflowMonitorPolName,fltType)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s"				
			}
		}
	}`, "fvRsBDToNetflowMonitorPol", dn))

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

func (sm *ServiceManager) DeleteRelationfvRsBDToNetflowMonitorPolFromBridgeDomain(parentDn , tnNetflowMonitorPolName,fltType string) error{
	dn := fmt.Sprintf("%s/rsBDToNetflowMonitorPol-[%s]-%s", parentDn, tnNetflowMonitorPolName,fltType)
	return sm.DeleteByDn(dn , "fvRsBDToNetflowMonitorPol")
}
func (sm *ServiceManager) CreateRelationfvRsIgmpsnFromBridgeDomain( parentDn, tnIgmpSnoopPolName string) error {
	dn := fmt.Sprintf("%s/rsigmpsn", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnIgmpSnoopPolName": "%s"
								
			}
		}
	}`, "fvRsIgmpsn", dn,tnIgmpSnoopPolName))

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
func (sm *ServiceManager) CreateRelationfvRsBdToEpRetFromBridgeDomain( parentDn, tnFvEpRetPolName string) error {
	dn := fmt.Sprintf("%s/rsbdToEpRet", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnFvEpRetPolName": "%s"
								
			}
		}
	}`, "fvRsBdToEpRet", dn,tnFvEpRetPolName))

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
func (sm *ServiceManager) CreateRelationfvRsBDToOutFromBridgeDomain( parentDn, tnL3extOutName string) error {
	dn := fmt.Sprintf("%s/rsBDToOut-%s", parentDn, tnL3extOutName)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s"				
			}
		}
	}`, "fvRsBDToOut", dn))

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

func (sm *ServiceManager) DeleteRelationfvRsBDToOutFromBridgeDomain(parentDn , tnL3extOutName string) error{
	dn := fmt.Sprintf("%s/rsBDToOut-%s", parentDn, tnL3extOutName)
	return sm.DeleteByDn(dn , "fvRsBDToOut")
}

