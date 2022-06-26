package auth

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/Okaki030/vacca-note-server/apperr"
	"github.com/Okaki030/vacca-note-server/log"
	"github.com/Okaki030/vacca-note-server/utils"
	jwt "github.com/form3tech-oss/jwt-go"
)

type jwtResponse struct {
	Token string `json:"token"`
}

func GetJWTToken(w http.ResponseWriter, req *http.Request) {
	log.Debugf("start GetJWTToken")
	utils.EnableCORS(w)

	// headerのセット
	token := jwt.New(jwt.SigningMethodHS256)

	// 電子署名
	tokenString, err := token.SignedString([]byte(os.Getenv("SIGNINGKEY")))
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, &apperr.ApplicationError{Code: "A001", Err: err})
		return
	}

	jwtResponse := jwtResponse{
		Token: tokenString,
	}

	utils.SendSuccess(w, jwtResponse)
}

func CheckJWTToken(req *http.Request) error {
	authHeader := req.Header.Get("Authorization")

	if authHeader == "" {
		return &apperr.ApplicationError{Code: "A002", Err: errors.New("not found authorization header")}
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		return &apperr.ApplicationError{Code: "A003", Err: errors.New("invalid auth header")}
	}

	if headerParts[0] != "Bearer" {
		return &apperr.ApplicationError{Code: "A004", Err: errors.New("unauthorized - no bearer")}
	}

	_, err := jwt.Parse(headerParts[1], func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SIGNINGKEY")), nil
	})
	if err != nil {
		return &apperr.ApplicationError{Code: "A005", Err: err}
	}

	return nil
}
