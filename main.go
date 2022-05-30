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
	ctx            context.Context
	err            error
	client         *mongo.Client
	collection     *mongo.Collection
)

func init() {

	ctx = context.Background()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongodbURI))
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")
	collection = client.Database(DB).Collection(COL)
	recipesHandler = handlers.NewRecipesHandler(ctx, collection)

}

func main() {
	router := gin.Default()
	router.POST("/recipes", recipesHandler.NewRecipeHandler)
	router.GET("/recipes", recipesHandler.ListRecipesHandler)
	router.PUT("/recipes/:id", recipesHandler.UpdateRecipeHandler)
	router.DELETE("/recipes/:id", recipesHandler.DeleteRecipeHandler)
	router.GET("/recipes/:id", recipesHandler.GetRecipeHandler)
	//router.GET("/recipes/search", SearchRecipesHandler)
	router.Run()
}
