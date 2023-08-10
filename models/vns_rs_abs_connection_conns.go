package models

import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/v2/container"
)

const (
	RnVnsRsAbsConnectionConns        = "rsabsConnectionConns-[%s]"
	DnVnsRsAbsConnectionConns        = "uni/tn-%s/AbsGraph-%s/AbsConnection-%s/rsabsConnectionConns-[%s]"
	ParentDnVnsRsAbsConnectionConns  = "uni/tn-%s/AbsGraph-%s/AbsConnection-%s"
	VnsRsAbsConnectionConnsClassName = "vnsRsAbsConnectionConns"
)

type RelationFromServiceGraphConnectionToServiceGraphConnectors struct {
	BaseAttributes
	RelationFromServiceGraphConnectionToServiceGraphConnectorsAttributes
}

type RelationFromServiceGraphConnectionToServiceGraphConnectorsAttributes struct {
	Annotation string `json:",omitempty"`
	TDn        string `json:",omitempty"`
}

func NewRelationFromServiceGraphConnectionToServiceGraphConnectors(vnsRsAbsConnectionConnsRn, parentDn, description string, vnsRsAbsConnectionConnsAttr RelationFromServiceGraphConnectionToServiceGraphConnectorsAttributes) *RelationFromServiceGraphConnectionToServiceGraphConnectors {
	dn := fmt.Sprintf("%s/%s", parentDn, vnsRsAbsConnectionConnsRn)
	return &RelationFromServiceGraphConnectionToServiceGraphConnectors{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         VnsRsAbsConnectionConnsClassName,
			Rn:                vnsRsAbsConnectionConnsRn,
		},
		RelationFromServiceGraphConnectionToServiceGraphConnectorsAttributes: vnsRsAbsConnectionConnsAttr,
	}
}

func (vnsRsAbsConnectionConns *RelationFromServiceGraphConnectionToServiceGraphConnectors) ToMap() (map[string]string, error) {
	vnsRsAbsConnectionConnsMap, err := vnsRsAbsConnectionConns.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	A(vnsRsAbsConnectionConnsMap, "annotation", vnsRsAbsConnectionConns.Annotation)
	A(vnsRsAbsConnectionConnsMap, "tDn", vnsRsAbsConnectionConns.TDn)
	return vnsRsAbsConnectionConnsMap, err
}

func RelationFromServiceGraphConnectionToServiceGraphConnectorsFromContainerList(cont *container.Container, index int) *RelationFromServiceGraphConnectionToServiceGraphConnectors {
	RelationFromServiceGraphConnectionToServiceGraphConnectorsCont := cont.S("imdata").Index(index).S(VnsRsAbsConnectionConnsClassName, "attributes")
	return &RelationFromServiceGraphConnectionToServiceGraphConnectors{
		BaseAttributes{
			DistinguishedName: G(RelationFromServiceGraphConnectionToServiceGraphConnectorsCont, "dn"),
			Description:       G(RelationFromServiceGraphConnectionToServiceGraphConnectorsCont, "descr"),
			Status:            G(RelationFromServiceGraphConnectionToServiceGraphConnectorsCont, "status"),
			ClassName:         VnsRsAbsConnectionConnsClassName,
			Rn:                G(RelationFromServiceGraphConnectionToServiceGraphConnectorsCont, "rn"),
		},
		RelationFromServiceGraphConnectionToServiceGraphConnectorsAttributes{
			Annotation: G(RelationFromServiceGraphConnectionToServiceGraphConnectorsCont, "annotation"),
			TDn:        G(RelationFromServiceGraphConnectionToServiceGraphConnectorsCont, "tDn"),
		},
	}
}

func RelationFromServiceGraphConnectionToServiceGraphConnectorsFromContainer(cont *container.Container) *RelationFromServiceGraphConnectionToServiceGraphConnectors {
	return RelationFromServiceGraphConnectionToServiceGraphConnectorsFromContainerList(cont, 0)
}

func RelationFromServiceGraphConnectionToServiceGraphConnectorsListFromContainer(cont *container.Container) []*RelationFromServiceGraphConnectionToServiceGraphConnectors {
	length, _ := strconv.Atoi(G(cont, "totalCount"))
	arr := make([]*RelationFromServiceGraphConnectionToServiceGraphConnectors, length)

	for i := 0; i < length; i++ {
		arr[i] = RelationFromServiceGraphConnectionToServiceGraphConnectorsFromContainerList(cont, i)
	}

	return arr
}
