package user

import (
	"POS-SRI/model/request"
	"POS-SRI/service/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	var Request request.User_Request
	Request.Username = c.FormValue("username")
	Request.Password = c.FormValue("password")

	result, err := user.Login(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
