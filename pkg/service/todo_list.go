package service

import (
	todo "github.com/ymoldabe/Todo_list"
	"github.com/ymoldabe/Todo_list/pkg/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{
		repo: repo,
	}
}

func (s *TodoListService) Create(list todo.Todo) error {
	return s.repo.Create(list)

}
func (s *TodoListService) Update(todoId primitive.ObjectID, form todo.Todo) error {
	return s.repo.Update(todoId, form)
}
func (s *TodoListService) Delete(id primitive.ObjectID) error {
	return s.repo.Delete(id)

}
func (s *TodoListService) MarkTaskAsDone(id primitive.ObjectID) error {
	return s.repo.MarkTaskAsDone(id)
}
func (s *TodoListService) GetTasksByStatus(status string) ([]todo.TodoRes, error) {
	return s.repo.GetTasksByStatus(status)
}

func (s *TodoListService) TaskExistsWithTitleAndActiveAt(input todo.TodoRequset) (bool, error) {
	return s.repo.TaskExistsWithTitleAndActiveAt(input)
}
