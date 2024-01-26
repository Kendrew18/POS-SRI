package quality_control

import (
	"POS-SRI/model/request"
	"POS-SRI/service/quality_control"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func ReadQualityControl(c echo.Context) error {
	var Request request.Read_quality_control_Request
	var Request_filter request.Read_Quality_Control_Filter_Request

	Request.Kode_gudang = c.FormValue("kode_gudang")

	Request_filter.Kode_lot = c.FormValue("kode_lot")
	Request_filter.Tanggal_awal = c.FormValue("tanggal_awal")
	Request_filter.Tanggal_akhir = c.FormValue("tanggal_akhir")

	result, err := quality_control.Read_Quality_Control(Request, Request_filter)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func UpdateBeratRillQualityControl(c echo.Context) error {
	var Request request.Update_Berat_Barang_Rill_Request
	var Request_kode request.Update_Quality_Control_Kode_Request

	Request_kode.Kode_barang_quality_control = c.FormValue("kode_barang_quality_control")

	Request.Berat_barang_rill, _ = strconv.ParseFloat(c.FormValue("berat_barang_rill"), 64)
	Request.Berat_barang_ditolak, _ = strconv.ParseFloat(c.FormValue("berat_barang_ditolak"), 64)
	Request.Kadar_air, _ = strconv.ParseFloat(c.FormValue("kadar_air"), 64)

	result, err := quality_control.Update_Berat_Rill_Quality_Control(Request, Request_kode)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func ReadDetailQualityControl(c echo.Context) error {
	var Request request.Read_Detail_Quality_Control_Request
	Request.Kode_quality_control = c.FormValue("kode_quality_control")

	result, err := quality_control.Read_Detail_Quality_Control(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func UpdateStatusQualityControl(c echo.Context) error {
	var Request request.Update_Status_Quality_Control_Request
	var Request_kode request.Kode_Quality_Control_Request

	Request.Status, _ = strconv.Atoi(c.FormValue("status"))

	Request_kode.Kode_quality_control = c.FormValue("kode_quality_control")

	result, err := quality_control.Update_Status_Quality_Control(Request, Request_kode)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
