package stock_utilitas

import (
	"POS-SRI/db"
	"POS-SRI/model/request"
	"POS-SRI/model/response"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func Input_Stock_Utilitas(Request request.Input_Stock_Utilitas_Request) (response.Response, error) {
	var res response.Response

	con := db.CreateConGorm().Table("stock_utilitas")

	co := 0

	err := con.Select("co").Order("co DESC").Limit(1).Scan(&co)

	Request.Co = co + 1
	Request.Kode_stock_utilitas = "STU-" + strconv.Itoa(Request.Co)

	date, _ := time.Parse("02-01-2006", Request.Tanggal)
	Request.Tanggal = date.Format("2006-01-02")

	fmt.Println(co)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = co
		return res, err.Error
	}

	err = con.Select("co", "kode_stock_utilitas", "nama_stock_utilitas", "jumlah", "tanggal", "kode_gudang").Create(&Request)

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

func Read_Stock_Utilitas(Request request.Read_Stock_Utilitas_Request) (response.Response, error) {
	var res response.Response

	var stock_utilitas []response.Stock_Utilitas_Response

	con := db.CreateConGorm()

	err := con.Table("stock_utilitas").Select("kode_stock_utilitas", "nama_stock_utilitas", "tanggal", "jumlah").Where("kode_gudang = ?", Request.Kode_gudang).Scan(&stock_utilitas)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	if stock_utilitas == nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = stock_utilitas

	} else {
		res.Status = http.StatusOK
		res.Message = "Suksess"
		res.Data = stock_utilitas
	}

	return res, nil
}

func Update_Stock_Utilitas(Request request.Update_Stock_Utilitas_Request, Request_kode request.Update_Stock_Utilitas_Kode_Request) (response.Response, error) {
	var res response.Response

	con := db.CreateConGorm().Table("stock_utilitas")

	date, _ := time.Parse("02-01-2006", Request.Tanggal)
	Request.Tanggal = date.Format("2006-01-02")

	err := con.Where("kode_stock_utilitas = ?", Request_kode.Kode_stock_utilitas).Select("nama_stock_utilitas", "tanggal", "jumlah").Updates(&Request)

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
