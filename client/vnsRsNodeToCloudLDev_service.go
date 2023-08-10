package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/v2/models"
)

func (sm *ServiceManager) CreateRelationFromAbsNodeToCloudLDev(function_node string, l4_l7_service_graph_template string, tenant string, description string, vnsRsNodeToCloudLDevAttr models.RelationFromAbsNodeToCloudLDevAttributes) (*models.RelationFromAbsNodeToCloudLDev, error) {
	parentDn := fmt.Sprintf(models.ParentDnVnsRsNodeToCloudLDev, tenant, l4_l7_service_graph_template, function_node)
	vnsRsNodeToCloudLDev := models.NewRelationFromAbsNodeToCloudLDev(models.RnVnsRsNodeToCloudLDev, parentDn, description, vnsRsNodeToCloudLDevAttr)
	err := sm.Save(vnsRsNodeToCloudLDev)
	return vnsRsNodeToCloudLDev, err
}

func (sm *ServiceManager) ReadRelationFromAbsNodeToCloudLDev(function_node string, l4_l7_service_graph_template string, tenant string) (*models.RelationFromAbsNodeToCloudLDev, error) {
	parentDn := fmt.Sprintf(models.ParentDnVnsRsNodeToCloudLDev, tenant, l4_l7_service_graph_template, function_node)
	dn := fmt.Sprintf("%s/%s", parentDn, models.RnVnsRsNodeToCloudLDev)
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}
	vnsRsNodeToCloudLDev := models.RelationFromAbsNodeToCloudLDevFromContainer(cont)
	return vnsRsNodeToCloudLDev, nil
}

func (sm *ServiceManager) DeleteRelationFromAbsNodeToCloudLDev(function_node string, l4_l7_service_graph_template string, tenant string) error {
	parentDn := fmt.Sprintf(models.ParentDnVnsRsNodeToCloudLDev, tenant, l4_l7_service_graph_template, function_node)
	dn := fmt.Sprintf("%s/%s", parentDn, models.RnVnsRsNodeToCloudLDev)
	return sm.DeleteByDn(dn, models.VnsRsNodeToCloudLDevClassName)
}

func (sm *ServiceManager) UpdateRelationFromAbsNodeToCloudLDev(function_node string, l4_l7_service_graph_template string, tenant string, description string, vnsRsNodeToCloudLDevAttr models.RelationFromAbsNodeToCloudLDevAttributes) (*models.RelationFromAbsNodeToCloudLDev, error) {
	parentDn := fmt.Sprintf(models.ParentDnVnsRsNodeToCloudLDev, tenant, l4_l7_service_graph_template, function_node)
	vnsRsNodeToCloudLDev := models.NewRelationFromAbsNodeToCloudLDev(models.RnVnsRsNodeToCloudLDev, parentDn, description, vnsRsNodeToCloudLDevAttr)
	vnsRsNodeToCloudLDev.Status = "modified"
	err := sm.Save(vnsRsNodeToCloudLDev)
	return vnsRsNodeToCloudLDev, err
}

func (sm *ServiceManager) ListRelationFromAbsNodeToCloudLDev(function_node string, l4_l7_service_graph_template string, tenant string) ([]*models.RelationFromAbsNodeToCloudLDev, error) {
	parentDn := fmt.Sprintf(models.ParentDnVnsRsNodeToCloudLDev, tenant, l4_l7_service_graph_template, function_node)
	dnUrl := fmt.Sprintf("%s/%s/%s.json", models.BaseurlStr, parentDn, models.VnsRsNodeToCloudLDevClassName)
	cont, err := sm.GetViaURL(dnUrl)
	list := models.RelationFromAbsNodeToCloudLDevListFromContainer(cont)
	return list, err
}
