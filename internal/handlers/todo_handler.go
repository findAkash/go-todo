package handlers

import (
	"go-todo/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Dependency Injection using a struct
type Handler struct {
	DB *gorm.DB
}

func RegisterRoutes(router *gin.Engine, db * gorm.DB){
	h := &Handler{DB: db} // Inject database
	router.POST("/todos", h.CreateTodoHandler)
}

func (h *Handler) CreateTodoHandler(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.DB.Create(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (h *Handler) ListToDo(c *gin.Context) {
	var todos []models.Todo
	if err := h.DB.Find(&todos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errror": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}
