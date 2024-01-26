package grade_barang

import (
	"POS-SRI/db"
	"POS-SRI/model/request"
	"POS-SRI/model/response"
	"POS-SRI/tools"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

func Input_Grade_Barang(Request request.Input_Grade_Barang_Request) (response.Response, error) {
	var res response.Response
	var err *gorm.DB

	Nama_grade_barang := tools.String_Separator_To_String(Request.Nama_grade_barang)

	var RequestV2 []request.Input_Grade_Barang_Request_V2

	con := db.CreateConGorm().Table("grade_barang")

	co := 0

	err = con.Select("co").Order("co DESC").Limit(1).Scan(&co)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = co
		return res, err.Error
	}

	for i := 0; i < len(Nama_grade_barang); i++ {

		RequestV2[i].Co = co + 1 + i
		RequestV2[i].Kode_grade_barang = "GB-" + strconv.Itoa(RequestV2[i].Co)
		RequestV2[i].Kode_barang = Request.Kode_barang
		RequestV2[i].Nama_grade_barang = Nama_grade_barang[i]
	}

	err = con.Select("co", "kode_barang", "kode_grade_barang", "nama_grade_barang").Create(&RequestV2)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	} else {
		res.Status = http.StatusOK
		res.Message = "Suksess"
		res.Data = map[string]int64{
			"rows": err.RowsAffected,
		}
	}

	return res, nil
}
