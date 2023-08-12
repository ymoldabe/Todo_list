package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	todo "github.com/ymoldabe/Todo_list"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const layout = "2006-01-02"

func (h *Handler) Create(c *gin.Context) {
	var input todo.TodoRequset

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	form, err := checkRequset(input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	ok, err := h.service.TaskExistsWithTitleAndActiveAt(input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "The task exists!")
		return
	}
	if ok {
		fmt.Println("Запись есть")
		newErrorResponse(c, http.StatusBadRequest, "The task exists!")
		return
	}

	if err := h.service.TodoList.Create(form); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"code": http.StatusNoContent,
	})

}

func (h *Handler) Update(c *gin.Context) {
	var uri todo.TodoURI
	if err := c.BindUri(&uri); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	todoId, err := primitive.ObjectIDFromHex(uri.Id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var input todo.TodoRequset

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	form, err := checkRequset(input)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.service.Update(todoId, form); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": http.StatusNoContent,
	})
}
func (h *Handler) Delete(c *gin.Context) {
	var uri todo.TodoURI
	if err := c.BindUri(&uri); err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"code": http.StatusNotFound,
		})
		return
	}
	todoId, err := primitive.ObjectIDFromHex(uri.Id)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"code": http.StatusNotFound,
		})
		return
	}
	if err := h.service.Delete(todoId); err != nil {
		log.Printf(err.Error())
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"code": http.StatusNotFound,
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": http.StatusNoContent,
	})

}
func (h *Handler) MarkTaskAsDone(c *gin.Context) {
	var uri todo.TodoURI
	if err := c.BindUri(&uri); err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"code": http.StatusNotFound,
		})
		return
	}
	todoId, err := primitive.ObjectIDFromHex(uri.Id)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"code": http.StatusNotFound,
		})
		return
	}
	if err := h.service.MarkTaskAsDone(todoId); err != nil {
		log.Printf(err.Error())
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"code": http.StatusNotFound,
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": http.StatusNoContent,
	})

}
func (h *Handler) GetTasksByStatus(c *gin.Context) {
	status := c.DefaultQuery("status", "active")
	todo_lists, err := h.service.GetTasksByStatus(status)
	if err != nil {
		log.Printf(err.Error())
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"code": http.StatusNotFound,
		})
	}
	c.JSON(http.StatusOK, todo_lists)
}
