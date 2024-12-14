package main

import (
	"todo-list/db"
	"todo-list/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	dsn := "host=localhost user=postgres password= dbname=to-do-list port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db.MigrateDatabase(dsn, "scripts/init.sql")
	db.ConnectDatabase(dsn)

	router := gin.Default()

	router.GET("/todos", handlers.GetTodoHandler)
	router.POST("/todos", handlers.CreateTodoHandler)
	router.PUT("/todos/:id", handlers.UpdateTodoHandler)
	router.DELETE("/todos/:id", handlers.DeleteHandler)

	router.Run(":8080")
}
