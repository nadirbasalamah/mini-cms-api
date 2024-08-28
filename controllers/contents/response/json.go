package response

import (
	"mini-cms-api/businesses/contents"
	"time"

	"gorm.io/gorm"
)

type Content struct {
	ID           uint           `json:"id" gorm:"primarykey"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Title        string         `json:"title"`
	Description  string         `json:"description"`
	CategoryID   uint           `json:"category_id"`
	CategoryName string         `json:"category_name"`
}

func FromDomain(domain contents.Domain) Content {
	return Content{
		ID:           domain.ID,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
		DeletedAt:    domain.DeletedAt,
		Title:        domain.Title,
		Description:  domain.Description,
		CategoryID:   domain.CategoryID,
		CategoryName: domain.CategoryName,
	}
}
