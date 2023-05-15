package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/measutosh/aggrss/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at`
	Name      string    `json:"name"`
}

// this will return the json in custom written format(SnakeCase) instead of the format(CamelCase) that comes from db
func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
	}
}
