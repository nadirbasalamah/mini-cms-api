package contents

import (
	"mini-cms-api/businesses/contents"
	"mini-cms-api/controllers"
	"mini-cms-api/controllers/contents/request"
	"mini-cms-api/controllers/contents/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ContentController struct {
	contentUseCase contents.UseCase
}

func NewContentController(contentUC contents.UseCase) *ContentController {
	return &ContentController{
		contentUseCase: contentUC,
	}
}

func (cc *ContentController) GetAll(c echo.Context) error {
	contentRecords, err := cc.contentUseCase.GetAll()

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to fetch contents", "")
	}

	contents := []response.Content{}

	for _, content := range contentRecords {
		contents = append(contents, response.FromDomain(content))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "all contents", contents)
}

func (cc *ContentController) GetByID(c echo.Context) error {
	contentID := c.Param("id")

	content, err := cc.contentUseCase.GetByID(contentID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "content not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "content found", response.FromDomain(content))
}

func (cc *ContentController) Create(c echo.Context) error {
	contentReq := request.Content{}

	if err := c.Bind(&contentReq); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid input", "")
	}

	err := contentReq.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "please insert all the required fields", "")
	}

	content, err := cc.contentUseCase.Create(contentReq.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to create a content", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, "success", "content created", response.FromDomain(content))
}

func (cc *ContentController) Update(c echo.Context) error {
	contentReq := request.Content{}

	if err := c.Bind(&contentReq); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid input", "")
	}

	contentID := c.Param("id")

	err := contentReq.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "please insert all the required fields", "")
	}

	content, err := cc.contentUseCase.Update(contentReq.ToDomain(), contentID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to update a content", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "content updated", response.FromDomain(content))
}

func (cc *ContentController) Delete(c echo.Context) error {
	contentID := c.Param("id")

	err := cc.contentUseCase.Delete(contentID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to delete a content", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "content deleted", "")
}
