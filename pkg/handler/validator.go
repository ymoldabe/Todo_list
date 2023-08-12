package handler

import (
	"errors"
	"time"

	todo "github.com/ymoldabe/Todo_list"
)

func checkRequset(input todo.TodoRequset) (todo.Todo, error) {
	if input.Title == "" || input.ActiveAt == "" {
		return todo.Todo{}, errors.New("Title and activeAt fields are required!")
	}
	if len(input.Title) > 200 {
		return todo.Todo{}, errors.New("Title is too long!")
	}
	activeAtTime, err := time.Parse(layout, input.ActiveAt)
	if err != nil {
		return todo.Todo{}, errors.New("Invalid date format!")
	}

	form := todo.Todo{
		Title:    input.Title,
		ActiveAt: activeAtTime.Format(layout),
		Status:   "active",
	}
	return form, nil
}
