package stock_masuk

import (
	"POS-SRI/db"
	"POS-SRI/model/request"
	"POS-SRI/model/response"
	"net/http"
	"time"
)

func Read_Stock_Masuk(Request request.Read_Stock_Masuk_Request, Request_filter request.Read_Stock_Masuk_Filter_Request) (response.Response, error) {

	var res response.Response
	var arr_data []response.Read_Stock_Masuk_Response
	var err error

	con := db.CreateConGorm()

	statement := "SELECT kode_stock_masuk, kode_lot, DATE_FORMAT(tanggal, '%d-%m-%Y') AS tanggal, kode_quality_control, nama_supplier FROM stock_masuk WHERE stock_masuk.kode_gudang = '" + Request.Kode_gudang + "'"

	if Request_filter.Kode_lot != "" {
		statement += " && kode_lot = '" + Request_filter.Kode_lot + "'"
	}

	if Request_filter.Tanggal_awal != "" && Request_filter.Tanggal_akhir != "" {

		date, _ := time.Parse("02-01-2006", Request_filter.Tanggal_awal)
		Request_filter.Tanggal_awal = date.Format("2006-01-02")

		date2, _ := time.Parse("02-01-2006", Request_filter.Tanggal_akhir)
		Request_filter.Tanggal_akhir = date2.Format("2006-01-02")

		statement += " AND (tanggal >= '" + Request_filter.Tanggal_awal + "' && tanggal <= '" + Request_filter.Tanggal_akhir + "' )"

	} else if Request_filter.Tanggal_awal != "" {

		date, _ := time.Parse("02-01-2006", Request_filter.Tanggal_awal)
		Request_filter.Tanggal_awal = date.Format("2006-01-02")

		statement += " && tanggal = '" + Request_filter.Tanggal_awal + "'"

	}

	statement += " ORDER BY stock_masuk.tanggal DESC"

	err = con.Raw(statement).Scan(&arr_data).Error

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = arr_data
		return res, err
	}

	if arr_data == nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = arr_data

	} else {
		res.Status = http.StatusOK
		res.Message = "Suksess"
		res.Data = arr_data
	}

	return res, nil
}

func Read_Detail_Stock_Masuk(Request request.Read_Detail_Stock_Masuk_Request) (response.Response, error) {

	var res response.Response
	var arr_data response.Read_Detail_Stock_Masuk_Response
	var data []response.Read_Barang_Stock_Masuk_Response

	con := db.CreateConGorm()

	err := con.Table("stock_masuk").Select("kode_stock_masuk", "kode_lot", "kode_quality_control", "tanggal", "nama_supplier").Where("kode_stock_masuk = ?", Request.Kode_stock_masuk).Scan(&arr_data)

	err = con.Table("barang_stock_masuk").Select("kode_barang_stock_masuk", "kode_stock_masuk", "barang_stock_masuk.kode_barang", "nama_barang", "barang_stock_masuk.kode_grade_barang", "nama_grade_barang", "berat_barang", "penyusutan", "kadar_air", "harga", "sub_total").Joins("JOIN barang b on b.kode_barang = barang_stock_masuk.kode_barang").Joins("JOIN grade_barang gb on gb.kode_grade_barang = barang_stock_masuk.kode_grade_barang").Where("kode_stock_masuk = ?", Request.Kode_stock_masuk).Scan(&data)

	arr_data.Detail_barang = data

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Update Error"
		res.Data = Request
		return res, err.Error
	}

	if data == nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = arr_data

	} else {
		res.Status = http.StatusOK
		res.Message = "Suksess"
		res.Data = arr_data
	}

	return res, nil
}
