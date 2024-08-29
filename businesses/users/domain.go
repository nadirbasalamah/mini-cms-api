package users

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Email     string
	Password  string
}

type UseCase interface {
	Register(userReq *Domain) (Domain, error)
	Login(userReq *Domain) (string, error)
	GetUserInfo(id string) (Domain, error)
}

type Repository interface {
	Register(userReq *Domain) (Domain, error)
	GetUserByEmail(userReq *Domain) (Domain, error)
	GetUserInfo(id string) (Domain, error)
}
