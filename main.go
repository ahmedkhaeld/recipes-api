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
	"github.com/ahmedkhaeld/recipes-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

var (
	recipesHandler *handlers.RecipesHandler
	authHandler    *handlers.AuthHandler
)

var Cfg, err = utils.LoadConfig(".")

func init() {

	if err != nil {
		log.Fatal("cannot load configuration variables", err)
	}

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(Cfg.MongoURI))
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")
	recipesCol := client.Database(Cfg.DB).Collection(Cfg.RecipesCol)

	redisClient := redis.NewClient(&redis.Options{
		Addr:     Cfg.RedisAddr,
		Password: "",
		DB:       0,
	})

	status := redisClient.Ping(ctx)
	log.Println(status)
	recipesHandler = handlers.NewRecipesHandler(ctx, recipesCol, redisClient)

	usersCol := client.Database(Cfg.DB).Collection(Cfg.UsersCol)
	authHandler = handlers.NewAuthHandler(ctx, usersCol)

}

func main() {
	// for public  api endpoints
	router := gin.Default()

	router.GET("/recipes", recipesHandler.ListRecipesHandler)

	// for private api endpoints
	authorized := router.Group("/")
	authorized.Use(authHandler.AuthMiddleware(Cfg.Auth0Domain, Cfg.Auth0APIIdentifier))
	{
		authorized.POST("/recipes", recipesHandler.NewRecipeHandler)
		authorized.PUT("/recipes/:id", recipesHandler.UpdateRecipeHandler)
		authorized.DELETE("/recipes/:id", recipesHandler.DeleteRecipeHandler)
		authorized.GET("/recipes/:id", recipesHandler.GetOneRecipeHandler)
	}

	router.Run()
}
