package handlers

import (
	"database/sql"
	"dimo-backend/drivers/postgres"
	"dimo-backend/models"
	"dimo-backend/models/api/review"
	"dimo-backend/repos/repoimpl"
	. "dimo-backend/utils"
	"encoding/json"
	"net/http"
)

var CreateReview = func(w http.ResponseWriter, r *http.Request) {

	reviewData := review.CreateRequest{}
	err := json.NewDecoder(r.Body).Decode(&reviewData)
	if err != nil {
		Respond(w, Message(http.StatusBadRequest, "Invalid Request"))
		return
	}

	review := models.Review{
		UserID:  reviewData.UserID,
		StoreID: reviewData.StoreID,
		Rating:  reviewData.Rating,
		Comment: sql.NullString{
			String: reviewData.Comment,
			Valid:  true,
		},
	}
	reviewRepo := repoimpl.NewReviewRepo(postgres.ConnectAsDefault().SQL)
	err = reviewRepo.Insert(&review)
	if err != nil {
		Respond(w, Message(http.StatusInternalServerError, ""))
		return
	}
	Respond(w, Message(http.StatusOK, "Review added successfully"))
}

var DeleteReview = func(w http.ResponseWriter, r *http.Request) {

	reviewData := review.DeleteRequest{}
	err := json.NewDecoder(r.Body).Decode(&reviewData)
	if err != nil {
		Respond(w, Message(http.StatusBadRequest, "Invalid Request"))
		return
	}

	reviewRepo := repoimpl.NewReviewRepo(postgres.ConnectAsDefault().SQL)
	review, err := reviewRepo.GetByStoreUserID(reviewData.UserID, reviewData.StoreID)
	if err != nil {
		Respond(w, Message(http.StatusNotFound, "Review not found"))
		return
	}
	err = reviewRepo.DeleteByID(review.ID)
	if err != nil {
		Respond(w, Message(http.StatusOK, "Review deleted successfully"))
		return
	}
	Respond(w, Message(http.StatusInternalServerError, "Error"))
}
