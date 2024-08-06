package controllers

import (
	"mini-cms-api/models"
	"mini-cms-api/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ContentController struct {
	service services.ContentService
}

func InitContentController() ContentController {
	return ContentController{
		service: services.InitContentService(),
	}
}

func (cc *ContentController) GetAll(c echo.Context) error {
	contents, err := cc.service.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "error when fetching contents",
		})
	}

	return c.JSON(http.StatusOK, models.Response[[]models.Content]{
		Status:  "success",
		Message: "all contents",
		Data:    contents,
	})
}

func (cc *ContentController) GetByID(c echo.Context) error {
	contentID := c.Param("id")

	content, err := cc.service.GetByID(contentID)

	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response[string]{
			Status:  "failed",
			Message: "content not found",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.Content]{
		Status:  "success",
		Message: "content found",
		Data:    content,
	})
}

func (cc *ContentController) Create(c echo.Context) error {
	var contentReq models.ContentRequest

	if err := c.Bind(&contentReq); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	content, err := cc.service.Create(contentReq)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to create a content",
		})
	}

	return c.JSON(http.StatusCreated, models.Response[models.Content]{
		Status:  "success",
		Message: "content created",
		Data:    content,
	})
}

func (cc *ContentController) Update(c echo.Context) error {
	var contentReq models.ContentRequest

	if err := c.Bind(&contentReq); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	contentID := c.Param("id")

	content, err := cc.service.Update(contentReq, contentID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to update a content",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.Content]{
		Status:  "success",
		Message: "content updated",
		Data:    content,
	})
}

func (cc *ContentController) Delete(c echo.Context) error {
	contentID := c.Param("id")

	err := cc.service.Delete(contentID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to delete content",
		})
	}

	return c.JSON(http.StatusOK, models.Response[string]{
		Status:  "success",
		Message: "content deleted",
	})
}
