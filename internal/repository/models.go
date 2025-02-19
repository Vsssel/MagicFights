// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package repository

import (
	"database/sql"
)

type Card struct {
	ID          int32
	Description string
	Damage      int32
	Power       int32
}

type Character struct {
	ID          int32
	Name        string
	Description string
}

type Room struct {
	Key      string
	IsActive sql.NullBool
}

type User struct {
	ID       int32
	Username string
	RoomKey  sql.NullString
	Hp       int32
}
