package database

import (
	"fmt"

	"github.com/angel-lizander/questions-answers/models"
	"github.com/stretchr/testify/mock"
)

type MockQuestionClient struct {
	mock.Mock
}

func (m *MockQuestionClient) Insert(questions models.Question) (models.Question, error) {
	args := m.Called(questions)
	return args.Get(0).(models.Question), args.Error(1)
}

func (m *MockQuestionClient) Update(id string, update interface{}) (models.QuestionUpdate, error) {
	args := m.Called(id, update)
	return args.Get(0).(models.QuestionUpdate), args.Error(1)
}

func (m *MockQuestionClient) Delete(id string) (models.QuestionDelete, error) {
	args := m.Called(id)
	return args.Get(0).(models.QuestionDelete), args.Error(1)
}

func (m *MockQuestionClient) Get(id string) (models.Question, error) {
	fmt.Println("call get mock function")
	args := m.Called(id)
	return args.Get(0).(models.Question), args.Error(1)
}

func (m *MockQuestionClient) Search(filter interface{}) ([]models.Question, error) {
	args := m.Called(filter)
	return args.Get(0).([]models.Question), args.Error(1)
}
