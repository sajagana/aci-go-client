package models

import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/v2/container"
)

const (
	RnCloudLB        = "clb-%s"
	DnCloudLB        = "uni/tn-%s/clb-%s"
	ParentDnCloudLB  = "uni/tn-%s"
	CloudLBClassName = "cloudLB"
)

type CloudL4L7LB struct {
	BaseAttributes
	CloudL4L7LBAttributes
}

type CloudL4L7LBAttributes struct {
	Version      string `json:",omitempty"`
	ActiveActive string `json:",omitempty"`
	AllowAll     string `json:",omitempty"`
	Annotation   string `json:",omitempty"`
	AutoScaling  string `json:",omitempty"`
	ContextAware string `json:",omitempty"`
	CustomRG     string `json:",omitempty"`
	DevType      string `json:",omitempty"`
	// FirewallMode                       string `json:",omitempty"`
	// FirewallStatus                     string `json:",omitempty"`
	FuncType                           string `json:",omitempty"`
	InstanceCount                      string `json:",omitempty"`
	IsCopy                             string `json:",omitempty"`
	IsInstantiation                    string `json:",omitempty"`
	IsStaticIP                         string `json:",omitempty"`
	L4L7DeviceApplicationSecurityGroup string `json:",omitempty"`
	L4L7ThirdPartyDevice               string `json:",omitempty"`
	Managed                            string `json:",omitempty"`
	MaxInstanceCount                   string `json:",omitempty"`
	MinInstanceCount                   string `json:",omitempty"`
	Mode                               string `json:",omitempty"`
	Name                               string `json:",omitempty"`
	NameAlias                          string `json:",omitempty"`
	NativeLBName                       string `json:",omitempty"`
	// Oid                                string `json:",omitempty"`
	PackageModel     string `json:",omitempty"`
	PromMode         string `json:",omitempty"`
	Scheme           string `json:",omitempty"`
	Size             string `json:",omitempty"`
	Sku              string `json:",omitempty"`
	SvcType          string `json:",omitempty"`
	TargetMode       string `json:",omitempty"`
	Trunking         string `json:",omitempty"`
	CloudL4L7LB_type string `json:",omitempty"`
}

func NewCloudL4L7LB(cloudLBRn, parentDn, description string, cloudLBAttr CloudL4L7LBAttributes) *CloudL4L7LB {
	dn := fmt.Sprintf("%s/%s", parentDn, cloudLBRn)
	return &CloudL4L7LB{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         CloudLBClassName,
			Rn:                cloudLBRn,
		},
		CloudL4L7LBAttributes: cloudLBAttr,
	}
}

func (cloudLB *CloudL4L7LB) ToMap() (map[string]string, error) {
	cloudLBMap, err := cloudLB.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	A(cloudLBMap, "Version", cloudLB.Version)
	A(cloudLBMap, "activeActive", cloudLB.ActiveActive)
	A(cloudLBMap, "allowAll", cloudLB.AllowAll)
	A(cloudLBMap, "annotation", cloudLB.Annotation)
	A(cloudLBMap, "autoScaling", cloudLB.AutoScaling)
	A(cloudLBMap, "contextAware", cloudLB.ContextAware)
	A(cloudLBMap, "customRG", cloudLB.CustomRG)
	A(cloudLBMap, "devtype", cloudLB.DevType)
	// A(cloudLBMap, "firewallMode", cloudLB.FirewallMode)
	// A(cloudLBMap, "firewallStatus", cloudLB.FirewallStatus)
	A(cloudLBMap, "funcType", cloudLB.FuncType)
	A(cloudLBMap, "instanceCount", cloudLB.InstanceCount)
	A(cloudLBMap, "isCopy", cloudLB.IsCopy)
	A(cloudLBMap, "isInstantiation", cloudLB.IsInstantiation)
	A(cloudLBMap, "isStaticIP", cloudLB.IsStaticIP)
	A(cloudLBMap, "l4L7DeviceApplicationSecurityGroup", cloudLB.L4L7DeviceApplicationSecurityGroup)
	A(cloudLBMap, "l4L7ThirdPartyDevice", cloudLB.L4L7ThirdPartyDevice)
	A(cloudLBMap, "managed", cloudLB.Managed)
	A(cloudLBMap, "maxInstanceCount", cloudLB.MaxInstanceCount)
	A(cloudLBMap, "minInstanceCount", cloudLB.MinInstanceCount)
	A(cloudLBMap, "mode", cloudLB.Mode)
	A(cloudLBMap, "name", cloudLB.Name)
	A(cloudLBMap, "nameAlias", cloudLB.NameAlias)
	A(cloudLBMap, "nativeLBName", cloudLB.NativeLBName)
	// A(cloudLBMap, "oid", cloudLB.Oid)
	A(cloudLBMap, "packageModel", cloudLB.PackageModel)
	A(cloudLBMap, "promMode", cloudLB.PromMode)
	A(cloudLBMap, "scheme", cloudLB.Scheme)
	A(cloudLBMap, "size", cloudLB.Size)
	A(cloudLBMap, "sku", cloudLB.Sku)
	A(cloudLBMap, "svcType", cloudLB.SvcType)
	A(cloudLBMap, "targetMode", cloudLB.TargetMode)
	A(cloudLBMap, "trunking", cloudLB.Trunking)
	A(cloudLBMap, "type", cloudLB.CloudL4L7LB_type)
	return cloudLBMap, err
}

func CloudL4L7LBFromContainerList(cont *container.Container, index int) *CloudL4L7LB {
	CloudL4L7LBCont := cont.S("imdata").Index(index).S(CloudLBClassName, "attributes")
	return &CloudL4L7LB{
		BaseAttributes{
			DistinguishedName: G(CloudL4L7LBCont, "dn"),
			Description:       G(CloudL4L7LBCont, "descr"),
			Status:            G(CloudL4L7LBCont, "status"),
			ClassName:         CloudLBClassName,
			Rn:                G(CloudL4L7LBCont, "rn"),
		},
		CloudL4L7LBAttributes{
			Version:      G(CloudL4L7LBCont, "Version"),
			ActiveActive: G(CloudL4L7LBCont, "activeActive"),
			AllowAll:     G(CloudL4L7LBCont, "allowAll"),
			Annotation:   G(CloudL4L7LBCont, "annotation"),
			AutoScaling:  G(CloudL4L7LBCont, "autoScaling"),
			ContextAware: G(CloudL4L7LBCont, "contextAware"),
			CustomRG:     G(CloudL4L7LBCont, "customRG"),
			DevType:      G(CloudL4L7LBCont, "devtype"),
			// FirewallMode:                       G(CloudL4L7LBCont, "firewallMode"),
			// FirewallStatus:                     G(CloudL4L7LBCont, "firewallStatus"),
			FuncType:                           G(CloudL4L7LBCont, "funcType"),
			InstanceCount:                      G(CloudL4L7LBCont, "instanceCount"),
			IsCopy:                             G(CloudL4L7LBCont, "isCopy"),
			IsInstantiation:                    G(CloudL4L7LBCont, "isInstantiation"),
			IsStaticIP:                         G(CloudL4L7LBCont, "isStaticIP"),
			L4L7DeviceApplicationSecurityGroup: G(CloudL4L7LBCont, "l4L7DeviceApplicationSecurityGroup"),
			L4L7ThirdPartyDevice:               G(CloudL4L7LBCont, "l4L7ThirdPartyDevice"),
			Managed:                            G(CloudL4L7LBCont, "managed"),
			MaxInstanceCount:                   G(CloudL4L7LBCont, "maxInstanceCount"),
			MinInstanceCount:                   G(CloudL4L7LBCont, "minInstanceCount"),
			Mode:                               G(CloudL4L7LBCont, "mode"),
			Name:                               G(CloudL4L7LBCont, "name"),
			NameAlias:                          G(CloudL4L7LBCont, "nameAlias"),
			NativeLBName:                       G(CloudL4L7LBCont, "nativeLBName"),
			// Oid:                                G(CloudL4L7LBCont, "oid"),
			PackageModel:     G(CloudL4L7LBCont, "packageModel"),
			PromMode:         G(CloudL4L7LBCont, "promMode"),
			Scheme:           G(CloudL4L7LBCont, "scheme"),
			Size:             G(CloudL4L7LBCont, "size"),
			Sku:              G(CloudL4L7LBCont, "sku"),
			SvcType:          G(CloudL4L7LBCont, "svcType"),
			TargetMode:       G(CloudL4L7LBCont, "targetMode"),
			Trunking:         G(CloudL4L7LBCont, "trunking"),
			CloudL4L7LB_type: G(CloudL4L7LBCont, "type"),
		},
	}
}

func CloudL4L7LBFromContainer(cont *container.Container) *CloudL4L7LB {
	return CloudL4L7LBFromContainerList(cont, 0)
}

func CloudL4L7LBListFromContainer(cont *container.Container) []*CloudL4L7LB {
	length, _ := strconv.Atoi(G(cont, "totalCount"))
	arr := make([]*CloudL4L7LB, length)

	for i := 0; i < length; i++ {
		arr[i] = CloudL4L7LBFromContainerList(cont, i)
	}

	return arr
}
