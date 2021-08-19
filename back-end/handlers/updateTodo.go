package handlers

import (
	"net/http"

	"github.com/angel-lizander/questions-answers/database"

	"github.com/gin-gonic/gin"
)

func UpdateQuestions(db database.QuestionInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var question interface{}
		id := c.Param("id")
		err := c.BindJSON(&question)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		res, err := db.Update(id, question)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}
