package handlers

import (
	"dimo-backend/drivers/postgres"
	"dimo-backend/drivers/recsys"
	"dimo-backend/models"
	"dimo-backend/models/api/store_api"
	"dimo-backend/repos/repoimpl"
	. "dimo-backend/utils"
	"fmt"
	"github.com/gorilla/mux"
	"math"
	"net/http"
	"sort"
	"strconv"
	"time"
)

func distance(lat1, lng1, lat2, lng2 float32) float64 {
	const PI float64 = 3.141592653589793

	radlat1 := PI * float64(lat1) / 180
	radlat2 := PI * float64(lat2) / 180

	theta := float64(lng1 - lng2)
	radtheta := float64(PI * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)

	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / PI
	dist = dist * 60 * 1.1515
	dist = dist * 1.609344
	return dist
}

func sortForUser(stores []*store_api.ResponseData) []*store_api.ResponseData {
	sort.SliceStable(stores, func(i, j int) bool {
		return stores[i].Distance < stores[j].Distance
	})

	brandDict := map[string]bool{}
	firstPart := make([]*store_api.ResponseData, 0)
	secondPart := make([]*store_api.ResponseData, 0)
	for _, store := range stores {
		if _, ok := brandDict[store.BrandName]; !ok {
			firstPart = append(firstPart, store)
			brandDict[store.BrandName] = true
		} else {
			secondPart = append(secondPart, store)
		}
	}

	sort.SliceStable(firstPart, func(i, j int) bool {
		return firstPart[i].RecommendRank < firstPart[j].RecommendRank
	})
	for _, store := range secondPart {
		firstPart = append(firstPart, store)
	}
	return firstPart
}

var GetStoreById = func(w http.ResponseWriter, r *http.Request) {
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
	storeId, err := strconv.Atoi(params["id"])
	if err != nil {
		Respond(w, Message(http.StatusBadRequest, "Invalid Store ID"))
		return
	}
	storeRepo := repoimpl.NewStoreRepo(postgres.ConnectAsDefault().SQL)
	store, err := storeRepo.GetByID(int64(storeId))
	if err != nil {
		Respond(w, Message(http.StatusNotFound, "Store not found"))
		panic(err)
		return
	}
	data := store_api.ResponseData{
		ID:            store.ID,
		BrandName:     store.BrandName,
		SubName:       store.SubName,
		Category:      store.Category,
		LogoUrl:       "https://d1nhio0ox7pgb.cloudfront.net/_img/g_collection_png/standard/512x512/store.png",
		Address:       store.Address,
		Latitude:      store.Latitude,
		Longitude:     store.Longitude,
		Distance:      distance(store.Latitude, store.Longitude, float32(userLat), float32(userLong)),
		RecommendRank: 0,
		AvgRating:     store.AvgRating,
		NumRating:     store.NumRating,
	}
	var message = Message(http.StatusOK, "Store found")
	message["data"] = data
	Respond(w, message)
}

var GetStoresByDistLimit = func(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	params := mux.Vars(r)
	kmLimit, err := strconv.ParseFloat(params["km_limit"], 32)
	if err != nil {
		Respond(w, Message(http.StatusBadRequest, "Invalid km limit"))
		return
	}
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
		Respond(w, Message(http.StatusNotFound, "Store not found"))
		return
	}
	brandRepo := repoimpl.NewBrandRepo(postgres.ConnectAsDefault().SQL)
	brandNameToID := map[string]int64{}
	brandIds := make([]int64, 0)

	finalStores := make([]*store_api.ResponseData, 0)
	for _, store := range stores {
		dist := distance(store.Latitude, store.Longitude, float32(userLat), float32(userLong))
		if dist <= kmLimit {
			cvtStore := store_api.ResponseData{
				ID:            store.ID,
				BrandName:     store.BrandName,
				SubName:       store.SubName,
				Category:      store.Category,
				LogoUrl:       "https://d1nhio0ox7pgb.cloudfront.net/_img/g_collection_png/standard/512x512/store.png",
				Address:       store.Address,
				Latitude:      store.Latitude,
				Longitude:     store.Longitude,
				Distance:      dist,
				RecommendRank: 0,
				AvgRating:     store.AvgRating,
				NumRating:     store.NumRating,
			}
			currBrand, err := brandRepo.GetByName(store.BrandName)
			if err == nil {
				brandId := currBrand.ID
				if _, ok := brandNameToID[store.BrandName]; !ok {
					brandNameToID[store.BrandName] = brandId
					brandIds = append(brandIds, brandId)
				}
			}
			finalStores = append(finalStores, &cvtStore)
		}
	}
	fmt.Println(brandIds)

	brandRank := map[int64]int{}
	userId, err := strconv.Atoi(params["user_id"])
	result, err := recsys.FactorizationRequest(int64(userId), brandIds)
	if len(result) < 1 || err != nil {
		sort.SliceStable(brandIds, func(i, j int) bool {
			ci, _ := storeRepo.CountByBrand(brandIds[i])
			cj, _ := storeRepo.CountByBrand(brandIds[j])
			return ci > cj
		})
	} else {
		brandIds = result
	}
	for idx, brandId := range brandIds {
		brandRank[brandId] = idx + 1
	}
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(brandIds)
	}

	for _, store := range finalStores {
		store.RecommendRank = brandRank[brandNameToID[store.BrandName]]
	}

	finalStores = sortForUser(finalStores)
	var message = Message(http.StatusOK, "Stores found")
	message["process_time"] = float64(time.Now().UnixNano()-startTime.UnixNano()) / 1000000000
	message["data"] = finalStores
	Respond(w, message)
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
	limit, _ := strconv.Atoi(params["limit"])
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
	stores = stores[:limit]
	var message = Message(http.StatusOK, "Store found")

	finalStores := make([]*store_api.ResponseData, 0)
	for _, store := range stores {
		cvtStore := store_api.ResponseData{
			ID:            store.ID,
			BrandName:     store.BrandName,
			SubName:       store.SubName,
			LogoUrl:       store.LogoUrl.String,
			Address:       store.Address,
			Distance:      distance(store.Latitude, store.Longitude, float32(userLat), float32(userLong)),
			RecommendRank: 0,
			AvgRating:     store.AvgRating,
			NumRating:     store.NumRating,
		}
		finalStores = append(finalStores, &cvtStore)
	}

	message["data"] = finalStores
	Respond(w, message)
}
