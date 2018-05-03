package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/minhajuddinkhan/gopansy/helpers"

	"github.com/jmoiron/sqlx"
	"github.com/minhajuddinkhan/gopansy/constants"

	"github.com/darahayes/go-boom"

	_ "database/sql"

	_ "github.com/lib/pq"
	"github.com/minhajuddinkhan/gopansy/models"
)

//CreatePermitOneForm CreatePermitOneForm
func CreatePermitOneForm(w http.ResponseWriter, r *http.Request) {

	payload := struct {
		models.PermitOneForm
		models.Applicant
		models.Importer
	}{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)

	var newApplicantID *string
	var newImporterID *string
	if err != nil {
		boom.BadRequest(w, "Invalid Request Body")
		return
	}

	err = helpers.Validate(payload)
	if err != nil {
		boom.BadRequest(w, err)
		return
	}

	db := r.Context().Value(constants.DbKey).(*sqlx.DB)

	tx := db.MustBegin()

	fmt.Println("payload.Applicant.ApplicantID", payload)
	applicant := models.Applicant{}
	row := applicant.FindByID(db, payload.Applicant.ApplicantID)
	row.StructScan(&applicant)

	if applicant.ApplicantID == nil {
		applicant = models.Applicant{
			ApplicantName:    payload.Applicant.ApplicantName,
			ApplicantAddress: payload.Applicant.ApplicantAddress,
		}

		row, err = applicant.CreateApplicantWithTransaction(tx)
		if err != nil {
			boom.Internal(w, "Cannot Create Applicant")
			return
		}
		row.StructScan(&newApplicantID)

		payload.Applicant.ApplicantID = newApplicantID
	} else {
		payload.Applicant.ApplicantID = applicant.ApplicantID
	}

	importer := models.Importer{}
	row = importer.FindByID(db, payload.Importer.ImporterID)
	row.StructScan(&importer)

	if importer.ImporterID != nil {
		importer = models.Importer{
			ImporterName:    payload.Importer.ImporterName,
			ImporterAddress: payload.Importer.ImporterAddress,
		}
		row, err = importer.CreateImporterWithTransaction(tx)
		if err != nil {
			boom.Internal(w, "Cannot Create Importert")
			return
		}
		row.StructScan(&newImporterID)
		payload.Importer.ImporterID = newImporterID
	} else {
		payload.Importer.ImporterID = importer.ImporterID
	}

	rows, err := payload.PermitOneForm.CreateWithTransaction(tx)
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
