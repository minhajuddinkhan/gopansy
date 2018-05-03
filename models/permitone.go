package models

import (
	_ "database/sql"

	"gopkg.in/go-playground/validator.v9"

	_ "github.com/lib/pq"

	sqlx "github.com/jmoiron/sqlx"
)

//PermitOneForm PermitOneForm
type PermitOneForm struct {
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
}

//CreateWithTransaction CreateWithTransaction
func (p *PermitOneForm) CreateWithTransaction(tx *sqlx.Tx) (*sqlx.Rows, error) {

	query := `INSERT INTO permitoneforms (
	isVerified, tokenNumber, diaryDate, diaryNumber, 
	importerId, applicantId, nameOfProduct, quantityOfProduct, countOfProduct, countryOrigin,
	importedFrom, numberOfConsignment, exporterName, exporterAddress, destinationPortName,
	meansOfTransport, portOfEntry, purposeOfImport, departmentUseOnly)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19) RETURNING permitOneId`

	stmt, err := tx.Preparex(query)
	if err != nil {
		return nil, err
	}
	return stmt.Queryx(
		p.IsVerified,
		p.TokenNumber,
		p.DiaryDate,
		p.DiaryNumber,
		p.ImporterID,
		p.ApplicantID,
		p.NameOfProduct,
		p.QuantityOfProduct,
		p.CountOfProduct,
		p.CountryOrigin,
		p.ImportedFrom,
		p.NumberOfConsignment,
		p.ExporterName,
		p.ExporterAddress,
		p.DestinationPortName,
		p.MeansOfTransport,
		p.PortOfEntry,
		p.PurposeOfImport,
		p.DepartmentUseOnly)

}

//CreatePermitOneForm CreatePermitOneForm
func (p *PermitOneForm) CreatePermitOneForm(db *sqlx.DB) (*sqlx.Row, error) {

	query := `
	INSERT INTO permitOneForms (
	isVerified, tokenNumber, diaryDate, diaryNumber, 
	importerId, applicantId, nameOfProduct, quantityOfProduct, countOfProduct, countryOrigin,
	importedFrom, numberOfConsignment, exporterName, exporterAddress, destinationPortName,
	meansOfTransport, portOfEntry, purposeOfImport, departmentUseOnly)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19) RETURNING permitOneId`

	stmt, err := db.Preparex(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRowx(
		p.IsVerified,
		p.TokenNumber,
		p.DiaryDate,
		p.DiaryNumber,
		p.ImporterID,
		p.ApplicantID,
		p.NameOfProduct,
		p.QuantityOfProduct,
		p.CountOfProduct,
		p.CountryOrigin,
		p.ImportedFrom,
		p.NumberOfConsignment,
		p.ExporterName,
		p.ExporterAddress,
		p.DestinationPortName,
		p.MeansOfTransport,
		p.PortOfEntry,
		p.PurposeOfImport,
		p.DepartmentUseOnly)

	return row, nil

}

//Validate Validate
func (p *PermitOneForm) Validate(v *validator.Validate) error {
	return v.Struct(p)
}
