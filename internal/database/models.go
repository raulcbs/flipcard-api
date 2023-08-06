// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Name      string
	Email     sql.NullString
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
