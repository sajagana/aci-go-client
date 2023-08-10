package models

import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/v2/container"
)

const (
	RnVnsRsNodeToCloudLDev        = "rsNodeToCloudLDev"
	DnVnsRsNodeToCloudLDev        = "uni/tn-%s/AbsGraph-%s/AbsNode-%s/rsNodeToCloudLDev"
	ParentDnVnsRsNodeToCloudLDev  = "uni/tn-%s/AbsGraph-%s/AbsNode-%s"
	VnsRsNodeToCloudLDevClassName = "vnsRsNodeToCloudLDev"
)

type RelationFromAbsNodeToCloudLDev struct {
	BaseAttributes
	RelationFromAbsNodeToCloudLDevAttributes
}

type RelationFromAbsNodeToCloudLDevAttributes struct {
	Annotation string `json:",omitempty"`
	TDn        string `json:",omitempty"`
}

func NewRelationFromAbsNodeToCloudLDev(vnsRsNodeToCloudLDevRn, parentDn, description string, vnsRsNodeToCloudLDevAttr RelationFromAbsNodeToCloudLDevAttributes) *RelationFromAbsNodeToCloudLDev {
	dn := fmt.Sprintf("%s/%s", parentDn, vnsRsNodeToCloudLDevRn)
	return &RelationFromAbsNodeToCloudLDev{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       "",
			Status:            "created, modified",
			ClassName:         VnsRsNodeToCloudLDevClassName,
			Rn:                vnsRsNodeToCloudLDevRn,
		},
		RelationFromAbsNodeToCloudLDevAttributes: vnsRsNodeToCloudLDevAttr,
	}
}

func (vnsRsNodeToCloudLDev *RelationFromAbsNodeToCloudLDev) ToMap() (map[string]string, error) {
	vnsRsNodeToCloudLDevMap, err := vnsRsNodeToCloudLDev.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	A(vnsRsNodeToCloudLDevMap, "annotation", vnsRsNodeToCloudLDev.Annotation)
	A(vnsRsNodeToCloudLDevMap, "tDn", vnsRsNodeToCloudLDev.TDn)
	return vnsRsNodeToCloudLDevMap, err
}

func RelationFromAbsNodeToCloudLDevFromContainerList(cont *container.Container, index int) *RelationFromAbsNodeToCloudLDev {
	RelationFromAbsNodeToCloudLDevCont := cont.S("imdata").Index(index).S(VnsRsNodeToCloudLDevClassName, "attributes")
	return &RelationFromAbsNodeToCloudLDev{
		BaseAttributes{
			DistinguishedName: G(RelationFromAbsNodeToCloudLDevCont, "dn"),
			Description:       G(RelationFromAbsNodeToCloudLDevCont, "descr"),
			Status:            G(RelationFromAbsNodeToCloudLDevCont, "status"),
			ClassName:         VnsRsNodeToCloudLDevClassName,
			Rn:                G(RelationFromAbsNodeToCloudLDevCont, "rn"),
		},
		RelationFromAbsNodeToCloudLDevAttributes{
			Annotation: G(RelationFromAbsNodeToCloudLDevCont, "annotation"),
			TDn:        G(RelationFromAbsNodeToCloudLDevCont, "tDn"),
		},
	}
}

func RelationFromAbsNodeToCloudLDevFromContainer(cont *container.Container) *RelationFromAbsNodeToCloudLDev {
	return RelationFromAbsNodeToCloudLDevFromContainerList(cont, 0)
}

func RelationFromAbsNodeToCloudLDevListFromContainer(cont *container.Container) []*RelationFromAbsNodeToCloudLDev {
	length, _ := strconv.Atoi(G(cont, "totalCount"))
	arr := make([]*RelationFromAbsNodeToCloudLDev, length)

	for i := 0; i < length; i++ {
		arr[i] = RelationFromAbsNodeToCloudLDevFromContainerList(cont, i)
	}

	return arr
}
