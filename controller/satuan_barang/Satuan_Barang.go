package satuan_barang

import (
	"POS-SRI/model/request"
	"POS-SRI/service/satuan_barang"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InputSatuanBarang(c echo.Context) error {
	var Request request.Input_Satuan_Barang_Request
	Request.Nama_satuan_barang = c.FormValue("nama_satuan_barang")
	Request.Kode_gudang = c.FormValue("kode_gudang")

	result, err := satuan_barang.Input_Satuan_Barang(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func ReadSatuanBarang(c echo.Context) error {
	var Request request.Read_Satuan_Barang_Request
	Request.Kode_gudang = c.FormValue("kode_gudang")

	result, err := satuan_barang.Read_Satuan_Barang(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
