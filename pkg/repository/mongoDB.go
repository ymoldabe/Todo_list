package repository

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	Driver   string
	Username string
	Password string
	Host     string
	Port     string
}

func NewMongoDB(cnf *Config) (*mongo.Client, error) {

	uri := fmt.Sprintf("%s://%s:%s", cnf.Driver, cnf.Host, cnf.Port)
	// Установите параметры подключения к MongoDB.
	clientOptions := options.Client().ApplyURI(uri)

	// Подключение к базе данных.
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Проверка, что соединение установлено.
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

//	docker run --name region-todo_mongo_1 -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=mongo -e MONGO_INITDB_ROOT_PASSWORD=123 -d mongo:4.4.23-focal

func Close(db *mongo.Client) {
	if err := db.Disconnect(context.Background()); err != nil {
		log.Fatal("Failed to close the database connection:", err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
