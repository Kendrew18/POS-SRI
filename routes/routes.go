package routes

import (
	"POS-SRI/controller/barang"
	"POS-SRI/controller/grade_barang"
	"POS-SRI/controller/pre_order"
	"POS-SRI/controller/quality_control"
	"POS-SRI/controller/sortir"
	"POS-SRI/controller/stock_masuk"
	"POS-SRI/controller/stock_utilitas"
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

	//Barang
	BR := e.Group("/BR")
	BR.POST("/barang", barang.InputBarang)
	BR.GET("/barang", barang.ReadBarang)
	BR.GET("/barang-stock", barang.ReadBarangStock)

	//Grade Barang
	GB := e.Group("/GB")
	GB.POST("/grade-barang", grade_barang.InputGradeBarang)

	//Stock Utilitas
	STU := e.Group("/STU")
	STU.POST("/stock-utilitas", stock_utilitas.InputStockUtilitas)
	STU.GET("/stock-utilitas", stock_utilitas.ReadStockUtilitas)
	STU.PUT("/stock-utilitas", stock_utilitas.UpdateStockUtilitas)

	//Pre Order
	PO := e.Group("/PO")
	PO.POST("/pre-order", pre_order.InputPreOrder)
	PO.GET("/pre-order", pre_order.ReadPreOrder)
	PO.GET("/detail-pre-order", pre_order.ReadDetailPreOrder)
	PO.PUT("/pre-order", pre_order.UpdatePreOrder)
	PO.DELETE("/pre-order", pre_order.DeletePreOrder)
	PO.PUT("/update-status", pre_order.UpdateStatusPreOrder)

	//Quality_Control
	QC := e.Group("/QC")
	QC.GET("/quality-control", quality_control.ReadQualityControl)
	QC.PUT("/quality-control", quality_control.UpdateBeratRillQualityControl)
	QC.GET("/detail", quality_control.ReadDetailQualityControl)
	QC.PUT("/status", quality_control.UpdateStatusQualityControl)

	//Stock_Masuk
	SM := e.Group("/SM")
	SM.GET("/stock-masuk", stock_masuk.ReadStockMasuk)
	SM.GET("/detail", stock_masuk.ReadDetailStockMasuk)

	//Sortir
	SRT := e.Group("/SRT")
	SRT.PUT("/update-status", sortir.UpdateStatusSortir)
	SRT.GET("/sortir", sortir.ReadSortir)
	SRT.GET("/detail", sortir.ReadDetailSortir)
	SRT.POST("/sortir", sortir.UpdateBeratBarangSetelahSortir)

	//Stock Keluar

	return e
}
