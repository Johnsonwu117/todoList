package routers

import (
	"todoList/model"

	"github.com/gin-gonic/gin"
)

func AddUserRouter(r *gin.RouterGroup) {
	user := r.Group("/todo")
	user.GET("/", model.GetAllList)
	user.GET("/now", model.Get3List)
	user.POST("/", model.CreateList)

	user.PUT("/:id", model.UpdateList)
	user.DELETE("/:id", model.DeleteList)
}
