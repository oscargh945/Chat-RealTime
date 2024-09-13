package entities

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID       uuid.UUID `json:"id" db:"id"`
	UserName string    `json:"user_name" db:"user_name"`
	Email    string    `json:"email" db:"email"`
	Password string    `json:"password" db:"password"`
	CreateAt time.Time `json:"created_at" db:"created_at"`
}

type CreateUserReq struct {
	UserName string `json:"user_name" db:"user_name"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type CreateUserResp struct {
	ID       uuid.UUID `json:"id" db:"id"`
	UserName string    `json:"user_name" db:"user_name"`
	Email    string    `json:"email" db:"email"`
}
