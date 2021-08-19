package handlers

import (
	"net/http"

	"github.com/angel-lizander/questions-answers/database"

	"github.com/gin-gonic/gin"
)

func GetQuestion(db database.QuestionInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		res, err := db.Get(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}