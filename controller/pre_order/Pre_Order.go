package pre_order

import (
	"POS-SRI/model/request"
	"POS-SRI/service/pre_order"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func InputPreOrder(c echo.Context) error {
	var Request request.Input_Pre_Order_Request
	var Request_Barang request.Input_Barang_Pre_Order_Request

	Request.Kode_lot = c.FormValue("kode_lot")
	Request.Tanggal = c.FormValue("tanggal")
	Request.Nama_supplier = c.FormValue("nama_supplier")
	Request.Tanggal_ETD = c.FormValue("tanggal_etd")
	Request.Kode_gudang = c.FormValue("kode_gudang")

	Request_Barang.Kode_barang = c.FormValue("kode_barang")
	Request_Barang.Kode_grade_barang = c.FormValue("kode_grade_barang")
	Request_Barang.Berat_barang = c.FormValue("berat_barang")
	Request_Barang.Harga = c.FormValue("harga")

	result, err := pre_order.Input_Pre_Order(Request, Request_Barang)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func ReadPreOrder(c echo.Context) error {
	var Request request.Read_Pre_Order_Request
	var Request_filter request.Read_Pre_Order_Filter_Request

	Request.Kode_gudang = c.FormValue("kode_gudang")

	Request_filter.Tanggal_etd = c.FormValue("tanggal_etd")
	Request_filter.Tanggal_awal = c.FormValue("tanggal_awal")
	Request_filter.Tanggal_akhir = c.FormValue("tanggal_akhir")

	result, err := pre_order.Read_Pre_Order(Request, Request_filter)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func ReadDetailPreOrder(c echo.Context) error {
	var Request request.Read_Detail_Pre_Order_Request

	Request.Kode_pre_order = c.FormValue("kode_pre_order")

	result, err := pre_order.Read_Detail_Pre_Order(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func UpdatePreOrder(c echo.Context) error {
	var Request request.Update_Pre_order_Request
	var Request_kode request.Update_Pre_Order_Kode_Request

	Request.Kode_barang = c.FormValue("kode_barang")
	Request.Kode_grade_barang = c.FormValue("kode_grade_barang")
	Request.Berat_barang, _ = strconv.ParseFloat(c.FormValue("berat_barang"), 64)
	Request.Harga, _ = strconv.ParseInt(c.FormValue("harga"), 10, 64)

	Request_kode.Kode_barang_pre_order = c.FormValue("kode_barang_pre_order")

	result, err := pre_order.Update_Pre_Order(Request, Request_kode)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func DeletePreOrder(c echo.Context) error {
	var Request request.Update_Pre_Order_Kode_Request

	Request.Kode_barang_pre_order = c.FormValue("kode_barang_pre_order")

	result, err := pre_order.Delete_Pre_Order(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func UpdateStatusPreOrder(c echo.Context) error {
	var Request request.Update_Status_Pre_Order_Request
	var Request_kode request.Kode_Pre_Order_Request

	Request.Status, _ = strconv.Atoi(c.FormValue("status"))
	Request_kode.Kode_pre_order = c.FormValue("kode_pre_order")

	result, err := pre_order.Update_Status_Pre_Order(Request, Request_kode)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
