package categories

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Name      string
}

type UseCase interface {
	GetAll() ([]Domain, error)
	GetByID(id string) (Domain, error)
	Create(categoryReq *Domain) (Domain, error)
	Update(categoryReq *Domain, id string) (Domain, error)
	Delete(id string) error
}

type Repository interface {
	GetAll() ([]Domain, error)
	GetByID(id string) (Domain, error)
	Create(categoryReq *Domain) (Domain, error)
	Update(categoryReq *Domain, id string) (Domain, error)
	Delete(id string) error
}
