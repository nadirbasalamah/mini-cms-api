package controllers

import (
	"mini-cms-api/models"
	"mini-cms-api/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	service services.CategoryService
}

func InitCategoryController() CategoryController {
	return CategoryController{
		service: services.InitCategoryService(),
	}
}

func (cc *CategoryController) GetAll(c echo.Context) error {
	categories, err := cc.service.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to retrieve categories",
		})
	}

	return c.JSON(http.StatusOK, models.Response[[]models.Category]{
		Status:  "success",
		Message: "all categories",
		Data:    categories,
	})
}

func (cc *CategoryController) GetByID(c echo.Context) error {
	categoryID := c.Param("id")

	category, err := cc.service.GetByID(categoryID)

	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response[string]{
			Status:  "failed",
			Message: "category not found",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.Category]{
		Status:  "success",
		Message: "category found",
		Data:    category,
	})
}

func (cc *CategoryController) Create(c echo.Context) error {
	var categoryRequest models.CategoryRequest

	if err := c.Bind(&categoryRequest); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	err := categoryRequest.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "please fill all the required fields",
		})
	}

	category, err := cc.service.Create(categoryRequest)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to create a category",
		})
	}

	return c.JSON(http.StatusCreated, models.Response[models.Category]{
		Status:  "success",
		Message: "category created",
		Data:    category,
	})
}

func (cc *CategoryController) Update(c echo.Context) error {
	categoryID := c.Param("id")

	var categoryRequest models.CategoryRequest

	if err := c.Bind(&categoryRequest); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	err := categoryRequest.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "please fill all the required fields",
		})
	}

	category, err := cc.service.Update(categoryRequest, categoryID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to update a category",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.Category]{
		Status:  "success",
		Message: "category updated",
		Data:    category,
	})
}

func (cc *CategoryController) Delete(c echo.Context) error {
	categoryID := c.Param("id")

	err := cc.service.Delete(categoryID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to delete a category",
		})
	}

	return c.JSON(http.StatusOK, models.Response[string]{
		Status:  "success",
		Message: "category deleted",
	})
}
