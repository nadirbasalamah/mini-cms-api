package drivers

import (
	contentDomain "mini-cms-api/businesses/contents"
	contentDB "mini-cms-api/drivers/mysql/contents"

	categoryDomain "mini-cms-api/businesses/categories"
	categoryDB "mini-cms-api/drivers/mysql/categories"

	"gorm.io/gorm"
)

func NewContentRepository(conn *gorm.DB) contentDomain.Repository {
	return contentDB.NewMySQLRepository(conn)
}

func NewCategoryRepository(conn *gorm.DB) categoryDomain.Repository {
	return categoryDB.NewMySQLRepository(conn)
}
