package jenis_pembayaran

import (
	"POS-SRI/model/request"
	"POS-SRI/service/jenis_pembayaran"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InputJenisPembayaran(c echo.Context) error {
	var Request request.Input_Jenis_Pembayaran_Request
	Request.Nama_jenis_pembayaran = c.FormValue("nama_jenis_pembayaran")
	Request.Kode_gudang = c.FormValue("kode_gudang")

	result, err := jenis_pembayaran.Input_Jenis_Pembayaran(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func ReadJenisPembayaran(c echo.Context) error {
	var Request request.Read_Jenis_Pembayaran_Request

	Request.Kode_gudang = c.FormValue("kode_gudang")

	result, err := jenis_pembayaran.Read_Jenis_Pembayaran(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
