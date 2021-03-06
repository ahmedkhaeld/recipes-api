package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	router := gin.Default()
	router.Run()
}

type Recipe struct {
	Name         string    `json:"name"`
	Tags         []string  `json:"tags"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	PublishedAt  time.Time `json:"publishedAt"`
}
