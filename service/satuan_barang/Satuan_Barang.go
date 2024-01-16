package satuan_barang

import (
	"POS-SRI/db"
	"POS-SRI/model/request"
	"POS-SRI/model/response"
	"net/http"
	"strconv"
)

func Input_Satuan_Barang(Request request.Input_Satuan_Barang_Request) (response.Response, error) {

	var res response.Response

	con := db.CreateConGorm().Table("satuan_barang")

	co := 0

	err := con.Select("co").Order("co DESC").Scan(&co)

	Request.Co = co + 1
	Request.Kode_satuan_barang = "SB-" + strconv.Itoa(Request.Co)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	err = con.Select("co", "kode_gudang", "kode_satuan_barang", "nama_satuan_barang").Create(&Request)

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

func Read_Satuan_Barang(Request request.Read_Satuan_Barang_Request) (response.Response, error) {

	var res response.Response
	var data []response.Read_Satuan_Barang_Response

	con := db.CreateConGorm().Table("satuan_barang")

	err := con.Select("kode_satuan_barang", "nama_satuan_barang").Where("kode_gudang = ?", Request.Kode_gudang).Order("co ASC").Scan(&data).Error

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
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
