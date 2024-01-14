package grade_barang

import (
	"POS-SRI/model/request"
	"POS-SRI/service/grade_barang"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InputGradeBarang(c echo.Context) error {
	var Request request.Input_Grade_Barang_Request
	Request.Nama_grade_barang = c.FormValue("nama_grade_barang")
	Request.Kode_barang = c.FormValue("kode_barang")

	result, err := grade_barang.Input_Grade_Barang(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
