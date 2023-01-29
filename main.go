package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"io/ioutil"
	"net/http"
	"time"
)

var recipes []Recipe

type Recipe struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Tags         []string `json:"tags"`
	Ingredients  []string `json:"ingredients"`
	Instructions []string `json:"instructions"`
	PublishAt    string   `json:"publish_at"`
}

func NewRecipeHandler(c *gin.Context) {
	var recipe Recipe
	if err := c.ShouldBind(&recipe); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	recipe.ID = xid.New().String()
	recipe.PublishAt = time.Now().Format(time.RFC3339)
	recipes = append(recipes, recipe)
	c.JSON(http.StatusOK, recipe)
}

func ListRecipeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, recipes)
}

func init() {
	recipes = make([]Recipe, 0)
	files, _ := ioutil.ReadFile("recipes.json")
	err := json.Unmarshal(files, &recipes)
	if err != nil {
		return
	}
}

func main() {
	r := gin.Default()
	r.POST("/recipes", NewRecipeHandler)
	r.GET("/recipes", ListRecipeHandler)
	err := r.Run()
	if err != nil {
		return
	}
}
