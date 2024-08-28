package contents

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID           uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
	Title        string
	Description  string
	CategoryID   uint
	CategoryName string
}

type UseCase interface {
	GetAll() ([]Domain, error)
	GetByID(id string) (Domain, error)
	Create(contentReq *Domain) (Domain, error)
	Update(contentReq *Domain, id string) (Domain, error)
	Delete(id string) error
}

type Repository interface {
	GetAll() ([]Domain, error)
	GetByID(id string) (Domain, error)
	Create(contentReq *Domain) (Domain, error)
	Update(contentReq *Domain, id string) (Domain, error)
	Delete(id string) error
}
