package supplier

import (
	"POS-SRI/model/request"
	"POS-SRI/service/supplier"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InputSupplier(c echo.Context) error {
	var Request request.Input_Supplier_Request
	var Request_Barang request.Input_Barang_Supplier_Request
	Request.Nama_supplier = c.FormValue("nama_supplier")
	Request.Nomor_telpon = c.FormValue("nomor_telpon")
	Request.Kode_gudang = c.FormValue("kode_gudang")

	Request_Barang.Kode_barang = c.FormValue("kode_barang")

	result, err := supplier.Input_Supplier(Request, Request_Barang)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func ReadSupplier(c echo.Context) error {
	var Request request.Read_Supplier_Request
	Request.Kode_gudang = c.FormValue("kode_gudang")

	result, err := supplier.Read_Supplier(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
