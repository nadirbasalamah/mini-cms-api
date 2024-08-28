package contents

import (
	"mini-cms-api/businesses/contents"
	"mini-cms-api/drivers/mysql/categories"
	"time"

	"gorm.io/gorm"
)

type Content struct {
	ID          uint                `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at"`
	DeletedAt   gorm.DeletedAt      `json:"deleted_at" gorm:"index"`
	Title       string              `json:"title"`
	Description string              `json:"description"`
	CategoryID  uint                `json:"category_id"`
	Category    categories.Category `json:"category"`
}

func (rec *Content) ToDomain() contents.Domain {
	return contents.Domain{
		ID:           rec.ID,
		CreatedAt:    rec.CreatedAt,
		UpdatedAt:    rec.UpdatedAt,
		DeletedAt:    rec.DeletedAt,
		Title:        rec.Title,
		Description:  rec.Description,
		CategoryID:   rec.CategoryID,
		CategoryName: rec.Category.Name,
	}
}

func FromDomain(domain *contents.Domain) *Content {
	return &Content{
		ID:          domain.ID,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
		DeletedAt:   domain.DeletedAt,
		Title:       domain.Title,
		Description: domain.Description,
		CategoryID:  domain.CategoryID,
	}
}
