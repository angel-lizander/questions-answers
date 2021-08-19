package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/angel-lizander/questions-answers/database"
	"github.com/angel-lizander/questions-answers/handlers"
	"github.com/angel-lizander/questions-answers/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetQuestions(t *testing.T) {
	client := &database.MockQuestionClient{}
	id := primitive.NewObjectID().Hex()

	tests := map[string]struct {
		id           string
		expectedCode int
	}{
		"should return 200": {
			id:           id,
			expectedCode: 200,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if test.expectedCode == 200 {
				client.On("Get", test.id).Return(models.Question{}, nil)
			}
			req, _ := http.NewRequest("GET", "/questions/"+test.id, nil)
			rec := httptest.NewRecorder()

			gin.SetMode(gin.ReleaseMode)
			r := gin.New()
			r.GET("/questions/:id", handlers.GetQuestion(client))
			r.ServeHTTP(rec, req)

			if test.expectedCode == 200 {
				client.AssertExpectations(t)
			} else {
				client.AssertNotCalled(t, "Get")
			}
		})
	}
}
