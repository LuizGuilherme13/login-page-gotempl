package repository

import (
	"database/sql"
	"fmt"

	"github.com/LuizGuilherme13/login-page-gotempl/internal/initializers"
	"github.com/LuizGuilherme13/login-page-gotempl/internal/models"
	_ "github.com/lib/pq"
)

func CreateUser(tx *sql.Tx, newUser models.User) error {
	query := "INSERT INTO users(username, password, email)"
	query += "VALUES($1, $2, $3)"

	_, err := tx.Exec(query, newUser.UserName, newUser.Password, newUser.Email)
	if err != nil {
		return fmt.Errorf("repository.CreateUser: %w", err)
	}

	return nil
}

func AuthWithUsernameOrEmail(username, password string) (*models.User, error) {
	db, err := sql.Open("postgres", initializers.DB.String())
	if err != nil {
		return nil, fmt.Errorf("repository.AuthWithUsernameOrEmail(): %w", err)
	}
	defer db.Close()

	var user models.User

	query := "SELECT"
	query += " user_id, username, password,"
	query += " email, created_on, last_login "
	query += "FROM users"
	query += " WHERE (username = $1 OR email= $1) AND password = $2"

	if err = db.QueryRow(query, username, password).Scan(
		&user.UserID,
		&user.UserName,
		&user.Password,
		&user.Email,
		&user.CreatedOn,
		&user.LastLogin,
	); err != nil {
		return nil, fmt.Errorf("repository.AuthWithUsernameOrEmail(): %w", err)
	}

	return &user, nil

}
