package utils

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Okaki030/vacca-note-server/apperr"
	"github.com/Okaki030/vacca-note-server/log"
)

func SendError(w http.ResponseWriter, status int, err error) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // cros
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.WriteHeader(status)

	log.Errorf(err.Error())

	// 独自エラーの時にエラーの内容を絞る
	switch e := err.(type) {
	case *apperr.ApplicationError:
		e.Message = apperr.ReturnErrorMessage(e.Code)
		json.NewEncoder(w).Encode(e.ResponseError())
	default:
		json.NewEncoder(w).Encode(errors.New("システムエラーが発生しました。管理者に連絡してください。").Error())
	}

	return
}

func SendSuccess(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // cros
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(data)
	return
}

func EnableCORS(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Access-Control-Allow-Origin", "*") // cros
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	return w
}
