package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/angel-lizander/questions-answers/database"
	"github.com/angel-lizander/questions-answers/handlers"
	"github.com/angel-lizander/questions-answers/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

func TestInsertQuestion(t *testing.T) {
	client := &database.MockQuestionClient{}
	tests := map[string]struct {
		payload      string
		expectedCode int
	}{
		"should return 200": {
			payload:      `{"userId":1,"title":"learning golang","completed":false}`,
			expectedCode: 200,
		},
		"should return 400": {
			payload:      "invalid json string",
			expectedCode: 400,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			client.On("Insert", mock.Anything).Return(models.Question{}, nil)
			req, _ := http.NewRequest("POST", "/questions", strings.NewReader(test.payload))
			rec := httptest.NewRecorder()

			r := gin.Default()
			r.POST("/questions", handlers.InsertQuestion(client))
			r.ServeHTTP(rec, req)

			if test.expectedCode == 200 {
				client.AssertExpectations(t)
			} else {
				client.AssertNotCalled(t, "Insert")
			}
		})
	}
}
