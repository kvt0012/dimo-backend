package controllers

import (
	"dimo-backend/driver/db"
	"dimo-backend/models"
	"dimo-backend/models/api"
	"dimo-backend/repos/repoimpl"
	"dimo-backend/utils"
	"encoding/json"
	"net/http"
)

var RegisterUser = func(w http.ResponseWriter, r *http.Request) {

	regData := api.Registration{}
	err := json.NewDecoder(r.Body).Decode(&regData)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	user := models.User{
		Name:     regData.Name,
		Phone:    regData.Phone,
		Password: regData.Password,
	}
	userRepo := repoimpl.NewUserRepo(db.ConnectDefault().SQL)
	err = userRepo.Insert(&user)
}
