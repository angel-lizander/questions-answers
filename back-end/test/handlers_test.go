package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/angel-lizander/questions-answers/apis"

	"github.com/stretchr/testify/assert"
)

func TestGetQuestions(t *testing.T) {
	client := &apis.MockHttpClient{}
	client.On("Get", URL).Return([]byte(`[{"QuestionAnswer":{}}]`), nil)

	req, _ := http.NewRequest("GET", "/todos", nil)
	rec := httptest.NewRecorder()
	h := http.HandlerFunc(GetQuestions(client))
	h.ServeHTTP(rec, req)

	assert.Contains(t, rec.Body.String(), `"QuestionAnswer":{}`)
	client.AssertExpectations(t)
}
