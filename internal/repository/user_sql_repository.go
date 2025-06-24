package repository

import (
	"database/sql"
	"fmt"
	"log"
	"main/internal/models"
)

type SQLUserRepository struct {
	db *sql.DB
}

func (ur *SQLUserRepository) FindAll(search string, page, limit int) ([]models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (ur *SQLUserRepository) FindByUUID(uuid string, user *models.User) error {
	row := ur.db.QueryRow("select * from users where user_id = $1", uuid)

	err := row.Scan(&user.UUID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("user not found")
		}
		return fmt.Errorf("failed to create user: %s", err)
	}
	return nil
}

func (ur *SQLUserRepository) FindByEmail(email string) (models.User, bool) {
	return models.User{}, false
}

func (ur *SQLUserRepository) Create(user models.User) error {
	log.Print("Create user sql repository..... ")
	//log.Printf("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING user_id)", user.Name, user.Email)
	row := ur.db.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING user_id", user.Name, user.Email)

	err := row.Scan(&user.Id)
	if err != nil {
		return fmt.Errorf("failed to create user: %s", err)
	}
	return nil
}

func (ur *SQLUserRepository) Update(uuid string, user models.User) error {
	//TODO implement me
	panic("implement me")
}

func (ur *SQLUserRepository) Delete(uuid string) error {
	//TODO implement me
	panic("implement me")
}

func NewSQLUserRepository(DB *sql.DB) UserRepository {
	return &SQLUserRepository{
		db: DB,
	}
}
