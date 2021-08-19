package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/angel-lizander/questions-answers/apis"
	"github.com/angel-lizander/questions-answers/models"
)

var URL = "https://raw.githubusercontent.com/angel-lizander/questions-answers/main/db.json"

func GetQuestions(client apis.HttpInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todos := []models.Question{}

		body, err := client.Get(URL)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}

		err = json.Unmarshal(body, &todos)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(todos)

	}
}
