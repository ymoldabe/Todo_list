package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	todo "github.com/ymoldabe/Todo_list"
	"github.com/ymoldabe/Todo_list/pkg/handler"
	"github.com/ymoldabe/Todo_list/pkg/repository"
	"github.com/ymoldabe/Todo_list/pkg/service"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initalizate config: %s ", err.Error())
	}

	db, err := repository.NewMongoDB(&repository.Config{
		Driver:   viper.GetString("db.driver"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
	})
	if err != nil {
		logrus.Fatalf("Failed to initialize DB: %s", err.Error())
	}
	defer repository.Close(db)

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handler := handler.NewHandler(service)

	srv := new(todo.Server)
	if err = srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		logrus.Fatalf("error of listening server %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

//docker run -d -p 27017:27017 --name mongodb_container mongo

//64d4f04f8e69b16aaf8f172a
//64d4f1468e69b16aaf8f172b
