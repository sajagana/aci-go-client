package models

import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/v2/container"
)

const (
	RnAaaRbacAnnotation        = "rbacDom-%s"
	AaaRbacAnnotationClassName = "aaaRbacAnnotation"
)

type AnnotationToCaptureRbacInfo struct {
	BaseAttributes
	AnnotationToCaptureRbacInfoAttributes
}

type AnnotationToCaptureRbacInfoAttributes struct {
	ChildRegex string `json:",omitempty"`
	Domain     string `json:",omitempty"`
}

func NewAnnotationToCaptureRbacInfo(aaaRbacAnnotationRn, parentDn string, aaaRbacAnnotationAttr AnnotationToCaptureRbacInfoAttributes) *AnnotationToCaptureRbacInfo {
	dn := fmt.Sprintf("%s/%s", parentDn, aaaRbacAnnotationRn)
	return &AnnotationToCaptureRbacInfo{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Status:            "created, modified",
			ClassName:         AaaRbacAnnotationClassName,
			Rn:                aaaRbacAnnotationRn,
		},
		AnnotationToCaptureRbacInfoAttributes: aaaRbacAnnotationAttr,
	}
}

func (aaaRbacAnnotation *AnnotationToCaptureRbacInfo) ToMap() (map[string]string, error) {
	aaaRbacAnnotationMap, err := aaaRbacAnnotation.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

	A(aaaRbacAnnotationMap, "childRegex", aaaRbacAnnotation.ChildRegex)
	A(aaaRbacAnnotationMap, "domain", aaaRbacAnnotation.Domain)
	return aaaRbacAnnotationMap, err
}

func AnnotationToCaptureRbacInfoFromContainerList(cont *container.Container, index int) *AnnotationToCaptureRbacInfo {
	AnnotationToCaptureRbacInfoCont := cont.S("imdata").Index(index).S(AaaRbacAnnotationClassName, "attributes")
	return &AnnotationToCaptureRbacInfo{
		BaseAttributes{
			DistinguishedName: G(AnnotationToCaptureRbacInfoCont, "dn"),
			Status:            G(AnnotationToCaptureRbacInfoCont, "status"),
			ClassName:         AaaRbacAnnotationClassName,
			Rn:                G(AnnotationToCaptureRbacInfoCont, "rn"),
		},
		AnnotationToCaptureRbacInfoAttributes{
			ChildRegex: G(AnnotationToCaptureRbacInfoCont, "childRegex"),
			Domain:     G(AnnotationToCaptureRbacInfoCont, "domain"),
		},
	}
}

func AnnotationToCaptureRbacInfoFromContainer(cont *container.Container) *AnnotationToCaptureRbacInfo {
	return AnnotationToCaptureRbacInfoFromContainerList(cont, 0)
}

func AnnotationToCaptureRbacInfoListFromContainer(cont *container.Container) []*AnnotationToCaptureRbacInfo {
	length, _ := strconv.Atoi(G(cont, "totalCount"))
	arr := make([]*AnnotationToCaptureRbacInfo, length)

	for i := 0; i < length; i++ {
		arr[i] = AnnotationToCaptureRbacInfoFromContainerList(cont, i)
	}

	return arr
}
