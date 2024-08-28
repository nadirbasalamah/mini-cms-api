package controllers

import "github.com/labstack/echo/v4"

type Response[T any] struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

func NewResponse[T any](c echo.Context, statusCode int, statusMsg string, message string, data T) error {
	return c.JSON(statusCode, Response[T]{
		Status:  statusMsg,
		Message: message,
		Data:    data,
	})
}
