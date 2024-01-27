package stock_masuk

import (
	"POS-SRI/model/request"
	"POS-SRI/service/stock_masuk"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ReadStockMasuk(c echo.Context) error {
	var Request request.Read_Stock_Masuk_Request
	var Request_filter request.Read_Stock_Masuk_Filter_Request

	Request.Kode_gudang = c.FormValue("kode_gudang")

	Request_filter.Kode_lot = c.FormValue("kode_lot")
	Request_filter.Tanggal_awal = c.FormValue("tanggal_awal")
	Request_filter.Tanggal_akhir = c.FormValue("tanggal_akhir")

	result, err := stock_masuk.Read_Stock_Masuk(Request, Request_filter)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func ReadDetailStockMasuk(c echo.Context) error {
	var Request request.Read_Detail_Stock_Masuk_Request
	Request.Kode_stock_masuk = c.FormValue("kode_stock_masuk")

	result, err := stock_masuk.Read_Detail_Stock_Masuk(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
