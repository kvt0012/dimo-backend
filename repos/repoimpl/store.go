package repoimpl

import (
	"database/sql"
	"dimo-backend/models"
	"dimo-backend/repos"
	"fmt"
)

type StoreRepoImpl struct {
	db *sql.DB
}


func (s *StoreRepoImpl) GetByID(id int64) (*models.Store, error) {
	queryStatement := `SELECT id, brand_id, subname, avg_rating, num_rating,
								address, latitude, longitude, district, city 
						FROM stores WHERE id=$1`
	rows, err := s.db.Query(queryStatement, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	rows.Next()
	store := models.Store{}
	var brand_id int64
	brandRepo := NewBrandRepo(s.db)

	err = rows.Scan(&store.ID, &brand_id, &store.SubName, &store.AvgRating, &store.NumRating,
					&store.Address, &store.Latitude, &store.Longitude, &store.District, &store.City)
	brand, _ := brandRepo.GetByID(brand_id)
	store.BrandName = brand.Name
	store.Category = brand.Category
	store.ImageUrl = brand.ImageUrl
	if err != nil {
		return &store, err
	}
	return &store, nil
}

func (s *StoreRepoImpl) GetAll() ([]*models.Store, error) {
	stores := make([]*models.Store, 0)
	rows, err := s.db.Query(`
		SELECT id, brand_id, subname, avg_rating, num_rating,
		address, latitude, longitude, district, city FROM stores`)
	if err != nil {
		return stores, err
	}
	defer rows.Close()
	brandRepo := NewBrandRepo(s.db)
	for rows.Next() {
		store := models.Store{}
		var brand_id int64
		err = rows.Scan(&store.ID, &brand_id, &store.SubName, &store.AvgRating, &store.NumRating,
			&store.Address, &store.Latitude, &store.Longitude, &store.District, &store.City)
		brand, _ := brandRepo.GetByID(brand_id)
		store.BrandName = brand.Name
		store.Category = brand.Category
		store.ImageUrl = brand.ImageUrl
		if err != nil {
			break
		}
		stores = append(stores, &store)
	}
	err = rows.Err()
	if err != nil {
		return stores, err
	}
	return stores, nil
}

func (s *StoreRepoImpl) GetByBrandName(brandName string) ([]*models.Store, error) {
	brandRepo := NewBrandRepo(s.db)
	brand, err := brandRepo.GetByName(brandName)
	if err != nil {
		return nil, err
	}

	stores := make([]*models.Store, 0)
	rows, err := s.db.Query(`
		SELECT id, subname, avg_rating, num_rating,
		address, latitude, longitude, district, city 
		FROM stores WHERE brand_id=$1`, brand.ID)
	if err != nil {
		return stores, err
	}
	defer rows.Close()
	for rows.Next() {
		store := models.Store{}
		err = rows.Scan(&store.ID, &store.SubName, &store.AvgRating, &store.NumRating,
			&store.Address, &store.Latitude, &store.Longitude, &store.District, &store.City)
		if err != nil {
			break
		}
		store.BrandName = brand.Name
		store.Category = brand.Category
		store.ImageUrl = brand.ImageUrl
		stores = append(stores, &store)
	}
	err = rows.Err()
	if err != nil {
		return stores, err
	}
	return stores, nil
}

func (s *StoreRepoImpl) GetByCity(city string) ([]*models.Store, error) {
	stores := make([]*models.Store, 0)
	rows, err := s.db.Query(`
		SELECT id, brand_id, subname, avg_rating, num_rating,
			address, latitude, longitude, district, city 
		FROM stores WHERE city LIKE $1`, city)
	if err != nil {
		return stores, err
	}
	defer rows.Close()
	brandRepo := NewBrandRepo(s.db)
	for rows.Next() {
		store := models.Store{}
		var brand_id int64
		err = rows.Scan(&store.ID, &brand_id, &store.SubName, &store.AvgRating, &store.NumRating,
			&store.Address, &store.Latitude, &store.Longitude, &store.District, &store.City)
		if err != nil {
			break
		}
		brand, err := brandRepo.GetByID(brand_id)
		if err != nil {
			continue
		}
		store.BrandName = brand.Name
		store.Category = brand.Category
		store.ImageUrl = brand.ImageUrl
		stores = append(stores, &store)
	}
	err = rows.Err()
	if err != nil {
		return stores, err
	}
	return stores, nil
}

func (s *StoreRepoImpl) GetByDistrict(district string) ([]*models.Store, error) {
	stores := make([]*models.Store, 0)
	rows, err := s.db.Query(`
		SELECT id, brand_id, subname, avg_rating, num_rating,
			address, latitude, longitude, district, city 
		FROM stores WHERE district LIKE $1`, district)
	if err != nil {
		return stores, err
	}
	defer rows.Close()
	brandRepo := NewBrandRepo(s.db)
	for rows.Next() {
		store := models.Store{}
		var brand_id int64
		err = rows.Scan(&store.ID, &brand_id, &store.SubName, &store.AvgRating, &store.NumRating,
			&store.Address, &store.Latitude, &store.Longitude, &store.District, &store.City)
		if err != nil {
			break
		}
		brand, err := brandRepo.GetByID(brand_id)
		if err != nil {
			continue
		}
		store.BrandName = brand.Name
		store.Category = brand.Category
		store.ImageUrl = brand.ImageUrl
		stores = append(stores, &store)
	}
	err = rows.Err()
	if err != nil {
		return stores, err
	}
	return stores, nil
}

func (s *StoreRepoImpl) GetByCategory(category string) ([]*models.Store, error) {
	brandRepo := NewBrandRepo(s.db)
	brands, err := brandRepo.GetByCategory(category)
	if err != nil {
		return nil, err
	}
	fmt.Println(brands[0])
	stores := make([]*models.Store, 0)
	for _, brand := range brands {
		rows, err := s.db.Query(`
		SELECT id, subname, avg_rating, num_rating,
		address, latitude, longitude, district, city 
		FROM stores WHERE brand_id=$1`, brand.ID)
		if err != nil {
			return stores, err
		}
		for rows.Next() {
			store := models.Store{}
			err = rows.Scan(&store.ID, &store.SubName, &store.AvgRating, &store.NumRating,
				&store.Address, &store.Latitude, &store.Longitude, &store.District, &store.City)
			if err != nil {
				break
			}
			store.BrandName = brand.Name
			store.Category = brand.Category
			store.ImageUrl = brand.ImageUrl
			stores = append(stores, &store)
		}
		err = rows.Err()
		if err != nil {
			return stores, err
		}
		rows.Close()
	}
	return stores, nil
}

func NewStoreRepo(db *sql.DB) repos.StoreRepo {
	return &StoreRepoImpl{db}
}
