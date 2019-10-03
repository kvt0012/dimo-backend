package repoimpl

import (
	"database/sql"
	"dimo-backend/models"
	"dimo-backend/repos"
	"fmt"
	"time"
)

func GetReviewImages(db *sql.DB, review_id int64) ([]sql.NullString, error) {
	queryStatement := `SELECT image_url
						FROM review_images WHERE review_id=$1`
	rows, err := db.Query(queryStatement, review_id)
	if err != nil {
		return nil, err
	}
	urls := make([]sql.NullString, 0)
	for rows.Next() {
		var url string
		err = rows.Scan(&url)
		if err != nil {
			break
		}
		urls = append(urls, sql.NullString{String: url})
	}
	err = rows.Err()
	if err != nil {
		return urls, err
	}
	return urls, nil
}

type ReviewRepoImpl struct {
	db *sql.DB
}

func (r *ReviewRepoImpl) GetAll() ([]*models.Review, error) {
	queryStatement := `SELECT id, user_id, store_id, rating, comment, created_at
						FROM reviews`
	rows, err := r.db.Query(queryStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	reviews := make([]*models.Review, 0)
	for rows.Next() {
		review := models.Review{}
		err = rows.Scan(&review.ID, &review.UserID, &review.StoreID, &review.Rating,
			&review.Comment, &review.CreatedAt)
		if err != nil {
			break
		}
		urls, _ := GetReviewImages(r.db, review.ID)
		review.ImageUrls = urls
		reviews = append(reviews, &review)
	}
	err = rows.Err()
	if err != nil {
		return reviews, err
	}
	return reviews, nil
}

func (r *ReviewRepoImpl) GetByID(id int64) (*models.Review, error) {
	queryStatement := `SELECT id, user_id, store_id, rating, comment, created_at
						FROM reviews WHERE id=$1`
	rows, err := r.db.Query(queryStatement, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	rows.Next()
	review := models.Review{}
	err = rows.Scan(&review.ID, &review.UserID, &review.StoreID, &review.Rating,
					&review.Comment, &review.CreatedAt)
	if err != nil {
		return nil, err
	}
	urls, err := GetReviewImages(r.db, review.ID)
	review.ImageUrls = urls
	if err != nil {
		return &review, err
	}
	return &review, nil
}

func (r *ReviewRepoImpl) DeleteByID(id int64) error {

	rows, err := r.db.Query(`DELETE FROM review_images WHERE review_id=$1`, id)
	if err != nil {
		return err
	}
	defer rows.Close()
	rows, err = r.db.Query(`DELETE FROM reviews WHERE id=$1`, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *ReviewRepoImpl) Insert(review *models.Review) error {
	insertStatement := `
	INSERT INTO reviews (user_id, store_id, rating, comment, created_at)
	VALUES ($1, $2, $3, $4, $5)`
	review.CreatedAt = time.Now()
	res, err := r.db.Exec(insertStatement, review.UserID,
		review.StoreID, review.Rating, review.Comment, review.CreatedAt)
	if err != nil {
		return err
	}
	review.ID, _ = res.LastInsertId()
	fmt.Println("Record added: ", review)
	return nil
}

func (r *ReviewRepoImpl) GetByUserID(userId int64) ([]*models.Review, error) {
	queryStatement := `SELECT id, user_id, store_id, rating, comment, created_at
						FROM reviews WHERE user_id=$1`
	rows, err := r.db.Query(queryStatement, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	reviews := make([]*models.Review, 0)
	for rows.Next() {
		review := models.Review{}
		err = rows.Scan(&review.ID, &review.UserID, &review.StoreID, &review.Rating,
			&review.Comment, &review.CreatedAt)
		if err != nil {
			break
		}
		urls, _ := GetReviewImages(r.db, review.ID)
		review.ImageUrls = urls
		reviews = append(reviews, &review)
	}
	err = rows.Err()
	if err != nil {
		return reviews, err
	}
	return reviews, nil
}

func (r *ReviewRepoImpl) GetByStoreID(storeId int64) ([]*models.Review, error) {
	queryStatement := `SELECT id, user_id, store_id, rating, comment, created_at
						FROM reviews WHERE store_id=$1`
	rows, err := r.db.Query(queryStatement, storeId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	reviews := make([]*models.Review, 0)
	for rows.Next() {
		review := models.Review{}
		err = rows.Scan(&review.ID, &review.UserID, &review.StoreID, &review.Rating,
			&review.Comment, &review.CreatedAt)
		if err != nil {
			break
		}
		urls, _ := GetReviewImages(r.db, review.ID)
		review.ImageUrls = urls
		reviews = append(reviews, &review)
	}
	err = rows.Err()
	if err != nil {
		return reviews, err
	}
	return reviews, nil
}

func NewReviewRepo(db *sql.DB) repos.ReviewRepo {
	return &ReviewRepoImpl{db:db}
}