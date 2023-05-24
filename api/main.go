package main

import (
	"crud_mux/models"
	"crud_mux/routes"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var _ = godotenv.Load(".env")

func main() {
	models.Init()

	r := mux.NewRouter()
	routes.Route(r)

	port := ":" + os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(port, r))
}
