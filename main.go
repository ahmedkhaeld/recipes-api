// Recipes API
//
// This is a sample recipes API.
//Distributed-Applications-in-Gin.
//
// Schemes: http
//Host: localhost:8080
//BasePath: /
//Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta
package main

import (
	"context"
	"github.com/ahmedkhaeld/recipes-api/handlers"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

const (
	mongodbURI = "mongodb://admin:password@localhost:27017/recipes-store?authSource=admin"
	DB         = "recipes-store"
	COL        = "recipes"
)

var (
	recipesHandler *handlers.RecipesHandler
)

func init() {

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodbURI))
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")
	collection := client.Database(DB).Collection(COL)

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	status := redisClient.Ping(ctx)
	log.Println(status)
	recipesHandler = handlers.NewRecipesHandler(ctx, collection, redisClient)

}

func main() {
	// for public  api endpoints
	router := gin.Default()
	router.GET("/recipes", recipesHandler.ListRecipesHandler)

	// for private api endpoints
	authorized := router.Group("/")
	authorized.Use(handlers.AuthMiddleware())
	{
		authorized.POST("/recipes", recipesHandler.NewRecipeHandler)
		authorized.PUT("/recipes/:id", recipesHandler.UpdateRecipeHandler)
		authorized.DELETE("/recipes/:id", recipesHandler.DeleteRecipeHandler)
		authorized.GET("/recipes/:id", recipesHandler.GetOneRecipeHandler)
	}

	router.Run()
}
