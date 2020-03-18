package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"
)

func (sm *ServiceManager) CreateNodeBlockMG(name string, pod_maintenance_group string, description string, fabricNodeBlkattr models.NodeBlockAttributes) (*models.NodeBlock, error) {
	rn := fmt.Sprintf("nodeblk-%s", name)
	parentDn := fmt.Sprintf("uni/fabric/maintgrp-%s", pod_maintenance_group)
	fabricNodeBlk := models.NewNodeBlock(rn, parentDn, description, fabricNodeBlkattr)
	err := sm.Save(fabricNodeBlk)
	return fabricNodeBlk, err
}

func (sm *ServiceManager) ReadNodeBlockMG(name string, pod_maintenance_group string) (*models.NodeBlock, error) {
	dn := fmt.Sprintf("uni/fabric/maintgrp-%s/nodeblk-%s", pod_maintenance_group, name)
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	fabricNodeBlk := models.NodeBlockFromContainer(cont)
	return fabricNodeBlk, nil
}

func (sm *ServiceManager) DeleteNodeBlockMG(name string, pod_maintenance_group string) error {
	dn := fmt.Sprintf("uni/fabric/maintgrp-%s/nodeblk-%s", pod_maintenance_group, name)
	return sm.DeleteByDn(dn, models.FabricnodeblkClassName)
}

func (sm *ServiceManager) UpdateNodeBlockMG(name string, pod_maintenance_group string, description string, fabricNodeBlkattr models.NodeBlockAttributes) (*models.NodeBlock, error) {
	rn := fmt.Sprintf("nodeblk-%s", name)
	parentDn := fmt.Sprintf("uni/fabric/maintgrp-%s", pod_maintenance_group)
	fabricNodeBlk := models.NewNodeBlock(rn, parentDn, description, fabricNodeBlkattr)

	fabricNodeBlk.Status = "modified"
	err := sm.Save(fabricNodeBlk)
	return fabricNodeBlk, err

}

func (sm *ServiceManager) ListNodeBlockMG(pod_maintenance_group string) ([]*models.NodeBlock, error) {

	baseurlStr := "/api/node/class"
	dnUrl := fmt.Sprintf("%s/uni/fabric/maintgrp-%s/fabricNodeBlk.json", baseurlStr, pod_maintenance_group)

	cont, err := sm.GetViaURL(dnUrl)
	list := models.NodeBlockListFromContainer(cont)

	return list, err
}
