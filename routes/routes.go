package routes

import (
	"POS-SRI/controller/barang"
	"POS-SRI/controller/grade_barang"
	"POS-SRI/controller/user"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Project-SRI")
	})

	//User
	US := e.Group("/US")
	US.GET("/login", user.Login)

	//Jenis Barang
	JB := e.Group("/BR")
	JB.POST("/barang", barang.InputJenisBarang)

	//Grade Barang
	GB := e.Group("/GB")
	GB.POST("/grade-barang", grade_barang.InputGradeBarang)

	return e
}
