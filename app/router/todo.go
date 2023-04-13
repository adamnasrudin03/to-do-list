package router

import (
	"adamnasrudin03/to-do-list/app/controller"

	"github.com/gin-gonic/gin"
)

func TodoRouter(e *gin.Engine, h controller.TodoController) {
	todoRoutes := e.Group("/todo-items")
	{

		todoRoutes.POST("/", h.Create)
		todoRoutes.GET("/", h.GetAll)
		todoRoutes.PATCH("/:id", h.Update)
		todoRoutes.GET("/:id", h.GetOne)
		todoRoutes.DELETE("/:id", h.Delete)
	}
}
