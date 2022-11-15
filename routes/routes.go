package routes

import (
	"github.com/aldotp/cakestore/controller"
	"github.com/gorilla/mux"
	// "log"
	// "net/http"
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

// func Router() {
// 	r := mux.NewRouter()
// 	r.HandleFunc("/cakes", controller.GetCakes).Methods("GET")           // 1
// 	r.HandleFunc("/cakes/{id}", controller.GetCakeDetail).Methods("GET") // 2
// 	r.HandleFunc("/cakes", controller.AddNewCake).Methods("POST")        // 3
// 	r.HandleFunc("/cakes/{id}", controller.UpdateCake).Methods("PATCH")  // 4
// 	r.HandleFunc("/cakes/{id}", controller.DeleteCake).Methods("DELETE") // 5

// 	log.Println("Server started on http://localhost:8080")
// 	http.ListenAndServe(":8080", r)
// }
