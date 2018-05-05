package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/minhajuddinkhan/gopansy/constants"

	"github.com/imdario/mergo"

	boom "github.com/darahayes/go-boom"
	"github.com/minhajuddinkhan/gopansy/helpers"

	_ "database/sql"

	_ "github.com/lib/pq"
	"github.com/minhajuddinkhan/gopansy/models"
)

//CreatePermitOneForm CreatePermitOneForm
func CreatePermitOneForm(w http.ResponseWriter, r *http.Request) {

	applicantID := 0
	importerID := 0

	requestBody := models.PermitOneFormCreateRequest{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		boom.BadRequest(w, "Unable to parse request body")
		return
	}

	applicant := models.Applicant{
		ApplicantID:      requestBody.ApplicantID,
		ApplicantAddress: requestBody.ApplicantAddress,
		ApplicantCell:    requestBody.ApplicantAddress,
		ApplicantName:    requestBody.ApplicantName,
		ApplicantPhone:   requestBody.ApplicantPhone,
	}

	importer := models.Importer{
		ImporterID:   requestBody.ImporterID,
		ImporterName: requestBody.ImporterName,
	}

	applicantFound := models.Applicant{}
	importerFound := models.Importer{}

	db := r.Context().Value(constants.DbKey).(*sqlx.DB)
	tx := db.MustBegin()

	row := applicant.FindByID(db, requestBody.ApplicantID)
	row.StructScan(&applicantFound)

	if len(*applicantFound.ApplicantID) == 0 {
		row, err := applicant.CreateApplicantWithTransaction(tx)
		if err != nil {
			boom.BadData(w)
			return
		}
		row.StructScan(&applicantID)

	}

	row = importer.FindByID(db, requestBody.ImporterID)
	if err != nil {
		boom.BadData(w)
		return
	}

	row.StructScan(&importerFound)
	if len(*importerFound.ImporterID) == 0 {
		row, err := importer.CreateImporterWithTransaction(tx)
		if err != nil {
			boom.BadData(w)
			return
		}

		row.Scan(&importerID)
		fmt.Println("SCANNING STRUCT", importerID)
	}

	helpers.Respond(w, importerID)
	return

	permitOne := models.PermitOneForm{}

	mergo.Map(&permitOne, requestBody)

	rows, err := permitOne.CreateWithTransaction(tx)
	if err != nil {
		boom.Internal(w)
		return
	}

	rows.Close()
	err = tx.Commit()
	if err != nil {
		fmt.Println("CANNOT CREATE PERMIT ONE!!", err)

		boom.BadData(w, err)
		return
	}
	var permitOneID int32
	row.Scan(&permitOneID)
	if err != nil {
		boom.Internal(w, err)
		return
	}

	helpers.Respond(w, struct {
		Success     bool  `json:"success"`
		PermitOneID int32 `json:"permitOneId"`
	}{
		Success:     true,
		PermitOneID: permitOneID,
	})
}
