package routes

import (
	"POS-SRI/controller/barang"
	"POS-SRI/controller/grade_barang"
	"POS-SRI/controller/jenis_pembayaran"
	"POS-SRI/controller/pre_order"
	"POS-SRI/controller/stock_utilitas"
	"POS-SRI/controller/supplier"
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
	BR := e.Group("/BR")
	BR.POST("/barang", barang.InputBarang)
	BR.GET("/barang", barang.ReadBarang)
	BR.GET("/barang-stock", barang.ReadBarangStock)

	//Grade Barang
	GB := e.Group("/GB")
	GB.POST("/grade-barang", grade_barang.InputGradeBarang)

	//Supplier
	SP := e.Group("/SP")
	SP.POST("/supplier", supplier.InputSupplier)
	SP.GET("/supplier", supplier.ReadSupplier)

	//Stock Utilitas
	STU := e.Group("/STU")
	STU.POST("/stock-utilitas", stock_utilitas.InputStockUtilitas)
	STU.GET("//stock-utilitas", stock_utilitas.ReadStockUtilitas)
	STU.PUT("//stock-utilitas", stock_utilitas.UpdateStockUtilitas)

	//Jenis Pembayaran
	JP := e.Group("/JP")
	JP.POST("/jenis-pembayaran", jenis_pembayaran.InputJenisPembayaran)
	JP.GET("/jenis-pembayaran", jenis_pembayaran.ReadJenisPembayaran)

	//Pre Order
	PO := e.Group("/PO")
	PO.POST("/pre-order", pre_order.InputPreOrder)
	PO.GET("/pre-order", pre_order.ReadPreOrder)
	PO.GET("/detail-pre-order", pre_order.ReadDetailPreOrder)
	PO.PUT("/pre-order", pre_order.UpdatePreOrder)
	PO.DELETE("/pre-order", pre_order.DeletePreOrder)

	return e
}
