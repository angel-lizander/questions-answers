package main

import (
	"context"
	"log"
	"net/http"

	"github.com/angel-lizander/questions-answers/config"
	"github.com/angel-lizander/questions-answers/database"
	"github.com/angel-lizander/questions-answers/handlers"
	"github.com/rs/cors"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func main() {

	conf := config.GetConfig()
	ctx := context.TODO()
	db := database.ConnectDB(ctx, conf.Mongo)
	collection := db.Collection(conf.Mongo.Collection)

	client := &database.QuestionClient{
		Col: collection,
		Ctx: ctx,
	}

	GetQuestionHandler := httptransport.NewServer(
		handlers.MakeGetEndpoint(client),
		handlers.DecodeGetRequest,
		handlers.EncodeResponse,
	)

	CreateQuestionHandler := httptransport.NewServer(
		handlers.MakeInsertEndpoint(client),
		handlers.DecodePostRequest,
		handlers.EncodeResponse,
	)

	GetByIdQuestionHandler := httptransport.NewServer(
		handlers.MakeGetByIdEndpoint(client),
		handlers.DecodeGetByIdRequest,
		handlers.EncodeResponse,
	)

	DeleteQuestionHandler := httptransport.NewServer(
		handlers.MakeDeleteEndpoint(client),
		handlers.DecodeGetByIdRequest,
		handlers.EncodeResponse,
	)

	UpdateQuestionHandler := httptransport.NewServer(
		handlers.MakeUpdateEndpoint(client),
		handlers.DecodeUpdateRequest,
		handlers.EncodeResponse,
	)

	r := mux.NewRouter()

	api := r.PathPrefix("/questions").Subrouter()
	api.Handle("/", GetQuestionHandler).Methods(http.MethodGet)
	api.Path("/").Queries("q", "").Handler(GetQuestionHandler).Name("SearchQuestions")
	api.Handle("/", CreateQuestionHandler).Methods(http.MethodPost)
	api.Handle("/{id}", GetByIdQuestionHandler).Methods(http.MethodGet)
	api.Handle("/{id}", UpdateQuestionHandler).Methods(http.MethodPatch)
	api.Handle("/{id}", DeleteQuestionHandler).Methods(http.MethodDelete)

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

}
