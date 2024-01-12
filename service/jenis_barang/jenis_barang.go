package jenis_barang

import (
	"POS-SRI/db"
	"POS-SRI/model/request"
	"POS-SRI/model/response"
	"fmt"
	"net/http"
	"strconv"
)

func Input_Jenis_Barang(Request request.Input_Jenis_Barang_Request) (response.Response, error) {
	var res response.Response

	con := db.CreateConGorm().Table("jenis_barang")

	co := 0

	err := con.Select("co").Order("co DESC").Limit(1).Scan(&co)

	Request.Co = co + 1
	Request.Kode_jenis_barang = "JB-" + strconv.Itoa(Request.Co)

	fmt.Println(co)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = co
		return res, err.Error
	}

	err = con.Select("co", "kode_jenis_barang", "nama_jenis_barang", "kode_gudang").Create(&Request)

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

func Read_Jenis_Pembayaran(Request request.Input_Jenis_Barang_Request) (response.Response, error) {
	var res response.Response

	return res, nil
}
