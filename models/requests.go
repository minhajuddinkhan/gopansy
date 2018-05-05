package models

//PermitOneFormCreateRequest PermitOneFormCreateRequest
type PermitOneFormCreateRequest struct {
	PermitOneID         *string `json:"permitOneId,omitempty"`
	IsVerified          *bool   `json:"isVerified" validate:"required"`
	TokenNumber         *string `json:"tokenNumber"`
	DiaryDate           *string `json:"diaryDate" validate:"required"`
	DiaryNumber         *string `json:"diaryNumber" validate:"required"`
	ImporterID          *string `json:"importerId"`
	ApplicantID         *string `json:"applicantId"`
	NameOfProduct       *string `json:"nameOfProduct" validate:"required"`
	QuantityOfProduct   *string `json:"quantityOfProduct" validate:"required"`
	CountOfProduct      *int32  `json:"countOfProduct" validate:"required"`
	CountryOrigin       *string `json:"countryOrigin" validate:"required"`
	ImportedFrom        *string `json:"importedFrom" validate:"required"`
	NumberOfConsignment *int32  `json:"numberOfConsignment" validate:"required"`
	ExporterName        *string `json:"exporterName" validate:"required"`
	ExporterAddress     *string `json:"exporterAddress" validate:"required"`
	DestinationPortName *string `json:"destinationPortName" validate:"required"`
	MeansOfTransport    *string `json:"meansOfTransport" validate:"required"`
	PortOfEntry         *string `json:"portOfEntry" validate:"required"`
	PurposeOfImport     *string `json:"purposeOfImport" validate:"required"`
	DepartmentUseOnly   *string `json:"departmentUseOnly"`
	ApplicantName       *string `json:"applicantName" validate:"required"`
	ApplicantAddress    *string `json:"applicantAddress"`
	ApplicantPhone      *string `json:"applicantPhone"`
	ApplicantCell       *string `json:"applicantCell"`
	ImporterName        *string `json:"importerName"`
	ImporterAddress     *string `json:"importerAddress"`
	ImporterPhone       *string `json:"importerPhone"`
	ImporterCell        *string `json:"importerCell"`
	ImporterFax         *string `json:"importerFax"`
	ImporterEmail       *string `json:"importerEmail"`
}
