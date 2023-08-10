package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/v2/models"
)

func (sm *ServiceManager) CreateRelationFromServiceGraphConnectionToServiceGraphConnectors(tDn string, connection string, l4_l7_service_graph_template string, tenant string, description string, vnsRsAbsConnectionConnsAttr models.RelationFromServiceGraphConnectionToServiceGraphConnectorsAttributes) (*models.RelationFromServiceGraphConnectionToServiceGraphConnectors, error) {
	rn := fmt.Sprintf(models.RnVnsRsAbsConnectionConns, tDn)
	parentDn := fmt.Sprintf(models.ParentDnVnsRsAbsConnectionConns, tenant, l4_l7_service_graph_template, connection)
	vnsRsAbsConnectionConns := models.NewRelationFromServiceGraphConnectionToServiceGraphConnectors(rn, parentDn, description, vnsRsAbsConnectionConnsAttr)
	err := sm.Save(vnsRsAbsConnectionConns)
	return vnsRsAbsConnectionConns, err
}

func (sm *ServiceManager) ReadRelationFromServiceGraphConnectionToServiceGraphConnectors(tDn string, connection string, l4_l7_service_graph_template string, tenant string) (*models.RelationFromServiceGraphConnectionToServiceGraphConnectors, error) {
	rn := fmt.Sprintf(models.RnVnsRsAbsConnectionConns, tDn)
	parentDn := fmt.Sprintf(models.ParentDnVnsRsAbsConnectionConns, tenant, l4_l7_service_graph_template, connection)
	dn := fmt.Sprintf("%s/%s", parentDn, rn)
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}
	vnsRsAbsConnectionConns := models.RelationFromServiceGraphConnectionToServiceGraphConnectorsFromContainer(cont)
	return vnsRsAbsConnectionConns, nil
}

func (sm *ServiceManager) DeleteRelationFromServiceGraphConnectionToServiceGraphConnectors(tDn string, connection string, l4_l7_service_graph_template string, tenant string) error {
	rn := fmt.Sprintf(models.RnVnsRsAbsConnectionConns, tDn)
	parentDn := fmt.Sprintf(models.ParentDnVnsRsAbsConnectionConns, tenant, l4_l7_service_graph_template, connection)
	dn := fmt.Sprintf("%s/%s", parentDn, rn)
	return sm.DeleteByDn(dn, models.VnsRsAbsConnectionConnsClassName)
}

func (sm *ServiceManager) UpdateRelationFromServiceGraphConnectionToServiceGraphConnectors(tDn string, connection string, l4_l7_service_graph_template string, tenant string, description string, vnsRsAbsConnectionConnsAttr models.RelationFromServiceGraphConnectionToServiceGraphConnectorsAttributes) (*models.RelationFromServiceGraphConnectionToServiceGraphConnectors, error) {
	rn := fmt.Sprintf(models.RnVnsRsAbsConnectionConns, tDn)
	parentDn := fmt.Sprintf(models.ParentDnVnsRsAbsConnectionConns, tenant, l4_l7_service_graph_template, connection)
	vnsRsAbsConnectionConns := models.NewRelationFromServiceGraphConnectionToServiceGraphConnectors(rn, parentDn, description, vnsRsAbsConnectionConnsAttr)
	vnsRsAbsConnectionConns.Status = "modified"
	err := sm.Save(vnsRsAbsConnectionConns)
	return vnsRsAbsConnectionConns, err
}

func (sm *ServiceManager) ListRelationFromServiceGraphConnectionToServiceGraphConnectors(connection string, l4_l7_service_graph_template string, tenant string) ([]*models.RelationFromServiceGraphConnectionToServiceGraphConnectors, error) {
	parentDn := fmt.Sprintf(models.ParentDnVnsRsAbsConnectionConns, tenant, l4_l7_service_graph_template, connection)
	dnUrl := fmt.Sprintf("%s/%s/%s.json", models.BaseurlStr, parentDn, models.VnsRsAbsConnectionConnsClassName)
	cont, err := sm.GetViaURL(dnUrl)
	list := models.RelationFromServiceGraphConnectionToServiceGraphConnectorsListFromContainer(cont)
	return list, err
}
