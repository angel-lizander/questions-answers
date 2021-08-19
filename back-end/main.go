package main

import (
	"context"

	"github.com/angel-lizander/questions-answers/config"
	"github.com/angel-lizander/questions-answers/database"
	"github.com/angel-lizander/questions-answers/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

	r := gin.Default()
	r.Use(cors.Default())
	questions := r.Group("/questions")
	{
		questions.GET("/", handlers.SearchQuestions(client))
		questions.GET("/:id", handlers.GetQuestion(client))
		questions.POST("/", handlers.InsertQuestion(client))
		questions.PATCH("/:id", handlers.UpdateQuestions(client))
		questions.DELETE("/:id", handlers.DeleteQuestion(client))
	}

	r.Run(":8080")
}
