package models

import (
	"github.com/jmoiron/sqlx"
	"gopkg.in/go-playground/validator.v9"
)

//Applicant Applicant
type Applicant struct {
	ApplicantID      *string `json:"applicantId"`
	ApplicantName    *string `json:"applicantName"`
	ApplicantAddress *string `json:"applicantAddress"`
	ApplicantPhone   *string `json:"applicantPhone"`
	ApplicantCell    *string `json:"applicantCell"`
}

//CreateApplicantWithTransaction CreateApplicantWithTransaction
func (a *Applicant) CreateApplicantWithTransaction(tx *sqlx.Tx) (*sqlx.Row, error) {

	query := `INSERT INTO applicants (applicantName, applicantAddress, applicantPhone, applicantCell) VALUES 
	($1, $2, $3, $4)`
	stmt, err := tx.Preparex(query)
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRowx(a.ApplicantName, a.ApplicantAddress, a.ApplicantPhone, a.ApplicantCell)
	return row, nil

}

//CreateApplicant CreateApplicant
func (a *Applicant) CreateApplicant(db *sqlx.DB) (*sqlx.Row, error) {

	query := `INSERT INTO applicants
	 ( 
		 applicantName, applicantAddress, applicantPhone, applicantCell
	) VALUES 
	(
		$1, $2, $3, $4
	)`
	stmt, err := db.Preparex(query)
	if err != nil {
		return nil, err
	}
	stmt.Close()

	row := stmt.QueryRowx(a.ApplicantName, a.ApplicantAddress, a.ApplicantPhone, a.CreateApplicant)
	return row, nil

}

//Validate Validate
func (a *Applicant) Validate(v *validator.Validate) error {
	return v.Struct(a)
}

//FindByID FindById
func (a *Applicant) FindByID(db *sqlx.DB, id *string) *sqlx.Row {
	return db.QueryRowx("SELECT * FROM applicants WHERE applicantId = $1", id)
}
