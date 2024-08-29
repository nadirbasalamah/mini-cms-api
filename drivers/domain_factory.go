package drivers

import (
	contentDomain "mini-cms-api/businesses/contents"
	contentDB "mini-cms-api/drivers/mysql/contents"

	categoryDomain "mini-cms-api/businesses/categories"
	categoryDB "mini-cms-api/drivers/mysql/categories"

	userDomain "mini-cms-api/businesses/users"
	userDB "mini-cms-api/drivers/mysql/users"

	"gorm.io/gorm"
)

func NewContentRepository(conn *gorm.DB) contentDomain.Repository {
	return contentDB.NewMySQLRepository(conn)
}

func NewCategoryRepository(conn *gorm.DB) categoryDomain.Repository {
	return categoryDB.NewMySQLRepository(conn)
}

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMySQLRepository(conn)
}
