package repository

import (
	todo "github.com/ymoldabe/Todo_list"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoList interface {
	Create(list todo.Todo) error
	TaskExistsWithTitleAndActiveAt(input todo.TodoRequset) (bool, error)
	GetTasksByStatus(status string) ([]todo.TodoRes, error)
	Update(todoId primitive.ObjectID, form todo.Todo) error
	Delete(id primitive.ObjectID) error
	MarkTaskAsDone(id primitive.ObjectID) error
}

type Repository struct {
	TodoList
}

func NewRepository(mdb *mongo.Client) *Repository {
	return &Repository{
		TodoList: NewTodoMongo(mdb.Database("todo_db")),
	}
}
