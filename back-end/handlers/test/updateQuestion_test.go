package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/angel-lizander/questions-answers/database"
	"github.com/angel-lizander/questions-answers/handlers"
	"github.com/angel-lizander/questions-answers/models"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

func TestUpdateQuestions(t *testing.T) {
	client := &database.MockQuestionClient{}
	id := primitive.NewObjectID().Hex()

	tests := map[string]struct {
		id           string
		payload      string
		expectedCode int
	}{
		"should return 200": {
			id:           id,
			payload:      `{"completed": true}`,
			expectedCode: 200,
		},
		"should return 404": {
			id:           "",
			expectedCode: 404,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if test.expectedCode == 200 {
				client.On("Update", test.id, mock.Anything).Return(models.QuestionUpdate{}, nil)
			}
			req, _ := http.NewRequest("PATCH", "/questions/"+test.id, strings.NewReader(test.payload))
			rec := httptest.NewRecorder()

			r := gin.Default()
			r.PATCH("/questions/:id", handlers.UpdateQuestions(client))
			r.ServeHTTP(rec, req)

			if test.expectedCode == 200 {
				client.AssertExpectations(t)
			} else {
				client.AssertNotCalled(t, "Update")
			}
		})
	}
}
