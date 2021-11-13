package main

import (
	"log"
	"net/http"

	"github.com/theAimOne/service-gateway/database"
	"github.com/theAimOne/service-gateway/service"
)

const basePath = "/api"

func main() {
	database.SetupDatabase()
	service.SetupRoutes(basePath)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
