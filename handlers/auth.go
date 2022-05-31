package handlers

import (
	"context"
	"github.com/auth0-community/go-auth0"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/square/go-jose.v2"
	"net/http"
)

const Auth0Domain = "DOMAIN.eu.auth0.com"
const AUTH0APIIDENTIFIER = "https://api.recipes.io"

type AuthHandler struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewAuthHandler(ctx context.Context, collection *mongo.Collection) *AuthHandler {
	return &AuthHandler{
		collection: collection,
		ctx:        ctx,
	}
}

func (handler *AuthHandler) AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		var auth0Domain = "https://" + Auth0Domain + "/"
		client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: auth0Domain + ".well-known/jwks.json"}, nil)
		configuration := auth0.NewConfiguration(client, []string{AUTH0APIIDENTIFIER}, auth0Domain, jose.RS256)
		validator := auth0.NewValidator(configuration, nil)

		_, err := validator.ValidateRequest(c.Request)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			c.Abort()
			return
		}
		c.Next()
	}
}
