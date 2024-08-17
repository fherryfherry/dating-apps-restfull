package error

import "github.com/labstack/echo/v4"

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func ErrUnauthorized(c echo.Context) error {
	return c.JSON(401, ErrorResponse{401, "UNAUTHORIZED"})
}

func ErrorResponseBadRequest(c echo.Context, message string) error {
	return c.JSON(400, ErrorResponse{400, message})
}

func ErrorResponseInternalError(c echo.Context, message string) error {
	return c.JSON(500, ErrorResponse{500, message})
}
