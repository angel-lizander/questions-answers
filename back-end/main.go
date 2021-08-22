package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/angel-lizander/questions-answers/config"
	"github.com/angel-lizander/questions-answers/database"
	"github.com/angel-lizander/questions-answers/models"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func SearchQuestions(w http.ResponseWriter, r *http.Request) {

	var filter interface{}

	query := r.FormValue("q")

	if query != "" {
		err := json.Unmarshal([]byte(query), &filter)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
	}

	res, err := client.Search(filter)
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

func GetQuestion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	res, err := client.Get(id)
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

func InsertQuestion(w http.ResponseWriter, r *http.Request) {

	question := models.Question{}
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.Unmarshal(reqBody, &question)
	res, err := client.Insert(question)

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

func DeleteQuestion(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id := params["id"]

	res, err := client.Delete(id)
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

func UpdateQuestions(w http.ResponseWriter, r *http.Request) {
	var question interface{}
	params := mux.Vars(r)
	id := params["id"]

	reqBody, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &question)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	res, err := client.Update(id, question)

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

var client = &database.QuestionClient{}

func main() {

	conf := config.GetConfig()
	ctx := context.TODO()
	db := database.ConnectDB(ctx, conf.Mongo)
	collection := db.Collection(conf.Mongo.Collection)

	client = &database.QuestionClient{
		Col: collection,
		Ctx: ctx,
	}

	r := mux.NewRouter()

	api := r.PathPrefix("/questions").Subrouter()
	api.HandleFunc("/", SearchQuestions).Methods(http.MethodGet)
	api.Path("/").Queries("q", "").HandlerFunc(SearchQuestions).Name("SearchQuestions")

	api.HandleFunc("/{id}", GetQuestion).Methods(http.MethodGet)
	api.HandleFunc("/", InsertQuestion).Methods(http.MethodPost)
	api.HandleFunc("/{id}", UpdateQuestions).Methods(http.MethodPatch)
	api.HandleFunc("/{id}", DeleteQuestion).Methods(http.MethodDelete)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PATCH"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	error := http.ListenAndServe(":8080", handler)

	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("Listening in port 8080")
	}

	/*questions := r.Group("/questions")
	{
		questions.GET("/", handlers.SearchQuestions(client))
		questions.GET("/:id", handlers.GetQuestion(client))
		questions.POST("/", handlers.InsertQuestion(client))
		questions.PATCH("/:id", handlers.UpdateQuestions(client))
		questions.DELETE("/:id", handlers.DeleteQuestion(client))
	}

	r.Run(":8080")*/
}
