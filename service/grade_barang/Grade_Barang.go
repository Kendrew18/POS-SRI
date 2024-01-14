package grade_barang

import (
	"POS-SRI/db"
	"POS-SRI/model/request"
	"POS-SRI/model/response"
	"POS-SRI/tools"
	"fmt"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

func Input_Grade_Barang(Request request.Input_Grade_Barang_Request) (response.Response, error) {
	var res response.Response
	var err *gorm.DB

	Nama_grade_barang := tools.String_Separator_To_String(Request.Nama_grade_barang)

	for i := 0; i < len(Nama_grade_barang); i++ {
		var RequestV2 request.Input_Grade_Barang_Request_V2

		con := db.CreateConGorm().Table("grade_barang")

		co := 0

		err = con.Select("co").Order("co DESC").Limit(1).Scan(&co)

		RequestV2.Co = co + 1
		RequestV2.Kode_grade_barang = "GB-" + strconv.Itoa(RequestV2.Co)
		RequestV2.Kode_barang = Request.Kode_barang
		RequestV2.Nama_grade_barang = Nama_grade_barang[i]

		fmt.Println(co)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = co
			return res, err.Error
		}

		err = con.Select("co", "kode_barang", "kode_grade_barang", "nama_jenis_barang").Create(&RequestV2)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}
	}

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
