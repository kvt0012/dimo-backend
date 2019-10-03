package handlers

import (
	"dimo-backend/drivers/postgres"
	"dimo-backend/models"
	"dimo-backend/models/api/user"
	"dimo-backend/repos/repoimpl"
	. "dimo-backend/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"net/http"
	"strconv"
)

var CreateUser = func(w http.ResponseWriter, r *http.Request) {

	regData := user.RegisterRequest{}
	err := json.NewDecoder(r.Body).Decode(&regData)
	if err != nil {
		Respond(w, Message(http.StatusBadRequest, "Invalid Request"))
		return
	}

	user := models.User{
		Name:     regData.Name,
		Phone:    regData.Phone,
		Password: regData.Password,
	}
	userRepo := repoimpl.NewUserRepo(postgres.ConnectAsDefault().SQL)
	err = userRepo.Insert(&user)
	if err != nil {
		pqErr := err.(*pq.Error)
		if pqErr.Code == "23505" {
			Respond(w, Message(http.StatusConflict, "Duplicated phone number"))
			return
		}
		Respond(w, Message(http.StatusInternalServerError, ""))
		return
	}
	Respond(w, Message(http.StatusOK, "Register successfully"))
}

var AuthenticateUser = func(w http.ResponseWriter, r *http.Request) {

	loginData := user.LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		Respond(w, Message(http.StatusBadRequest, "Invalid Request"))
		return
	}

	userRepo := repoimpl.NewUserRepo(postgres.ConnectAsDefault().SQL)
	user, err := userRepo.GetByPhone(loginData.Phone)
	if err != nil || user.Password != loginData.Password {
		Respond(w, Message(http.StatusUnauthorized, "Unauthorized"))
		return
	}
	Respond(w, Message(http.StatusOK, "Login successfully"))
}

var GetUserInfo = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.Atoi(params["id"])
	if err != nil {
		Respond(w, Message(http.StatusBadRequest, "Invalid User ID"))
		return
	}
	userRepo := repoimpl.NewUserRepo(postgres.ConnectAsDefault().SQL)
	data, err := userRepo.GetByID(int64(userId))
	if err != nil {
		Respond(w, Message(http.StatusNotFound, "User not found"))
		return
	}
	fmt.Println(data)
	var message = Message(http.StatusOK, "User found")
	message["data"] = data
	Respond(w, message)
}
