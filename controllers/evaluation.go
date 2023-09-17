package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ryanma3003/hris/db"
	"github.com/ryanma3003/hris/models"
)

func EvaluationIndex(c *gin.Context) {
	yearNow := time.Now().Year()

	var evaluations []models.Evaluation
	err := db.DB.Where("extract('year' from period::date) = ?", yearNow).Find(&evaluations).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": evaluations,
	})
}
