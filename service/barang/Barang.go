package barang

import (
	"POS-SRI/db"
	"POS-SRI/model/request"
	"POS-SRI/model/response"
	"fmt"
	"net/http"
	"strconv"
)

func Input_Barang(Request request.Input_Jenis_Barang_Request) (response.Response, error) {
	var res response.Response

	con := db.CreateConGorm().Table("barang")

	co := 0

	err := con.Select("co").Order("co DESC").Limit(1).Scan(&co)

	Request.Co = co + 1
	Request.Kode_barang = "BR-" + strconv.Itoa(Request.Co)

	fmt.Println(co)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = co
		return res, err.Error
	}

	err = con.Select("co", "kode_barang", "nama_barang", "kode_gudang").Create(&Request)

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

func Read_Barang(Request request.Read_Jenis_Barang_Request) (response.Response, error) {
	var res response.Response

	var jenis_barang []response.Read_Jenis_Barang_Response
	var data response.Read_Jenis_Barang_Response

	con := db.CreateConGorm()

	rows, err := con.Table("barang").Select("kode_barang", "nama_barang").Where("kode_gudang=?", Request.Kode_gudang).Rows()

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = jenis_barang
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		var grade_barang []response.Read_Grade_Barang_Response

		err = rows.Scan(&data.Kode_barang, &data.Nama_barang)

		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = jenis_barang
			return res, err
		}

		err = con.Table("grade_barang").Select("kode_grade_barang", "nama_grade_barang").Where("kode_barang = ?", data.Kode_barang).Scan(&grade_barang).Error

		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = jenis_barang
			return res, err
		}

		data.Grade_barang = grade_barang

		jenis_barang = append(jenis_barang, data)
	}

	if jenis_barang == nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = jenis_barang

	} else {
		res.Status = http.StatusOK
		res.Message = "Suksess"
		res.Data = jenis_barang
	}

	return res, nil
}

func Read_Barang_Stock(Request request.Read_Barang_Stock_Request) (response.Response, error) {
	var res response.Response

	var arr_barang []response.Read_Barang_Stock_Response

	con := db.CreateConGorm()

	err := con.Table("barang").Select("kode_barang", "nama_barang", "total_berat", "nama_satuan_barang").Joins("JOIN satuan_barang sb on sb.kode_satuan_barang = barang.kode_satuan_barang").Where("kode_gudang=?", Request.Kode_gudang)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	if arr_barang == nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = arr_barang

	} else {
		res.Status = http.StatusOK
		res.Message = "Suksess"
		res.Data = arr_barang
	}

	return res, nil
}
