package handlers

import (
	"net/http"
	"strconv"
	"todo-list/db"
	"todo-list/models"

	"github.com/gin-gonic/gin"
)

func GetTodoHandler(c *gin.Context) {
	var todos []models.Todo
	db.DB.Find(&todos)
	c.JSON(http.StatusOK, todos)
}

func CreateTodoHandler(c *gin.Context) {
	var newTodo models.Todo
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.DB.Create(&newTodo)
	c.JSON(http.StatusCreated, newTodo)
}

func UpdateTodoHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}

	var todo models.Todo
	if err := db.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	var updateTodo models.Todo
	if err := c.ShouldBindJSON(&updateTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo.Title = updateTodo.Title
	todo.Status = updateTodo.Status
	db.DB.Save(&todo)

	c.JSON(http.StatusOK, todo)
}

func DeleteHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid ID"})
		return
	}

	var todo models.Todo
	if err := db.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	db.DB.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}
