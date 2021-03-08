package main

import (
	"context"
	"log"
	"todo-go/config"
	"todo-go/route"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	connectionDB, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(
		config.MongoDBEndpoint,
	))
	if err != nil {
		log.Fatal(err)

	}
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	router.Use(cors.New(config))

	route.NoteRoute(connectionDB, router)
	router.Run()
}
