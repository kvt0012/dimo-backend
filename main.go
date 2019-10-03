package main

import (
	"dimo-backend/config"
	"dimo-backend/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/api/user/register", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/api/user/login", handlers.AuthenticateUser).Methods("POST")
	router.HandleFunc("/api/user/{id}", handlers.GetUserInfo).Methods("GET")

	router.HandleFunc("/api/store/?id={id}",
		handlers.GetStoreById).Methods("GET")
	router.HandleFunc("/api/store/nearest/lat={lat}&long={long}&limit={limit}",
		handlers.GetAllNearestStores).Methods("GET")
	router.HandleFunc("/api/store/nearest/by={key}&value={value}&lat={lat}&long={long}&limit={limit}",
		handlers.GetNearestStoresBy).Methods("GET")
	//	router.HandleFunc("/api/store/nearest/by={key}&value={value}&lat={lat}&long={long}",
	//		handlers.GetRecommendedStores).Methods("GET")

	port := config.ApiPort
	fmt.Println("Listening on port :", port)
	err := http.ListenAndServe(":"+strconv.FormatInt(int64(port), 10), router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
