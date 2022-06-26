package repository

import (
	"database/sql"

	"github.com/Okaki030/vacca-note-server/model"
)

type NoteRepository interface {
	GetNotes() (*sql.Rows, error)
	GetReccomendNotes(map[string]string) (*sql.Rows, error)
	GetNote(string) (*sql.Rows, error)
	PostNote(model.Note) (int, error)
	GetAnalysisTemperature(string, string) (*sql.Rows, error)
}
