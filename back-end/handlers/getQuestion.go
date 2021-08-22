package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/angel-lizander/questions-answers/database"
	"github.com/gorilla/mux"
)

func GetQuestion(db database.QuestionInterface, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	res, err := db.Get(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
