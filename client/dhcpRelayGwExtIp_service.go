package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/v2/models"
)

func (sm *ServiceManager) CreateUsetheexternalsecondaryaddressforDHCPrelaygateway(parentDn string, description string, dhcpRelayGwExtIpAttr models.UsetheexternalsecondaryaddressforDHCPrelaygatewayAttributes) (*models.UsetheexternalsecondaryaddressforDHCPrelaygateway, error) {
	rn := fmt.Sprintf(models.RndhcpRelayGwExtIp)
	dhcpRelayGwExtIp := models.NewUsetheexternalsecondaryaddressforDHCPrelaygateway(rn, parentDn, description, dhcpRelayGwExtIpAttr)
	err := sm.Save(dhcpRelayGwExtIp)
	return dhcpRelayGwExtIp, err
}

func (sm *ServiceManager) ReadUsetheexternalsecondaryaddressforDHCPrelaygateway(parentDn string) (*models.UsetheexternalsecondaryaddressforDHCPrelaygateway, error) {
	rn := fmt.Sprintf(models.RndhcpRelayGwExtIp)
	dn := fmt.Sprintf("%s/%s", parentDn, rn)

	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	dhcpRelayGwExtIp := models.UsetheexternalsecondaryaddressforDHCPrelaygatewayFromContainer(cont)
	return dhcpRelayGwExtIp, nil
}

func (sm *ServiceManager) DeleteUsetheexternalsecondaryaddressforDHCPrelaygateway(parentDn string) error {
	rn := fmt.Sprintf(models.RndhcpRelayGwExtIp)
	dn := fmt.Sprintf("%s/%s", parentDn, rn)
	return sm.DeleteByDn(dn, models.DhcprelaygwextipClassName)
}

func (sm *ServiceManager) UpdateUsetheexternalsecondaryaddressforDHCPrelaygateway(parentDn string, description string, dhcpRelayGwExtIpAttr models.UsetheexternalsecondaryaddressforDHCPrelaygatewayAttributes) (*models.UsetheexternalsecondaryaddressforDHCPrelaygateway, error) {
	rn := fmt.Sprintf(models.RndhcpRelayGwExtIp)
	dhcpRelayGwExtIp := models.NewUsetheexternalsecondaryaddressforDHCPrelaygateway(rn, parentDn, description, dhcpRelayGwExtIpAttr)
	dhcpRelayGwExtIp.Status = "modified"
	err := sm.Save(dhcpRelayGwExtIp)
	return dhcpRelayGwExtIp, err
}

func (sm *ServiceManager) ListUsetheexternalsecondaryaddressforDHCPrelaygateway(parentDn string) ([]*models.UsetheexternalsecondaryaddressforDHCPrelaygateway, error) {
	dnUrl := fmt.Sprintf("%s/%s/dhcpRelayGwExtIp.json", models.BaseurlStr, parentDn)
	cont, err := sm.GetViaURL(dnUrl)
	list := models.UsetheexternalsecondaryaddressforDHCPrelaygatewayListFromContainer(cont)
	return list, err
}
