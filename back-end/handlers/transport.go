package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/angel-lizander/questions-answers/database"
	"github.com/angel-lizander/questions-answers/models"
	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
)

// Endpoint for the User service.

func MakeInsertEndpoint(s database.QuestionInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateQuestionRequest)
		resp, err := s.Insert(req.question)

		if err != nil {
			return GetQuestionResponse{Err: err}, nil
		}

		return CreateQuestionResponse{Question: resp, Err: err}, nil
	}

}

func MakeUpdateEndpoint(s database.QuestionInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateQuestionRequest)

		if req.ID == "" {
			return GetQuestionResponse{Err: errors.New("Parameter ID is required ")}, nil

		}

		resp, err := s.Update(req.ID, req.question)

		if err != nil {
			return GetQuestionResponse{Err: err}, nil
		}

		return UpdateQuestionResponse{Question: resp, Err: err}, nil
	}

}

func MakeGetEndpoint(s database.QuestionInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetQuestionRequest)

		var filter interface{}
		query := req.Query

		if query != "" {
			err := json.Unmarshal([]byte(query), &filter)
			if err != nil {
				return GetQuestionResponse{Err: err}, nil
			}
		}

		res, err := s.Search(filter)
		if err != nil {
			return GetQuestionResponse{Err: err}, nil
		}
		return GetQuestionResponse{Questions: res, Err: err}, nil

	}

}

func MakeGetByIdEndpoint(s database.QuestionInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetQuestionByIdRequest)

		resp, err := s.Get(req.id)

		if err != nil {
			return GetQuestionResponse{Err: err}, nil
		}

		return CreateQuestionResponse{Question: resp, Err: err}, nil

	}

}

func MakeDeleteEndpoint(s database.QuestionInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetQuestionByIdRequest)

		resp, err := s.Delete(req.id)

		if resp.DeletedCount == 0 {
			return DeleteQuestionResponse{Err: errors.New("Parameter ID no found")}, nil

		}

		if err != nil {
			return DeleteQuestionResponse{Err: err}, nil
		}
		return DeleteQuestionResponse{Question: resp, Err: err}, nil

	}

}

func DecodeGetAllQuestionsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req GetAllQuestionsRequest

	return req, nil
}

func DecodeGetRequest(_ context.Context, r *http.Request) (interface{}, error) {

	var req GetQuestionRequest
	req = GetQuestionRequest{
		Query: r.FormValue("q"),
	}

	return req, nil
}

func DecodeGetByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {

	var req GetQuestionByIdRequest
	params := mux.Vars(r)
	req = GetQuestionByIdRequest{
		id: params["id"],
	}
	return req, nil
}

func DecodePostRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req CreateQuestionRequest

	if err := json.NewDecoder(r.Body).Decode(&req.question); err != nil {
		return nil, err
	}
	return req, nil
}

func DecodeUpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req UpdateQuestionRequest
	params := mux.Vars(r)
	var question interface{}
	reqBody, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &question)
	if err != nil {
		return nil, err

	}

	req.ID = params["id"]
	req.question = question

	return req, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	log.Println("THIS IS THE RESPONDE", response)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type (
	CreateQuestionRequest struct {
		question models.Question
	}
	GetAllQuestionsRequest struct {
	}

	GetQuestionRequest struct {
		Query string
	}
	GetQuestionByIdRequest struct {
		id string
	}
	GetQuestionResponse struct {
		Questions []models.Question `json:"Questions"`
		Err       error
	}

	UpdateQuestionRequest struct {
		ID       string `json:"id"`
		question interface{}
	}

	CreateQuestionResponse struct {
		Question models.Question `json:"Question"`
		Err      error
	}
	DeleteQuestionResponse struct {
		Question models.QuestionDelete
		Err      error
	}
	UpdateQuestionResponse struct {
		Question models.QuestionUpdate
		Err      error
	}
)
