package barang

import (
	"POS-SRI/model/request"
	"POS-SRI/service/barang"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func InputBarang(c echo.Context) error {
	var Request request.Input_Jenis_Barang_Request
	Request.Nama_barang = c.FormValue("nama_barang")
	Request.Kode_gudang = c.FormValue("kode_gudang")
	Request.Status, _ = strconv.Atoi(c.FormValue("status"))

	result, err := barang.Input_Barang(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func ReadBarang(c echo.Context) error {
	var Request request.Read_Jenis_Barang_Request
	Request.Kode_gudang = c.FormValue("kode_gudang")

	result, err := barang.Read_Barang(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func ReadBarangStock(c echo.Context) error {
	var Request request.Read_Barang_Stock_Request
	Request.Kode_gudang = c.FormValue("kode_gudang")

	result, err := barang.Read_Barang_Stock(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
