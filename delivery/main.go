package main

import (
	repository "coffee/project/Repository"
	"coffee/project/delivery/controllers"
	"coffee/project/delivery/routers"
	"coffee/project/usecase"
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := ConnectDb()
	if err != nil {
		fmt.Print(err.Error())
		panic("message while connecting!")
	}
	userRepository := repository.NewUserRepository(client, "taskcluster", "clients")
	router := Setup(userRepository)
	router.Run("localhost:8000")

}

// method for connecting to the database
func ConnectDb() (*mongo.Client, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}
	url := os.Getenv("URL")
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	return client, err
}

// method for setting up the all system
func Setup(user *repository.UserRepository) *gin.Engine {
	router := gin.Default()
	userUseCase := usecase.NewRegisterUSeCase(user)
	controller := controllers.NewRegisterController(userUseCase)

	loginUseCase := usecase.NewLoginUseCase(user)
	loginController := controllers.NewController(loginUseCase)
	routers.SetUpRoute(router,
		controller,
		loginController)
	return router

}
