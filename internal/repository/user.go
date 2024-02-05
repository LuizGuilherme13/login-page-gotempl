package repository

import (
	"database/sql"
	"fmt"

	"github.com/LuizGuilherme13/login-page-gotempl/internal/initializers"
	"github.com/LuizGuilherme13/login-page-gotempl/internal/models"
	"github.com/LuizGuilherme13/norm/pkg/db"
	"github.com/LuizGuilherme13/norm/pkg/norm"
	"github.com/LuizGuilherme13/norm/pkg/norm/repository/postgres"

	_ "github.com/lib/pq"
)

func CreateUser(tx *sql.Tx, newUser models.User) error {

	nORM := norm.NewService(postgres.New(initializers.DB))

	err := nORM.InTable("users").FromModel(&newUser).Omit("user_id").Create()
	if err != nil {
		return fmt.Errorf("repository.CreateUser: %w", err)
	}

	return nil
}

func AuthWithUsernameOrEmail(username, password string) (*models.User, error) {
	user := models.User{}

	nORM := norm.NewService(postgres.New(initializers.DB))

	err := nORM.InTable("users").ToModel(&user).Only("*").WithConditions(
		norm.And(db.Condition{
			Condition: "username = $1 OR email = $1", Value: username}),
		norm.And(db.Condition{
			Condition: "password = $2", Value: password}),
	).Find()
	if err != nil {
		return nil, fmt.Errorf("repository.AuthWithUsernameOrEmail(): %w", err)
	}

	return &user, nil

}
