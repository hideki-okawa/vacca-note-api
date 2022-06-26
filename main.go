package main

import (
	"fmt"
	"net/http"

	"github.com/Okaki030/vacca-note-server/auth"
	"github.com/Okaki030/vacca-note-server/controller"
	"github.com/Okaki030/vacca-note-server/log"
	"github.com/gorilla/mux"
)

func main() {

	port := "80"

	router := mux.NewRouter()

	router.HandleFunc("/", healthCheck).Methods("GET", "OPTIONS")
	router.HandleFunc("/note", controller.PostNote).Methods("POST", "OPTIONS")
	router.HandleFunc("/note/{id}", controller.GetNote).Methods("GET", "OPTIONS")
	router.HandleFunc("/notes/recommend", controller.GetReccomendNotes).Methods("GET", "OPTIONS")
	router.HandleFunc("/notes/analysis/temperature", controller.GetAnalysisTemperature).Methods("GET", "OPTIONS")
	router.HandleFunc("/notes", controller.GetNotes).Methods("GET", "OPTIONS")

	router.HandleFunc("/auth", auth.GetJWTToken).Methods("GET")

	router.HandleFunc("/stats", getStats).Methods("GET")
	router.HandleFunc("/stats", deleteStats).Methods("DELETE")

	log.Infof("Server Running at port:%s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}

// healthCheck はヘルスチェック用の関数
func healthCheck(w http.ResponseWriter, req *http.Request) {
	log.Debugf("run healthCheck")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "Hello vacca-note!!")
}
