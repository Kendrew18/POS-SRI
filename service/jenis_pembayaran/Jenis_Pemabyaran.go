package jenis_pembayaran

import (
	"POS-SRI/db"
	"POS-SRI/model/request"
	"POS-SRI/model/response"
	"net/http"
	"strconv"
)

func Input_Jenis_Pembayaran(Request request.Input_Jenis_Pembayaran_Request) (response.Response, error) {

	var res response.Response

	con := db.CreateConGorm().Table("jenis_pembayaran")

	co := 0

	err := con.Select("co").Order("co DESC").Limit(1).Scan(&co)

	Request.Co = co + 1
	Request.Kode_jenis_pembayaran = "JP-" + strconv.Itoa(Request.Co)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	err = con.Select("co", "kode_jenis_pembayaran", "nama_jenis_pembayaran", "kode_gudang").Create(&Request)

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

func Read_Jenis_Pembayaran(Request request.Read_Jenis_Pembayaran_Request) (response.Response, error) {

	var res response.Response
	var data []response.Read_Jenis_Pembayaran_Response

	con := db.CreateConGorm().Table("jenis_pembayaran")

	err := con.Select("kode_jenis_pembayaran", "nama_jenis_pembayaran").Where("kode_kasir = ?", Request.Kode_gudang).Order("co ASC").Scan(&data).Error

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = data
		return res, err
	}

	if data == nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = data

	} else {
		res.Status = http.StatusOK
		res.Message = "Suksess"
		res.Data = data
	}

	return res, nil
}
