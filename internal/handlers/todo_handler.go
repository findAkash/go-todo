package handlers

import (
	"go-todo/internal/database"
	"go-todo/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTodoHandler(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.Create(&todo)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}

	c.JSON(http.StatusOK, todo)
}