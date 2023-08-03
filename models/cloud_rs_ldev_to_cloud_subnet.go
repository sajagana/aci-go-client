package models

import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/v2/container"
)

const (
	RnCloudRsLDevToCloudSubnet        = "rsLDevToCloudSubnet-[%s]"
	CloudRsLDevToCloudSubnetClassName = "cloudRsLDevToCloudSubnet"
)

type RelationFromCloudLDevToCloudSubnet struct {
	BaseAttributes
	RelationFromCloudLDevToCloudSubnetAttributes
}

type RelationFromCloudLDevToCloudSubnetAttributes struct {
	Annotation string `json:",omitempty"`
	TDn        string `json:",omitempty"`
}

func NewRelationFromCloudLDevToCloudSubnet(cloudRsLDevToCloudSubnetRn, parentDn string, cloudRsLDevToCloudSubnetAttr RelationFromCloudLDevToCloudSubnetAttributes) *RelationFromCloudLDevToCloudSubnet {
	dn := fmt.Sprintf("%s/%s", parentDn, cloudRsLDevToCloudSubnetRn)
	return &RelationFromCloudLDevToCloudSubnet{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Status:            "created, modified",
			ClassName:         CloudRsLDevToCloudSubnetClassName,
			Rn:                cloudRsLDevToCloudSubnetRn,
		},
		RelationFromCloudLDevToCloudSubnetAttributes: cloudRsLDevToCloudSubnetAttr,
	}
}

func (cloudRsLDevToCloudSubnet *RelationFromCloudLDevToCloudSubnet) ToMap() (map[string]string, error) {
	cloudRsLDevToCloudSubnetMap, err := cloudRsLDevToCloudSubnet.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	A(cloudRsLDevToCloudSubnetMap, "annotation", cloudRsLDevToCloudSubnet.Annotation)
	A(cloudRsLDevToCloudSubnetMap, "tDn", cloudRsLDevToCloudSubnet.TDn)
	return cloudRsLDevToCloudSubnetMap, err
}

func RelationFromCloudLDevToCloudSubnetFromContainerList(cont *container.Container, index int) *RelationFromCloudLDevToCloudSubnet {
	RelationFromCloudLDevToCloudSubnetCont := cont.S("imdata").Index(index).S(CloudRsLDevToCloudSubnetClassName, "attributes")
	return &RelationFromCloudLDevToCloudSubnet{
		BaseAttributes{
			DistinguishedName: G(RelationFromCloudLDevToCloudSubnetCont, "dn"),
			Status:            G(RelationFromCloudLDevToCloudSubnetCont, "status"),
			ClassName:         CloudRsLDevToCloudSubnetClassName,
			Rn:                G(RelationFromCloudLDevToCloudSubnetCont, "rn"),
		},
		RelationFromCloudLDevToCloudSubnetAttributes{
			Annotation: G(RelationFromCloudLDevToCloudSubnetCont, "annotation"),
			TDn:        G(RelationFromCloudLDevToCloudSubnetCont, "tDn"),
		},
	}
}

func RelationFromCloudLDevToCloudSubnetFromContainer(cont *container.Container) *RelationFromCloudLDevToCloudSubnet {
	return RelationFromCloudLDevToCloudSubnetFromContainerList(cont, 0)
}

func RelationFromCloudLDevToCloudSubnetListFromContainer(cont *container.Container) []*RelationFromCloudLDevToCloudSubnet {
	length, _ := strconv.Atoi(G(cont, "totalCount"))
	arr := make([]*RelationFromCloudLDevToCloudSubnet, length)

	for i := 0; i < length; i++ {
		arr[i] = RelationFromCloudLDevToCloudSubnetFromContainerList(cont, i)
	}

	return arr
}
