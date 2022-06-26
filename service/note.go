package service

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/Okaki030/vacca-note-server/apperr"
	"github.com/Okaki030/vacca-note-server/log"
	"github.com/Okaki030/vacca-note-server/model"
	"github.com/Okaki030/vacca-note-server/repository"
)

type NoteService struct {
	nr repository.NoteRepository
}

func NewNoteService(nr repository.NoteRepository) NoteService {
	return NoteService{
		nr: nr,
	}
}

func (ns NoteService) GetNotes() ([]model.Note, error) {
	log.Debugf("start service GetNotes")

	rows, err := ns.nr.GetNotes()
	if err != nil {
		return nil, err
	}

	var notes []model.Note
	for rows.Next() {
		var note model.Note
		var log, remarks string
		err := rows.Scan(
			&note.ID, &note.Gender, &note.Age, &note.VaccineType,
			&note.NuberOfVaccination, &note.MaxTemperature,
			&log, &remarks, &note.GoodCount, &note.CreatedAt, &note.SecondVaccineType,
		)
		if err != nil {
			return nil, &apperr.ApplicationError{Code: "S002", Err: err}
		}

		// logは先頭の3行を取得する or 先頭の180バイトを抽出する
		note.Log = formatParagraph(log)

		// remarksは先頭の3行を取得する or 先頭の180バイトを抽出する
		note.Remarks = formatParagraph(remarks)

		notes = append(notes, note)
	}

	return notes, nil
}

func (ns NoteService) GetReccomendNotes(searchConditions map[string]string) ([]model.Note, error) {
	log.Debugf("start service GetReccomendNotes")

	rows, err := ns.nr.GetReccomendNotes(searchConditions)
	if err != nil {
		return nil, err
	}

	var notes []model.Note
	for rows.Next() {
		var note model.Note
		var log, remarks string
		err := rows.Scan(
			&note.ID, &note.Gender, &note.Age, &note.VaccineType,
			&note.NuberOfVaccination, &note.MaxTemperature, &log,
			&remarks, &note.GoodCount, &note.CreatedAt, &note.SecondVaccineType,
		)
		if err != nil {
			return nil, &apperr.ApplicationError{Code: "S003", Err: err}
		}

		// logは先頭の3行を取得する or 先頭の180バイトを抽出する
		note.Log = formatParagraph(log)

		// remarksは先頭の3行を取得する or 先頭の180バイトを抽出する
		note.Remarks = formatParagraph(remarks)

		notes = append(notes, note)
	}

	return notes, nil
}

func formatParagraph(originalText string) string {
	slice := strings.Split(originalText, "\n")
	if len(slice) > 3 {
		return strings.Join(slice[:3], "\n") + "..."
	} else {
		if len([]rune(originalText)) > 100 {
			return string([]rune(originalText)[:100]) + "..."
		} else {
			return originalText
		}
	}
}

func (ns NoteService) GetAnalysisTemperature() ([]model.TemperatureData, error) {
	log.Debugf("start service GetAnalysisTemperature")

	// 0: ワクチンの種類(F→ファイザー M→モデルナ)
	// 1: 接種回数
	targetSlice := [][]string{
		{"F", "2"},
		{"M", "2"},
		{"F", "1"},
		{"M", "1"},
	}

	var temperatureList []model.TemperatureData
	for _, target := range targetSlice {
		var temperatureData model.TemperatureData

		// グラフタイトルの生成
		var vaccineName string
		if target[0] == "F" {
			vaccineName = "ファイザー"
		} else {
			vaccineName = "モデルナ"
		}
		temperatureData.Name = fmt.Sprintf("%s(%s回目)接種後の体温", vaccineName, target[1])

		// 体温リストの生成
		var temperatureCountList []model.TemperatureCount
		rows, err := ns.nr.GetAnalysisTemperature(target[0], target[1])
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var temperatureCount model.TemperatureCount
			err := rows.Scan(&temperatureCount.Num, &temperatureCount.Count)
			if err != nil {
				return nil, &apperr.ApplicationError{Code: "S016", Err: err}
			}
			temperatureCountList = append(temperatureCountList, temperatureCount)
		}
		temperatureData.List = temperatureCountList

		temperatureList = append(temperatureList, temperatureData)
	}

	return temperatureList, nil
}

func (ns NoteService) GetNote(id string) (model.Note, error) {
	log.Debugf("start service GetNote")

	rows, err := ns.nr.GetNote(id)
	if err != nil {
		return model.Note{}, err
	}

	var note model.Note
	for rows.Next() {
		err := rows.Scan(
			&note.ID, &note.Name, &note.Gender, &note.Age, &note.VaccineType,
			&note.NuberOfVaccination, &note.MaxTemperature, &note.Log,
			&note.Remarks, &note.GoodCount, &note.CreatedAt, &note.UpdatedAt, &note.SecondVaccineType,
		)
		if err != nil {
			return model.Note{}, &apperr.ApplicationError{Code: "S004", Err: err}
		}
	}

	if note.ID == 0 {
		return model.Note{}, &apperr.ApplicationError{Code: "S005", Err: errors.New("Note not found")}
	}

	return note, nil
}

func (ns NoteService) PostNote(note model.Note) (int, error) {
	log.Debugf("start service PostNote")

	// バリデーション
	// 名前
	if note.Name == "" {
		note.Name = "匿名"
	}

	nameWordCount := len([]rune(note.Name))
	if nameWordCount > 50 {
		return 0, &apperr.ApplicationError{Code: "S013", Err: errors.New("Please enter a name that is less than 50 characters long")}
	}

	// 性別
	if note.Gender != "0" && note.Gender != "1" && note.Gender != "2" && note.Gender != "9" {
		return 0, &apperr.ApplicationError{Code: "S006", Err: errors.New("Please enter the correct gender")}
	}

	// 年齢
	age, err := strconv.Atoi(note.Age)
	if err != nil {
		return 0, &apperr.ApplicationError{Code: "S012", Err: err}

	}
	if age <= 0 || age > 9 {
		return 0, &apperr.ApplicationError{Code: "S007", Err: errors.New("Please enter the correct age")}
	}

	// ワクチン種別
	if note.VaccineType != "M" && note.VaccineType != "F" && note.VaccineType != "A" {
		return 0, &apperr.ApplicationError{Code: "S008", Err: errors.New("Please enter the correct vaccine_type")}
	}

	if note.SecondVaccineType != "M" && note.SecondVaccineType != "F" && note.SecondVaccineType != "A" && note.SecondVaccineType != "" {
		return 0, &apperr.ApplicationError{Code: "S017", Err: errors.New("Please enter the correct second_vaccine_type")}
	}

	// ワクチンの接種回数
	if note.NuberOfVaccination <= 0 || note.NuberOfVaccination > 4 {
		return 0, &apperr.ApplicationError{Code: "S009", Err: errors.New("Please enter the correct number of vaccination")}
	}

	// 最高体温
	maxTemperature, err := strconv.Atoi(note.MaxTemperature)
	if err != nil {
		return 0, &apperr.ApplicationError{Code: "S010", Err: err}
	}
	if maxTemperature <= 0 || maxTemperature > 10 {
		return 0, &apperr.ApplicationError{Code: "S001", Err: errors.New("Please enter the correct maxTemperature")}
	}

	// 経過記録
	logWordCount := len([]rune(note.Log))
	if logWordCount > 3000 {
		return 0, &apperr.ApplicationError{Code: "S014", Err: errors.New("Please enter a log that is less than 3000 characters long")}
	}

	// TODO:自由コメントのバリデーション
	remarksWordCount := len([]rune(note.Remarks))
	if remarksWordCount > 3000 {
		return 0, &apperr.ApplicationError{Code: "S015", Err: errors.New("Please enter a remarks that is less than 3000 characters long")}
	}

	insertID, err := ns.nr.PostNote(note)
	if err != nil {
		return 0, &apperr.ApplicationError{Code: "S011", Err: err}
	}

	return insertID, nil
}
