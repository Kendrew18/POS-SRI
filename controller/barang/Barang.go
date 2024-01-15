package barang

import (
	"POS-SRI/model/request"
	"POS-SRI/service/barang"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InputJenisBarang(c echo.Context) error {
	var Request request.Input_Jenis_Barang_Request
	Request.Nama_barang = c.FormValue("nama_barang")
	Request.Kode_gudang = c.FormValue("kode_gudang")

	result, err := barang.Input_Jenis_Barang(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func ReadJenisBarang(c echo.Context) error {
	var Request request.Read_Jenis_Barang_Response
	Request.Kode_gudang = c.FormValue("kode_gudang")

	result, err := barang.Read_Jenis_Barang(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
