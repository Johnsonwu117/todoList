package main

import (
	database "todoList/database"
	rou "todoList/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/list")
	rou.AddUserRouter(v1)
	go func() {
		database.DD()
	}()
	router.Run(":1216")
}
