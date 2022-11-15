package routes

import (
	"github.com/aldotp/cakestore/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/cakes", controller.GetCakes).Methods("GET")           // 1
	r.HandleFunc("/cakes/{id}", controller.GetCakeDetail).Methods("GET") // 2
	r.HandleFunc("/cakes", controller.AddNewCake).Methods("POST")        // 3
	r.HandleFunc("/cakes/{id}", controller.UpdateCake).Methods("PATCH")  // 4
	r.HandleFunc("/cakes/{id}", controller.DeleteCake).Methods("DELETE") // 5
	return r
}
