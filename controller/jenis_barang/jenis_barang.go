package jenis_barang

import (
	"POS-SRI/model/request"
	"POS-SRI/service/jenis_barang"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InputJenisBarang(c echo.Context) error {
	var Request request.Input_Jenis_Barang_Request
	Request.Nama_jenis_barang = c.FormValue("nama_jenis_barang")
	Request.Kode_gudang = c.FormValue("kode_gudang")

	result, err := jenis_barang.Input_Jenis_Barang(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
