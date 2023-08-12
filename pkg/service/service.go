package service

import (
	todo "github.com/ymoldabe/Todo_list"
	"github.com/ymoldabe/Todo_list/pkg/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoList interface {
	Create(list todo.Todo) error
	Update(todoId primitive.ObjectID, form todo.Todo) error
	Delete(id primitive.ObjectID) error
	MarkTaskAsDone(id primitive.ObjectID) error
	GetTasksByStatus(status string) ([]todo.TodoRes, error)
	TaskExistsWithTitleAndActiveAt(input todo.TodoRequset) (bool, error)
}

type Service struct {
	TodoList
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		TodoList: NewTodoListService(repo.TodoList),
	}
}
