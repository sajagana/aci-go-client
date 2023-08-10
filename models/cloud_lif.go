package models

import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/v2/container"
)

const (
	RnCloudLIf        = "clif-%s"
	DnCloudLIf        = "uni/tn-%s/cld-%s/clif-%s"
	ParentDnCloudLIf  = "uni/tn-%s/cld-%s"
	CloudLIfClassName = "cloudLIf"
)

type CloudL4L7LogicalInterface struct {
	BaseAttributes
	CloudL4L7LogicalInterfaceAttributes
}

type CloudL4L7LogicalInterfaceAttributes struct {
	AllowAll   string `json:",omitempty"`
	Annotation string `json:",omitempty"`
	Name       string `json:",omitempty"`
	NameAlias  string `json:",omitempty"`
}

func NewCloudL4L7LogicalInterface(cloudLIfRn, parentDn string, cloudLIfAttr CloudL4L7LogicalInterfaceAttributes) *CloudL4L7LogicalInterface {
	dn := fmt.Sprintf("%s/%s", parentDn, cloudLIfRn)
	return &CloudL4L7LogicalInterface{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Status:            "created, modified",
			ClassName:         CloudLIfClassName,
			Rn:                cloudLIfRn,
		},
		CloudL4L7LogicalInterfaceAttributes: cloudLIfAttr,
	}
}

func (cloudLIf *CloudL4L7LogicalInterface) ToMap() (map[string]string, error) {
	cloudLIfMap, err := cloudLIf.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	A(cloudLIfMap, "allowAll", cloudLIf.AllowAll)
	A(cloudLIfMap, "annotation", cloudLIf.Annotation)
	A(cloudLIfMap, "name", cloudLIf.Name)
	A(cloudLIfMap, "nameAlias", cloudLIf.NameAlias)
	return cloudLIfMap, err
}

func CloudL4L7LogicalInterfaceFromContainerList(cont *container.Container, index int) *CloudL4L7LogicalInterface {
	CloudL4L7LogicalInterfaceCont := cont.S("imdata").Index(index).S(CloudLIfClassName, "attributes")
	return &CloudL4L7LogicalInterface{
		BaseAttributes{
			DistinguishedName: G(CloudL4L7LogicalInterfaceCont, "dn"),
			Status:            G(CloudL4L7LogicalInterfaceCont, "status"),
			ClassName:         CloudLIfClassName,
			Rn:                G(CloudL4L7LogicalInterfaceCont, "rn"),
		},
		CloudL4L7LogicalInterfaceAttributes{
			AllowAll:   G(CloudL4L7LogicalInterfaceCont, "allowAll"),
			Annotation: G(CloudL4L7LogicalInterfaceCont, "annotation"),
			Name:       G(CloudL4L7LogicalInterfaceCont, "name"),
			NameAlias:  G(CloudL4L7LogicalInterfaceCont, "nameAlias"),
		},
	}
}

func CloudL4L7LogicalInterfaceFromContainer(cont *container.Container) *CloudL4L7LogicalInterface {
	return CloudL4L7LogicalInterfaceFromContainerList(cont, 0)
}

func CloudL4L7LogicalInterfaceListFromContainer(cont *container.Container) []*CloudL4L7LogicalInterface {
	length, _ := strconv.Atoi(G(cont, "totalCount"))
	arr := make([]*CloudL4L7LogicalInterface, length)

	for i := 0; i < length; i++ {
		arr[i] = CloudL4L7LogicalInterfaceFromContainerList(cont, i)
	}

	return arr
}
