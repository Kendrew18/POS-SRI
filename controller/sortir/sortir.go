package sortir

import (
	"POS-SRI/model/request"
	"POS-SRI/service/sortir"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func UpdateStatusSortir(c echo.Context) error {
	var Request request.Update_Status_Sortir_Request
	var Request_kode request.Kode_Stock_Masuk_Request
	Request.Status_sortir, _ = strconv.Atoi(c.FormValue("status_sortir"))

	Request_kode.Kode_stock_masuk = c.FormValue("kode_stock_masuk")

	result, err := sortir.Update_Status_Sortir(Request, Request_kode)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func ReadSortir(c echo.Context) error {
	var Request request.Read_Sortir_Request
	var Request_filter request.Read_Sortir_Filter_Request
	Request.Kode_gudang = c.FormValue("kode_gudang")

	Request_filter.Kode_lot = c.FormValue("kode_lot")
	Request_filter.Tanggal_awal = c.FormValue("tanggal_awal")
	Request_filter.Tanggal_akhir = c.FormValue("tanggal_akhir")

	result, err := sortir.Read_Sortir(Request, Request_filter)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func ReadDetailSortir(c echo.Context) error {
	var Request request.Read_Detail_Sortir_Request

	Request.Kode_sortir = c.FormValue("kode_sortir")

	result, err := sortir.Read_Detail_Sortir(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func UpdateBeratBarangSetelahSortir(c echo.Context) error {
	var Request request.Input_Barang_Sortir_Request

	Request.Kode_barang_sortir = c.FormValue("kode_barang_sortir")
	Request.Kode_sortir = c.FormValue("kode_sortir")
	Request.Kode_grade_barang = c.FormValue("kode_grade_barang")
	Request.Berat_setelah_sortir, _ = strconv.ParseFloat(c.FormValue("berat_setelah_sortir"), 64)
	Request.Kadar_air, _ = strconv.ParseFloat(c.FormValue("kadar_air"), 64)

	result, err := sortir.Update_Berat_Barang_Setelah_Sortir(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
