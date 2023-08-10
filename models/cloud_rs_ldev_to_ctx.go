package models

import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/v2/container"
)

const (
	RnCloudRsLDevToCtx        = "rsLDevToCtx"
	DnCloudRsLDevToCtx        = "uni/tn-%s/cld-%s/rsLDevToCtx"
	ParentDnCloudRsLDevToCtx  = "uni/tn-%s/cld-%s"
	CloudRsLDevToCtxClassName = "cloudRsLDevToCtx"
)

type RelationFromCloudLDevToCloudCtx struct {
	BaseAttributes
	RelationFromCloudLDevToCloudCtxAttributes
}

type RelationFromCloudLDevToCloudCtxAttributes struct {
	Annotation string `json:",omitempty"`
	TDn        string `json:",omitempty"`
}

func NewRelationFromCloudLDevToCloudCtx(cloudRsLDevToCtxRn, parentDn, description string, cloudRsLDevToCtxAttr RelationFromCloudLDevToCloudCtxAttributes) *RelationFromCloudLDevToCloudCtx {
	dn := fmt.Sprintf("%s/%s", parentDn, cloudRsLDevToCtxRn)
	return &RelationFromCloudLDevToCloudCtx{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         CloudRsLDevToCtxClassName,
			Rn:                cloudRsLDevToCtxRn,
		},
		RelationFromCloudLDevToCloudCtxAttributes: cloudRsLDevToCtxAttr,
	}
}

func (cloudRsLDevToCtx *RelationFromCloudLDevToCloudCtx) ToMap() (map[string]string, error) {
	cloudRsLDevToCtxMap, err := cloudRsLDevToCtx.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	A(cloudRsLDevToCtxMap, "annotation", cloudRsLDevToCtx.Annotation)
	A(cloudRsLDevToCtxMap, "tDn", cloudRsLDevToCtx.TDn)
	return cloudRsLDevToCtxMap, err
}

func RelationFromCloudLDevToCloudCtxFromContainerList(cont *container.Container, index int) *RelationFromCloudLDevToCloudCtx {
	RelationFromCloudLDevToCloudCtxCont := cont.S("imdata").Index(index).S(CloudRsLDevToCtxClassName, "attributes")
	return &RelationFromCloudLDevToCloudCtx{
		BaseAttributes{
			DistinguishedName: G(RelationFromCloudLDevToCloudCtxCont, "dn"),
			Description:       G(RelationFromCloudLDevToCloudCtxCont, "descr"),
			Status:            G(RelationFromCloudLDevToCloudCtxCont, "status"),
			ClassName:         CloudRsLDevToCtxClassName,
			Rn:                G(RelationFromCloudLDevToCloudCtxCont, "rn"),
		},
		RelationFromCloudLDevToCloudCtxAttributes{
			Annotation: G(RelationFromCloudLDevToCloudCtxCont, "annotation"),
			TDn:        G(RelationFromCloudLDevToCloudCtxCont, "tDn"),
		},
	}
}

func RelationFromCloudLDevToCloudCtxFromContainer(cont *container.Container) *RelationFromCloudLDevToCloudCtx {
	return RelationFromCloudLDevToCloudCtxFromContainerList(cont, 0)
}

func RelationFromCloudLDevToCloudCtxListFromContainer(cont *container.Container) []*RelationFromCloudLDevToCloudCtx {
	length, _ := strconv.Atoi(G(cont, "totalCount"))
	arr := make([]*RelationFromCloudLDevToCloudCtx, length)

	for i := 0; i < length; i++ {
		arr[i] = RelationFromCloudLDevToCloudCtxFromContainerList(cont, i)
	}

	return arr
}
