package models

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

//Importer Importer
type Importer struct {
	ImporterID      *string `json:"importerId"`
	ImporterName    *string `json:"importerName"`
	ImporterAddress *string `json:"importerAddress"`
	ImporterPhone   *string `json:"importerPhone"`
	ImporterCell    *string `json:"importerCell"`
	ImporterFax     *string `json:"importerFax"`
	ImporterEmail   *string `json:"importerEmail"`
}

//CreateImporterWithTransaction CreateImporterWithTransaction
func (i *Importer) CreateImporterWithTransaction(tx *sqlx.Tx) (*sql.Row, error) {

	query := `INSERT INTO importers 
			(
				importerName, importerAddress, importerPhone, importerCell, importerFax, importerEmail
			) VALUES
			(
				$1, $2, $3, $4, $5, $6
			) RETURNING importerId`

	stmt, err := tx.Preparex(query)
	if err != nil {

		return nil, err
	}

	row := stmt.QueryRow(i.ImporterName, i.ImporterAddress, i.ImporterPhone, i.ImporterCell, i.ImporterFax, i.ImporterEmail)

	return row, err

}

//CreateImporter CreateImporter
func (i *Importer) CreateImporter(db *sqlx.DB) (*sqlx.Row, error) {

	query := `INSERT INTO importers 
			(
				importerName, importerAddress, importerPhone, importerCell, importerFax, importerEmail
			) VALUES
			(
				$1, $2, $3, $4, $5, $6
			)`

	stmt, err := db.Preparex(query)
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRowx(i.ImporterName, i.ImporterAddress, i.ImporterPhone, i.ImporterCell, i.ImporterFax, i.ImporterEmail)
	return row, err

}

//FindByID FindByID
func (i *Importer) FindByID(db *sqlx.DB, id *string) *sqlx.Row {

	return db.QueryRowx("SELECT * FROM importers WHERE importerId = $1", id)
}
