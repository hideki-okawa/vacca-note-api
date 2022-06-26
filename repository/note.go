package repository

import (
	"database/sql"

	"github.com/Okaki030/vacca-note-server/apperr"
	"github.com/Okaki030/vacca-note-server/log"
	"github.com/Okaki030/vacca-note-server/model"
)

type NoteDBRepository struct {
	db *sql.DB
}

func NewNoteDBRepository(db *sql.DB) NoteDBRepository {
	return NoteDBRepository{
		db: db,
	}
}

func (nr NoteDBRepository) GetNotes() (*sql.Rows, error) {
	log.Debugf("start repository GetNotes")

	rows, err := nr.db.Query(`
	SELECT 
		id, gender, age, vaccine_type, number_of_vaccination, max_temperature,
		log, remarks, good_count, created_at, second_vaccine_type
		FROM notes 
		order by id desc
	`)
	if err != nil {
		return nil, &apperr.ApplicationError{Code: "R001", Err: err}
	}

	return rows, nil
}

func (nr NoteDBRepository) GetReccomendNotes(searchConditions map[string]string) (*sql.Rows, error) {
	log.Debugf("start repository GetNotes")

	var rows *sql.Rows
	var err error
	if searchConditions["age"] != "" {
		rows, err = nr.db.Query(`
		SELECT 
			id, gender, age, vaccine_type, number_of_vaccination, max_temperature,
			log, remarks, good_count, created_at, second_vaccine_type
			FROM notes
			where id!=?
			and vaccine_type=?
			and number_of_vaccination=?
			order by abs(?-age), id desc
			limit 4
		`, searchConditions["id"], searchConditions["vaccineType"], searchConditions["numberOfVaccination"], searchConditions["age"])
		if err != nil {
			return nil, &apperr.ApplicationError{Code: "R002", Err: err}
		}
	} else {
		rows, err = nr.db.Query(`
		SELECT 
			id, gender, age, vaccine_type, second_vaccine_type, number_of_vaccination, max_temperature,
			log, remarks, good_count, created_at
			FROM notes
			where id!=?
			and vaccine_type=?
			and number_of_vaccination=?
			order by id desc
			limit 4
		`, searchConditions["id"], searchConditions["vaccineType"], searchConditions["numberOfVaccination"])
		if err != nil {
			return nil, &apperr.ApplicationError{Code: "R003", Err: err}
		}
	}

	return rows, nil
}

func (nr NoteDBRepository) GetAnalysisTemperature(vaccineType string, numberOfVaccination string) (*sql.Rows, error) {
	log.Debugf("start repository GetAnalysisTemperature")

	rows, err := nr.db.Query(`
	SELECT 
		max_temperature, count(*) 
		FROM notes
		WHERE vaccine_type=? AND number_of_vaccination=?
		GROUP BY vaccine_type, number_of_vaccination, max_temperature;
	`, vaccineType, numberOfVaccination)
	if err != nil {
		return nil, &apperr.ApplicationError{Code: "R007", Err: err}
	}

	return rows, nil
}

func (nr NoteDBRepository) GetNote(id string) (*sql.Rows, error) {
	log.Debugf("start repository GetNote")

	rows, err := nr.db.Query(`
	SELECT 
		*
	FROM notes 
	where id=?
	`, id)
	if err != nil {
		return nil, &apperr.ApplicationError{Code: "R004", Err: err}
	}

	return rows, nil
}

func (nr NoteDBRepository) PostNote(note model.Note) (int, error) {
	log.Debugf("start repository PostNote")

	result, err := nr.db.Exec(`
	INSERT INTO 
		notes (name, gender, age, vaccine_type, second_vaccine_type, number_of_vaccination, 
			max_temperature, log, remarks, good_count,created_at, updated_at)
		values (?,?,?,?,?,?,?,?,?,0,now(),now())`,
		note.Name, note.Gender, note.Age, note.VaccineType,
		note.SecondVaccineType, note.NuberOfVaccination, note.MaxTemperature, note.Log, note.Remarks)
	if err != nil {
		return 0, &apperr.ApplicationError{Code: "R005", Err: err}
	}

	insertID, err := result.LastInsertId()
	if err != nil {
		return 0, &apperr.ApplicationError{Code: "R006", Err: err}
	}

	return int(insertID), nil
}
