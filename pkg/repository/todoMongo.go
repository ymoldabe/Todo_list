package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	todo "github.com/ymoldabe/Todo_list"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TodoMongoDB struct {
	mdb *mongo.Collection
}

func NewTodoMongo(mdb *mongo.Database) *TodoMongoDB {
	return &TodoMongoDB{
		mdb: mdb.Collection("task"),
	}
}

func (r *TodoMongoDB) Create(list todo.Todo) error {
	res, err := r.mdb.InsertOne(context.Background(), list)
	fmt.Println(res.InsertedID)
	return err

}

func (r *TodoMongoDB) TaskExistsWithTitleAndActiveAt(input todo.TodoRequset) (bool, error) {
	filter := bson.M{
		"$or": bson.A{
			bson.M{"title": input.Title},
			bson.M{"activeAt": input.ActiveAt},
		},
	}
	count, err := r.mdb.CountDocuments(context.Background(), filter)
	if err != nil {
		fmt.Print(err.Error())
		return true, err
	}
	if count == 0 {
		return false, nil
	}

	fmt.Println(count, "запись существует")

	return true, nil
}

func (r *TodoMongoDB) Update(todoId primitive.ObjectID, form todo.Todo) error {
	filter := bson.M{"_id": todoId}
	update := bson.D{{"$set", bson.M{
		"title":    form.Title,
		"activeAt": form.ActiveAt,
	}}}

	_, err := r.mdb.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	var updatedTodo todo.Todo
	err = r.mdb.FindOne(context.TODO(), filter).Decode(&updatedTodo)
	if err != nil {
		return err
	}
	fmt.Println(updatedTodo)

	return nil
}

func (r *TodoMongoDB) Delete(id primitive.ObjectID) error {
	res, err := r.mdb.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("Error on deleting one Hero")
	}
	return nil

}
func (r *TodoMongoDB) MarkTaskAsDone(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	statusUpdate := bson.D{{"$set", bson.M{"status": "done"}}}

	_, err := r.mdb.UpdateOne(context.TODO(), filter, statusUpdate)
	if err != nil {
		return err
	}
	var updatedTodo todo.Todo
	err = r.mdb.FindOne(context.TODO(), filter).Decode(&updatedTodo)
	if err != nil {
		return err
	}
	fmt.Println(updatedTodo)
	return nil
}
func (r *TodoMongoDB) GetTasksByStatus(status string) ([]todo.TodoRes, error) {
	opts := options.Find()
	opts.SetSort(bson.D{{"activeAt", 1}})

	sortCursor, err := r.mdb.Find(context.TODO(), bson.M{"status": status}, opts)
	if err != nil {
		return []todo.TodoRes{}, err
	}

	var tasks []todo.TodoRes

	for sortCursor.Next(context.TODO()) {
		var task todo.Todo
		if err := sortCursor.Decode(&task); err != nil {
			return nil, err
		}
		ok, day := chekTime(task.ActiveAt)
		if ok {
			tasks = append(tasks, addList(task, day))
		}
	}

	if err := sortCursor.Err(); err != nil {
		return nil, err
	}

	return tasks, nil

}

func addList(task todo.Todo, weekend string) todo.TodoRes {
	list := todo.TodoRes{
		Title:    weekend + task.Title,
		ActiveAt: task.ActiveAt,
	}
	return list

}

func chekTime(activeAT string) (bool, string) {
	active, err := time.Parse("2006-01-02", activeAT)
	if err != nil {
		log.Printf("Error Parse time: %s", err.Error())
		return false, ""
	}
	if active.Before(time.Now()) || active.Equal(time.Now()) {
		if active.Weekday() == time.Sunday || active.Weekday() == time.Saturday {
			return true, "ВЫХОДНОЙ - "
		} else {
			return true, ""
		}
	}
	return false, ""
}
