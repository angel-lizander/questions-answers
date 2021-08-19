package handlers

import (
	"net/http"

	"github.com/angel-lizander/questions-answers/database"
	"github.com/angel-lizander/questions-answers/models"

	"github.com/gin-gonic/gin"
)

func InsertQuestion(db database.QuestionInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		question := models.Question{}
		err := c.BindJSON(&question)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		res, err := db.Insert(question)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}
