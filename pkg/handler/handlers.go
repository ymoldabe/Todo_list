package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ymoldabe/Todo_list/pkg/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(serice *service.Service) *Handler {
	return &Handler{
		service: serice,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		todoList := api.Group("/todo-list")
		{
			task := todoList.Group("task")
			{
				task.POST("/", h.Create)
				task.PUT("/:id", h.Update)
				task.DELETE("/:id", h.Delete)
				task.PUT("/:id/done", h.MarkTaskAsDone)
				task.GET("/", h.GetTasksByStatus)
			}
		}
	}

	return router

}
