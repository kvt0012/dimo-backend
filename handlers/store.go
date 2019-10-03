package handlers

import (
	"dimo-backend/drivers/postgres"
	"dimo-backend/models"
	"dimo-backend/repos/repoimpl"
	. "dimo-backend/utils"
	"fmt"
	"github.com/gorilla/mux"
	"math"
	"net/http"
	"sort"
	"strconv"
)

func distance(lat1, long1, lat2, long2 float32) float64 {
	dist := math.Pow(float64(lat1-lat2), 2) + math.Pow(float64(long1-long2), 2)
	return dist
}

func sortByDistance(stores []*models.Store, orgLat, orgLong float32) []*models.Store {
	sort.SliceStable(stores, func(i, j int) bool {
		return distance(stores[i].Latitude, stores[i].Longitude, orgLat, orgLong) <
			distance(stores[j].Latitude, stores[j].Longitude, orgLat, orgLong)
	})
	return stores
}

var GetStoreById = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	storeId, err := strconv.Atoi(params["id"])
	if err != nil {
		Respond(w, Message(http.StatusBadRequest, "Invalid Store ID"))
		return
	}
	storeRepo := repoimpl.NewStoreRepo(postgres.ConnectAsDefault().SQL)
	data, err := storeRepo.GetByID(int64(storeId))
	if err != nil {
		Respond(w, Message(http.StatusNotFound, "Store not found"))
		return
	}
	fmt.Println(data)
	var message = Message(http.StatusOK, "Store found")
	message["data"] = data
	Respond(w, message)
}

var GetAllNearestStores = func(w http.ResponseWriter, r *http.Request) {

	// key: category or brand
	params := mux.Vars(r)
	userLat, err := strconv.ParseFloat(params["lat"], 32)
	if err != nil {
		Respond(w, Message(http.StatusBadRequest, "Invalid User's latitude"))
		return
	}
	userLong, err := strconv.ParseFloat(params["long"], 32)
	if err != nil {
		Respond(w, Message(http.StatusBadRequest, "Invalid User's longitude"))
		return
	}
	storeRepo := repoimpl.NewStoreRepo(postgres.ConnectAsDefault().SQL)
	stores, err := storeRepo.GetAll()
	if err != nil {
		Respond(w, Message(http.StatusNotFound, "Brand not found"))
		return
	}
}

var GetNearestStoresBy = func(w http.ResponseWriter, r *http.Request) {

	// key: category or brand
	params := mux.Vars(r)
	userLat, err := strconv.ParseFloat(params["lat"], 32)
	if err != nil {
		Respond(w, Message(http.StatusBadRequest, "Invalid User's latitude"))
		return
	}
	userLong, err := strconv.ParseFloat(params["long"], 32)
	if err != nil {
		Respond(w, Message(http.StatusBadRequest, "Invalid User's longitude"))
		return
	}
	key := params["key"]
	value := params["value"]
	storeRepo := repoimpl.NewStoreRepo(postgres.ConnectAsDefault().SQL)
	stores := make([]*models.Store, 0)
	if key == "brand" {
		stores, err = storeRepo.GetByBrandName(value)
		if err != nil {
			Respond(w, Message(http.StatusNotFound, "Brand not found"))
			return
		}
	} else if key == "category" {
		stores, err = storeRepo.GetByCategory(value)
		if err != nil {
			Respond(w, Message(http.StatusNotFound, "Category not found"))
			return
		}
	} else {
		Respond(w, Message(http.StatusNotFound, "Invalid Request Key"))
		return
	}
	stores = sortByDistance(stores, float32(userLat), float32(userLong))
	var message = Message(http.StatusOK, "Store found")

	fmt.Println(distance(float32(userLat), float32(userLong), stores[0].Latitude, stores[0].Longitude))
	message["data"] = stores[0]
	Respond(w, message)
}
