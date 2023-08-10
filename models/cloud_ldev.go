package models

import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/v2/container"
)

const (
	RnCloudLDev        = "cld-%s"
	DnCloudLDev        = "uni/tn-%s/cld-%s"
	ParentDnCloudLDev  = "uni/tn-%s"
	CloudLDevClassName = "cloudLDev"
)

type CloudL4L7Device struct {
	BaseAttributes
	CloudL4L7DeviceAttributes
}

type CloudL4L7DeviceAttributes struct {
	Version                            string `json:",omitempty"`
	ActiveActive                       string `json:",omitempty"`
	Annotation                         string `json:",omitempty"`
	ContextAware                       string `json:",omitempty"`
	CustomRG                           string `json:",omitempty"`
	DevType                            string `json:",omitempty"`
	FuncType                           string `json:",omitempty"`
	InstanceCount                      string `json:",omitempty"`
	IsCopy                             string `json:",omitempty"`
	IsInstantiation                    string `json:",omitempty"`
	L4L7DeviceApplicationSecurityGroup string `json:",omitempty"`
	L4L7ThirdPartyDevice               string `json:",omitempty"`
	Managed                            string `json:",omitempty"`
	Mode                               string `json:",omitempty"`
	Name                               string `json:",omitempty"`
	NameAlias                          string `json:",omitempty"`
	PackageModel                       string `json:",omitempty"`
	PromMode                           string `json:",omitempty"`
	SvcType                            string `json:",omitempty"`
	TargetMode                         string `json:",omitempty"`
	Trunking                           string `json:",omitempty"`
}

func NewCloudL4L7Device(cloudLDevRn, parentDn, description string, cloudLDevAttr CloudL4L7DeviceAttributes) *CloudL4L7Device {
	dn := fmt.Sprintf("%s/%s", parentDn, cloudLDevRn)
	return &CloudL4L7Device{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         CloudLDevClassName,
			Rn:                cloudLDevRn,
		},
		CloudL4L7DeviceAttributes: cloudLDevAttr,
	}
}

func (cloudLDev *CloudL4L7Device) ToMap() (map[string]string, error) {
	cloudLDevMap, err := cloudLDev.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	A(cloudLDevMap, "Version", cloudLDev.Version)
	A(cloudLDevMap, "activeActive", cloudLDev.ActiveActive)
	A(cloudLDevMap, "annotation", cloudLDev.Annotation)
	A(cloudLDevMap, "contextAware", cloudLDev.ContextAware)
	A(cloudLDevMap, "customRG", cloudLDev.CustomRG)
	A(cloudLDevMap, "devtype", cloudLDev.DevType)
	A(cloudLDevMap, "funcType", cloudLDev.FuncType)
	A(cloudLDevMap, "instanceCount", cloudLDev.InstanceCount)
	A(cloudLDevMap, "isCopy", cloudLDev.IsCopy)
	A(cloudLDevMap, "isInstantiation", cloudLDev.IsInstantiation)
	A(cloudLDevMap, "l4L7DeviceApplicationSecurityGroup", cloudLDev.L4L7DeviceApplicationSecurityGroup)
	A(cloudLDevMap, "l4L7ThirdPartyDevice", cloudLDev.L4L7ThirdPartyDevice)
	A(cloudLDevMap, "managed", cloudLDev.Managed)
	A(cloudLDevMap, "mode", cloudLDev.Mode)
	A(cloudLDevMap, "name", cloudLDev.Name)
	A(cloudLDevMap, "nameAlias", cloudLDev.NameAlias)
	A(cloudLDevMap, "packageModel", cloudLDev.PackageModel)
	A(cloudLDevMap, "promMode", cloudLDev.PromMode)
	A(cloudLDevMap, "svcType", cloudLDev.SvcType)
	A(cloudLDevMap, "targetMode", cloudLDev.TargetMode)
	A(cloudLDevMap, "trunking", cloudLDev.Trunking)
	return cloudLDevMap, err
}

func CloudL4L7DeviceFromContainerList(cont *container.Container, index int) *CloudL4L7Device {
	CloudL4L7DeviceCont := cont.S("imdata").Index(index).S(CloudLDevClassName, "attributes")
	return &CloudL4L7Device{
		BaseAttributes{
			DistinguishedName: G(CloudL4L7DeviceCont, "dn"),
			Description:       G(CloudL4L7DeviceCont, "descr"),
			Status:            G(CloudL4L7DeviceCont, "status"),
			ClassName:         CloudLDevClassName,
			Rn:                G(CloudL4L7DeviceCont, "rn"),
		},
		CloudL4L7DeviceAttributes{
			Version:                            G(CloudL4L7DeviceCont, "Version"),
			ActiveActive:                       G(CloudL4L7DeviceCont, "activeActive"),
			Annotation:                         G(CloudL4L7DeviceCont, "annotation"),
			ContextAware:                       G(CloudL4L7DeviceCont, "contextAware"),
			CustomRG:                           G(CloudL4L7DeviceCont, "customRG"),
			DevType:                            G(CloudL4L7DeviceCont, "devtype"),
			FuncType:                           G(CloudL4L7DeviceCont, "funcType"),
			InstanceCount:                      G(CloudL4L7DeviceCont, "instanceCount"),
			IsCopy:                             G(CloudL4L7DeviceCont, "isCopy"),
			IsInstantiation:                    G(CloudL4L7DeviceCont, "isInstantiation"),
			L4L7DeviceApplicationSecurityGroup: G(CloudL4L7DeviceCont, "l4L7DeviceApplicationSecurityGroup"),
			L4L7ThirdPartyDevice:               G(CloudL4L7DeviceCont, "l4L7ThirdPartyDevice"),
			Managed:                            G(CloudL4L7DeviceCont, "managed"),
			Mode:                               G(CloudL4L7DeviceCont, "mode"),
			Name:                               G(CloudL4L7DeviceCont, "name"),
			NameAlias:                          G(CloudL4L7DeviceCont, "nameAlias"),
			PackageModel:                       G(CloudL4L7DeviceCont, "packageModel"),
			PromMode:                           G(CloudL4L7DeviceCont, "promMode"),
			SvcType:                            G(CloudL4L7DeviceCont, "svcType"),
			TargetMode:                         G(CloudL4L7DeviceCont, "targetMode"),
			Trunking:                           G(CloudL4L7DeviceCont, "trunking"),
		},
	}
}

func CloudL4L7DeviceFromContainer(cont *container.Container) *CloudL4L7Device {
	return CloudL4L7DeviceFromContainerList(cont, 0)
}

func CloudL4L7DeviceListFromContainer(cont *container.Container) []*CloudL4L7Device {
	length, _ := strconv.Atoi(G(cont, "totalCount"))
	arr := make([]*CloudL4L7Device, length)

	for i := 0; i < length; i++ {
		arr[i] = CloudL4L7DeviceFromContainerList(cont, i)
	}

	return arr
}
