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
	router.HandleFunc("/", handlers.Default).Methods("GET")

	router.HandleFunc("/api/store/id={id}&lat={lat}&long={long}",
		handlers.GetStoreById).Methods("GET")
	router.HandleFunc("/api/store/user_id={user_id}&lat={lat}&long={long}&km_limit={km_limit}",
		handlers.GetStoresByDistLimit).Methods("GET")

	router.HandleFunc("/api/review/create", handlers.CreateReview).Methods("POST")
	router.HandleFunc("/api/review/delete", handlers.DeleteReview).Methods("DELETE")

	port := config.ApiPort
	fmt.Println("Listening on port:", port)
	err := http.ListenAndServe(":"+strconv.FormatInt(int64(port), 10), router)
	if err != nil {
		fmt.Print(err)
	}
}
