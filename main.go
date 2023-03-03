package main

import (
	"fmt"

	"github.com/eduardor2m/api-go-rest/database"
	"github.com/eduardor2m/api-go-rest/models"
	"github.com/eduardor2m/api-go-rest/routes"
)

func main() {
	models.Personalities = []models.Personality{}

	database.Connect()
	fmt.Println("App started")
	routes.HandleRequest()
}
