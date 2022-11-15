package main

import (
	"github.com/aldotp/cakestore/config"
	"github.com/aldotp/cakestore/routes"
	"log"
	"net/http"
)

func main() {
	config.LoadEnv(".env")
	config.ConnectToDatabase()
	r := routes.Router()
	log.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
