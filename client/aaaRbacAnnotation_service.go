package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/v2/models"
)

func (sm *ServiceManager) CreateAnnotationToCaptureRbacInfo(domain string, parentDn string, aaaRbacAnnotationAttr models.AnnotationToCaptureRbacInfoAttributes) (*models.AnnotationToCaptureRbacInfo, error) {
	rn := fmt.Sprintf(models.RnAaaRbacAnnotation, domain)
	aaaRbacAnnotation := models.NewAnnotationToCaptureRbacInfo(rn, parentDn, aaaRbacAnnotationAttr)
	err := sm.Save(aaaRbacAnnotation)
	return aaaRbacAnnotation, err
}

func (sm *ServiceManager) ReadAnnotationToCaptureRbacInfo(domain string, parentDn string) (*models.AnnotationToCaptureRbacInfo, error) {
	rn := fmt.Sprintf(models.RnAaaRbacAnnotation, domain)
	dn := fmt.Sprintf("%s/%s", parentDn, rn)
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}
	aaaRbacAnnotation := models.AnnotationToCaptureRbacInfoFromContainer(cont)
	return aaaRbacAnnotation, nil
}

func (sm *ServiceManager) DeleteAnnotationToCaptureRbacInfo(domain string, parentDn string) error {
	rn := fmt.Sprintf(models.RnAaaRbacAnnotation, domain)
	dn := fmt.Sprintf("%s/%s", parentDn, rn)
	return sm.DeleteByDn(dn, models.AaaRbacAnnotationClassName)
}

func (sm *ServiceManager) UpdateAnnotationToCaptureRbacInfo(domain string, parentDn string, aaaRbacAnnotationAttr models.AnnotationToCaptureRbacInfoAttributes) (*models.AnnotationToCaptureRbacInfo, error) {
	rn := fmt.Sprintf(models.RnAaaRbacAnnotation, domain)
	aaaRbacAnnotation := models.NewAnnotationToCaptureRbacInfo(rn, parentDn, aaaRbacAnnotationAttr)
	aaaRbacAnnotation.Status = "modified"
	err := sm.Save(aaaRbacAnnotation)
	return aaaRbacAnnotation, err
}

func (sm *ServiceManager) ListAnnotationToCaptureRbacInfo(parentDn string) ([]*models.AnnotationToCaptureRbacInfo, error) {
	dnUrl := fmt.Sprintf("%s/%s/%s.json", models.BaseurlStr, parentDn, models.AaaRbacAnnotationClassName)
	cont, err := sm.GetViaURL(dnUrl)
	list := models.AnnotationToCaptureRbacInfoListFromContainer(cont)
	return list, err
}
