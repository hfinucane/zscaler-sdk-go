package scimattributeheader

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/zscaler/zscaler-sdk-go/zpa/services/common"
)

const (
	mgmtConfig       = "/mgmtconfig/v1/admin/customers/"
	userConfig       = "/userconfig/v1/customers"
	idpId            = "/idp"
	scimAttrEndpoint = "/scimattribute"
)

type ScimAttributeHeader struct {
	CanonicalValues []string `json:"canonicalValues,omitempty"`
	CaseSensitive   bool     `json:"caseSensitive,omitempty"`
	CreationTime    string   `json:"creationTime,omitempty,"`
	DataType        string   `json:"dataType,omitempty"`
	Description     string   `json:"description,omitempty"`
	ID              string   `json:"id,omitempty"`
	IdpID           string   `json:"idpId,omitempty"`
	ModifiedBy      string   `json:"modifiedBy,omitempty"`
	ModifiedTime    string   `json:"modifiedTime,omitempty"`
	MultiValued     bool     `json:"multivalued,omitempty"`
	Mutability      string   `json:"mutability,omitempty"`
	Name            string   `json:"name,omitempty"`
	Required        bool     `json:"required,omitempty"`
	Returned        string   `json:"returned,omitempty"`
	SchemaURI       string   `json:"schemaURI,omitempty"`
	Uniqueness      bool     `json:"uniqueness,omitempty"`
}

func (service *Service) Get(idpId, scimAttrHeaderID string) (*ScimAttributeHeader, *http.Response, error) {
	v := new(ScimAttributeHeader)
	relativeURL := fmt.Sprintf("%s/idp/%s/scimattribute/%s", mgmtConfig+service.Client.Config.CustomerID, idpId, scimAttrHeaderID)
	resp, err := service.Client.NewRequestDo("GET", relativeURL, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetValues(idpId, ScimAttrHeaderID string) ([]string, error) {
	var v struct {
		List []string `json:"list"`
	}
	relativeURL := fmt.Sprintf("%s/%s/scimattribute/idpId/%s/attributeId/%s", userConfig, service.Client.Config.CustomerID, idpId, ScimAttrHeaderID)
	_, err := service.Client.NewRequestDo("GET", relativeURL, common.Pagination{PageSize: common.DefaultPageSize}, nil, &v)
	if err != nil {
		return nil, err
	}

	return v.List, nil
}

func (service *Service) GetByName(scimAttributeName, IdpId string) (*ScimAttributeHeader, *http.Response, error) {
	var v struct {
		List []ScimAttributeHeader `json:"list"`
	}
	relativeURL := fmt.Sprintf("%s/%s%s", mgmtConfig+service.Client.Config.CustomerID+idpId, IdpId, scimAttrEndpoint)
	resp, err := service.Client.NewRequestDo("GET", relativeURL, common.Pagination{PageSize: common.DefaultPageSize, Search: scimAttributeName}, nil, &v)
	if err != nil {
		return nil, nil, err
	}
	for _, scimAttribute := range v.List {
		if strings.EqualFold(scimAttribute.Name, scimAttributeName) {
			return &scimAttribute, resp, nil
		}
	}
	return nil, resp, fmt.Errorf("no scim named '%s' was found", scimAttributeName)
}

func (service *Service) GetAllByIdpId(IdpId string) ([]ScimAttributeHeader, *http.Response, error) {
	var v struct {
		List []ScimAttributeHeader `json:"list"`
	}
	relativeURL := fmt.Sprintf("%s/%s%s", mgmtConfig+service.Client.Config.CustomerID+idpId, IdpId, scimAttrEndpoint)
	resp, err := service.Client.NewRequestDo("GET", relativeURL, common.Pagination{PageSize: common.DefaultPageSize}, nil, &v)
	if err != nil {
		return nil, nil, err
	}
	return v.List, resp, nil
}
