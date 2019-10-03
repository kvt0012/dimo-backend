package repoimpl

import (
	"database/sql"
	"dimo-backend/models"
	"dimo-backend/repos"
	"fmt"
)

type UserRepoImpl struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) repos.UserRepo {
	return &UserRepoImpl{db}
}

func (u *UserRepoImpl) GetByID(id int64) (*models.User, error) {
	queryStatement := `SELECT id, name, phone, password, image_url, city 
						FROM users WHERE id=$1`
	rows, err := u.db.Query(queryStatement, id)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	rows.Next()
	user := models.User{}
	err = rows.Scan(&user.ID, &user.Name, &user.Phone, &user.Password, &user.ImageUrl, &user.City)
	if err != nil {
		return &user, err
	}
	return &user, nil
}

func (u *UserRepoImpl) GetByPhone(phone string) (*models.User, error) {
	queryStatement := `SELECT id, name, phone, password, image_url, city 
						FROM users WHERE phone=$1`
	rows, err := u.db.Query(queryStatement, phone)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	rows.Next()
	user := models.User{}
	err = rows.Scan(&user.ID, &user.Name, &user.Phone, &user.Password, &user.ImageUrl, &user.City)
	if err != nil {
		return &user, err
	}
	return &user, nil
}

func (u *UserRepoImpl) GetAll() ([]*models.User, error) {
	users := make([]*models.User, 0)
	rows, err := u.db.Query(`SELECT * FROM users 
									WHERE name NOT LIKE '%dummy%'`)
	if err != nil {
		return users, err
	}
	defer rows.Close()
	for rows.Next() {
		user := models.User{}
		err := rows.Scan(
			&user.ID, &user.Name, &user.Phone, &user.Password,
			&user.ImageUrl, user.City,
			&user.CreatedAt)
		if err != nil {
			break
		}
		users = append(users, &user)
	}
	err = rows.Err()
	if err != nil {
		return users, err
	}
	return users, nil
}

func (u *UserRepoImpl) Insert(user *models.User) error {
	insertStatement := `
	INSERT INTO users (name, phone, password)
	VALUES ($1, $2, $3)`
	res, err := u.db.Exec(insertStatement, user.Name, user.Phone, user.Password)
	if err != nil {
		return err
	}
	user.ID, _ = res.LastInsertId()
	fmt.Println("Record added: ", user)
	return nil
}
