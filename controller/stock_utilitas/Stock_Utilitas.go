package stock_utilitas

import (
	"POS-SRI/model/request"
	"POS-SRI/service/stock_utilitas"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func InputStockUtilitas(c echo.Context) error {
	var Request request.Input_Stock_Utilitas_Request
	Request.Nama_stock_utilitas = c.FormValue("nama_stock_utilitas")
	Request.Kode_gudang = c.FormValue("kode_gudang")
	Request.Tanggal = c.FormValue("kode_gudang")
	Request.Jumlah, _ = strconv.Atoi(c.FormValue("kode_gudang"))

	result, err := stock_utilitas.Input_Stock_Utilitas(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func ReadStockUtilitas(c echo.Context) error {
	var Request request.Read_Stock_Utilitas_Request
	Request.Kode_gudang = c.FormValue("kode_gudang")

	result, err := stock_utilitas.Read_Stock_Utilitas(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func UpdateStockUtilitas(c echo.Context) error {
	var Request request.Update_Stock_Utilitas_Request
	var Request_kode request.Update_Stock_Utilitas_Kode_Request

	Request.Nama_stock_utilitas = c.FormValue("nama_stock_utilitas")
	Request.Jumlah, _ = strconv.Atoi(c.FormValue("kode_gudang"))
	Request.Tanggal = c.FormValue("tanggal")

	Request_kode.Kode_stock_utilitas = c.FormValue("kode_stock_utilitas")

	result, err := stock_utilitas.Update_Stock_Utilitas(Request, Request_kode)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
