package categories

import (
	"mini-cms-api/businesses/categories"
	"mini-cms-api/controllers"
	"mini-cms-api/controllers/categories/request"
	"mini-cms-api/controllers/categories/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	categoryUseCase categories.UseCase
}

func NewCategoryController(categoryUC categories.UseCase) *CategoryController {
	return &CategoryController{
		categoryUseCase: categoryUC,
	}
}

func (cc *CategoryController) GetAll(c echo.Context) error {
	categoryRecords, err := cc.categoryUseCase.GetAll()

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to fetch categories", "")
	}

	categories := []response.Category{}

	for _, category := range categoryRecords {
		categories = append(categories, response.FromDomain(category))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "all categories", categories)
}

func (cc *CategoryController) GetByID(c echo.Context) error {
	categoryID := c.Param("id")

	category, err := cc.categoryUseCase.GetByID(categoryID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "category not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "category found", response.FromDomain(category))
}

func (cc *CategoryController) Create(c echo.Context) error {
	categoryReq := request.Category{}

	if err := c.Bind(&categoryReq); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid input", "")
	}

	err := categoryReq.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "please insert all the required fields", "")
	}

	category, err := cc.categoryUseCase.Create(categoryReq.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to create a category", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, "success", "category created", response.FromDomain(category))
}

func (cc *CategoryController) Update(c echo.Context) error {
	categoryReq := request.Category{}

	if err := c.Bind(&categoryReq); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid input", "")
	}

	categoryID := c.Param("id")

	err := categoryReq.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "please insert all the required fields", "")
	}

	category, err := cc.categoryUseCase.Update(categoryReq.ToDomain(), categoryID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to update a category", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "category updated", response.FromDomain(category))
}

func (cc *CategoryController) Delete(c echo.Context) error {
	categoryID := c.Param("id")

	err := cc.categoryUseCase.Delete(categoryID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to delete a category", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "category deleted", "")
}
