package repoimpl

import (
	"database/sql"
	"dimo-backend/models"
	"dimo-backend/repos"
)

type BrandRepoImpl struct {
	db *sql.DB
}

func (b *BrandRepoImpl) GetAll() ([]*models.Brand, error) {
	rows, err := b.db.Query(`SELECT id, name, category, logo_url FROM brands`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	brands := make([]*models.Brand, 0)
	for rows.Next() {
		brand := models.Brand{}
		err := rows.Scan(&brand.ID, &brand.Name, &brand.Category, &brand.LogoUrl)
		if err != nil {
			break
		}
		brands = append(brands, &brand)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return brands, nil
}

func (b *BrandRepoImpl) GetByID(id int64) (*models.Brand, error) {
	queryStatement := `SELECT id, name, category, logo_url 
						FROM brands WHERE id=$1`
	rows, err := b.db.Query(queryStatement, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	rows.Next()
	brand := models.Brand{}
	err = rows.Scan(&brand.ID, &brand.Name, &brand.Category, &brand.LogoUrl)
	if err != nil {
		return &brand, err
	}
	return &brand, nil
}

func (b *BrandRepoImpl) GetByName(name string) (*models.Brand, error) {
	queryStatement := `SELECT id, name, category, logo_url 
						FROM brands WHERE name=$1`
	rows, err := b.db.Query(queryStatement, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	rows.Next()
	brand := models.Brand{}
	err = rows.Scan(&brand.ID, &brand.Name, &brand.Category, &brand.LogoUrl)
	if err != nil {
		return &brand, err
	}
	return &brand, nil
}

func (b *BrandRepoImpl) GetByCategory(category string) ([]*models.Brand, error) {
	rows, err := b.db.Query(`SELECT id, name, category, logo_url FROM brands
									WHERE category=$1`, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	brands := make([]*models.Brand, 0)
	for rows.Next() {
		brand := models.Brand{}
		err := rows.Scan(&brand.ID, &brand.Name, &brand.Category, &brand.LogoUrl)
		if err != nil {
			continue
		}
		brands = append(brands, &brand)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return brands, nil
}

func (b *BrandRepoImpl) GetByTag(tag string) ([]*models.Brand, error) {
	rows, err := b.db.Query(`SELECT DISTINCT brand_id FROM brands
									WHERE tag=$1`, tag)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var brandIds []int64
	for rows.Next() {
		var brandId int64
		err := rows.Scan(&brandId)
		if err != nil {
			break
		}
		brandIds = append(brandIds, brandId)
	}
	brands := make([]*models.Brand, 0)
	for _, brandId := range brandIds {
		incomingBrand, _ := b.GetByID(brandId)
		brands = append(brands, incomingBrand)
	}
	return brands, err
}

func NewBrandRepo(db *sql.DB) repos.BrandRepo {
	return &BrandRepoImpl{db}
}
