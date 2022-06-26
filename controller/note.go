package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Okaki030/vacca-note-server/apperr"
	"github.com/Okaki030/vacca-note-server/auth"
	"github.com/Okaki030/vacca-note-server/db"
	"github.com/Okaki030/vacca-note-server/log"
	"github.com/Okaki030/vacca-note-server/model"
	repository "github.com/Okaki030/vacca-note-server/repository"
	"github.com/Okaki030/vacca-note-server/service"
	"github.com/Okaki030/vacca-note-server/utils"
	"github.com/gorilla/mux"
)

func GetNotes(w http.ResponseWriter, req *http.Request) {
	log.Debugf("start controller GetNotes")

	if req.Method == "OPTIONS" {
		utils.SendSuccess(w, nil)
		return
	}

	utils.EnableCORS(w)

	err := auth.CheckJWTToken(req)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, &apperr.ApplicationError{Code: "C001", Err: err})
		return
	}
	defer db.Close()

	ns := service.NewNoteService(repository.NewNoteDBRepository(db))
	notes, err := ns.GetNotes()
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err)
		return
	}

	utils.SendSuccess(w, notes)
}

func GetReccomendNotes(w http.ResponseWriter, req *http.Request) {
	log.Debugf("start controller GetReccomendNotes")

	if req.Method == "OPTIONS" {
		utils.SendSuccess(w, nil)
		return
	}

	utils.EnableCORS(w)

	err := auth.CheckJWTToken(req)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err)
		return
	}

	searchConditions := map[string]string{
		"id":                  "",
		"vaccineType":         "",
		"numberOfVaccination": "",
		"age":                 "",
	}

	searchConditions["id"] = req.URL.Query().Get("id")
	if searchConditions["id"] == "" {
		utils.SendError(w, http.StatusInternalServerError, errors.New("id is not specified"))
		return
	}

	searchConditions["vaccineType"] = req.URL.Query().Get("vaccine_type")
	if searchConditions["vaccineType"] == "" {
		utils.SendError(w, http.StatusInternalServerError, errors.New("vaccine_type is not specified"))
		return
	}

	searchConditions["numberOfVaccination"] = req.URL.Query().Get("number_of_vaccination")
	if searchConditions["numberOfVaccination"] == "" {
		utils.SendError(w, http.StatusInternalServerError, errors.New("number_of_vaccination is not specified"))
		return
	}

	searchConditions["age"] = req.URL.Query().Get("age")

	db, err := db.Connect()
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	ns := service.NewNoteService(repository.NewNoteDBRepository(db))
	notes, err := ns.GetReccomendNotes(searchConditions)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err)
		return
	}

	utils.SendSuccess(w, notes)
}

func GetAnalysisTemperature(w http.ResponseWriter, req *http.Request) {
	log.Debugf("start controller GetAnalysisTemperature")

	if req.Method == "OPTIONS" {
		utils.SendSuccess(w, nil)
		return
	}

	utils.EnableCORS(w)

	err := auth.CheckJWTToken(req)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	ns := service.NewNoteService(repository.NewNoteDBRepository(db))
	temperaturList, err := ns.GetAnalysisTemperature()
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err)
		return
	}

	utils.SendSuccess(w, temperaturList)
}

func GetNote(w http.ResponseWriter, req *http.Request) {
	log.Debugf("start controller GetNote")

	if req.Method == "OPTIONS" {
		utils.SendSuccess(w, nil)
		return
	}

	w = utils.EnableCORS(w)

	err := auth.CheckJWTToken(req)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err)
		return
	}

	// IDの取得
	vars := mux.Vars(req) //パスパラメータ取得

	db, err := db.Connect()
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	ns := service.NewNoteService(repository.NewNoteDBRepository(db))
	notes, err := ns.GetNote(vars["id"])
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err)
		return
	}

	utils.SendSuccess(w, notes)
}

func PostNote(w http.ResponseWriter, req *http.Request) {
	log.Debugf("start controller PostNote")

	// preflight requestを回避するために必要
	// https://qiita.com/nnishimura/items/1f156f05b26a5bce3672
	if req.Method == "OPTIONS" {
		utils.SendSuccess(w, nil)
		return
	}

	utils.EnableCORS(w)

	err := auth.CheckJWTToken(req)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	var note model.Note
	json.NewDecoder(req.Body).Decode(&note)
	log.Infof("request body: %+v", note)

	ns := service.NewNoteService(repository.NewNoteDBRepository(db))
	insertID, err := ns.PostNote(note)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err)
		return
	}

	utils.SendSuccess(w, insertID)
}
