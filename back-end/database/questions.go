package database

import (
	"context"
	"encoding/json"

	"github.com/angel-lizander/questions-answers/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type QuestionInterface interface {
	Insert(models.Question) (models.Question, error)
	Update(string, interface{}) (models.QuestionUpdate, error)
	Delete(string) (models.QuestionDelete, error)
	Get(string) (models.Question, error)
	Search(interface{}) ([]models.Question, error)
}

type QuestionClient struct {
	Ctx context.Context
	Col *mongo.Collection
}

func (c *QuestionClient) Insert(docs models.Question) (models.Question, error) {
	question := models.Question{}

	res, err := c.Col.InsertOne(c.Ctx, docs)
	if err != nil {
		return question, err
	}
	id := res.InsertedID.(primitive.ObjectID).Hex()
	return c.Get(id)
}
func (c *QuestionClient) Update(id string, update interface{}) (models.QuestionUpdate, error) {
	result := models.QuestionUpdate{
		ModifiedCount: 0,
	}
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, err
	}

	question, err := c.Get(id)
	if err != nil {
		return result, err
	}

	var exist map[string]interface{}
	b, err := json.Marshal(question)
	if err != nil {
		return result, err
	}
	json.Unmarshal(b, &exist)

	change := update.(map[string]interface{})

	if len(change) == 0 {
		return result, nil
	}

	res, err := c.Col.UpdateOne(c.Ctx, bson.M{"_id": _id}, bson.M{"$set": change})
	if err != nil {
		return result, err
	}

	newQuestion, err := c.Get(id)
	if err != nil {
		return result, err
	}

	result.ModifiedCount = res.ModifiedCount
	result.Result = newQuestion
	return result, nil
}
func (c *QuestionClient) Delete(id string) (models.QuestionDelete, error) {
	result := models.QuestionDelete{
		DeletedCount: 0,
	}
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, err
	}

	res, err := c.Col.DeleteOne(c.Ctx, bson.M{"_id": _id})
	if err != nil {
		return result, err
	}
	result.DeletedCount = res.DeletedCount
	return result, nil
}
func (c *QuestionClient) Get(id string) (models.Question, error) {
	todo := models.Question{}

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return todo, err
	}

	err = c.Col.FindOne(c.Ctx, bson.M{"_id": _id}).Decode(&todo)
	if err != nil {
		return todo, err
	}

	return todo, nil
}
func (c *QuestionClient) Search(filter interface{}) ([]models.Question, error) {
	question := []models.Question{}
	if filter == nil {
		filter = bson.M{}
	}

	cursor, err := c.Col.Find(c.Ctx, filter)
	if err != nil {
		return question, err
	}

	for cursor.Next(c.Ctx) {
		row := models.Question{}
		cursor.Decode(&row)
		question = append(question, row)
	}

	return question, nil
}
