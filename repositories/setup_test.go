package repositories_test

import (
	"mini-cms-api/models"
	"mini-cms-api/repositories/mocks"
	"mini-cms-api/services"
	"os"
	"testing"
)

var (
	userRepository mocks.UserRepository
	userService    services.UserService

	categoryRepository mocks.CategoryRepository
	categoryService    services.CategoryService

	contentRepository mocks.ContentRepository
	contentService    services.ContentService
)

func TestMain(m *testing.M) {
	userService = services.InitUserService(models.JWTOptions{})
	userService.Repository = &userRepository

	categoryService = services.InitCategoryService()
	categoryService.Repository = &categoryRepository

	contentService = services.InitContentService()
	contentService.Repository = &contentRepository

	os.Exit(m.Run())
}
