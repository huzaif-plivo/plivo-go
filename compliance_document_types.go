package plivo

import "time"

type ComplianceDocumentTypeService struct {
	client *Client
}

type GetComplianceDocumentTypeResponse struct {
	APIID string `json:"api_id"`
	Meta  struct {
		Limit      int         `json:"limit"`
		Next       interface{} `json:"next"`
		Offset     int         `json:"offset"`
		Previous   interface{} `json:"previous"`
		TotalCount int         `json:"total_count"`
	} `json:"meta"`
	Objects []struct {
		CreatedAt      time.Time `json:"created_at"`
		DocumentTypeID string    `json:"document_type_id"`
		DocumentName   string    `json:"document_name"`
		Description    string    `json:"description"`
		ProofRequired  string    `json:"proof_required"`
		Information    []struct {
			FieldName    string `json:"field_name"`
			FieldType    string `json:"field_type"`
			FriendlyName string `json:"friendly_name"`
			HelpText     string `json:"help_text,omitempty"`
			MaxLength    int    `json:"max_length,omitempty"`
			MinLength    int    `json:"min_length,omitempty"`
			Format       string `json:"format,omitempty"`
		} `json:"information"`
	} `json:"objects"`
}

type ListComplianceDocumentTypeResponse struct {
	APIID string `json:"api_id"`
	Meta  struct {
		Limit      int         `json:"limit"`
		Next       interface{} `json:"next"`
		Offset     int         `json:"offset"`
		Previous   interface{} `json:"previous"`
		TotalCount int         `json:"total_count"`
	} `json:"meta"`
	Objects []struct {
		CreatedAt      time.Time `json:"created_at"`
		Description    string    `json:"description"`
		DocumentName   string    `json:"document_name"`
		DocumentTypeID string    `json:"document_type_id"`
		Information    []struct {
			FieldName    string `json:"field_name"`
			FieldType    string `json:"field_type"`
			FriendlyName string `json:"friendly_name"`
			MaxLength    int    `json:"max_length,omitempty"`
			MinLength    int    `json:"min_length,omitempty"`
		} `json:"information"`
		ProofRequired interface{} `json:"proof_required"`
	} `json:"objects"`
}

func (service *ComplianceDocumentTypeService) Get(docId string) (response *GetComplianceDocumentTypeResponse, err error) {
	req, err := service.client.NewRequest("GET", nil, "ComplianceDocumentType/%s/", docId)
	response = &GetComplianceDocumentTypeResponse{}
	err = service.client.ExecuteRequest(req, response)
	return
}

func (service *ComplianceDocumentTypeService) List(params EndUserListParams) (response *ListComplianceDocumentTypeResponse, err error) {
	request, err := service.client.NewRequest("GET", params, "ComplianceDocumentType/")
	if err != nil {
		return
	}
	response = &ListComplianceDocumentTypeResponse{}
	err = service.client.ExecuteRequest(request, response)
	return
}
