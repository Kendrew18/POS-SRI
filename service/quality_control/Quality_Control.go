package quality_control

import (
	"POS-SRI/db"
	"POS-SRI/model/request"
	"POS-SRI/model/response"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"
)

func Read_Quality_Control(Request request.Read_quality_control_Request, Request_filter request.Read_Quality_Control_Filter_Request) (response.Response, error) {

	var res response.Response
	var arr_data []response.Read_Quality_Control_Response
	var err error

	con := db.CreateConGorm()

	statement := "SELECT kode_quality_control, kode_lot, DATE_FORMAT(tanggal_masuk, '%d-%m-%Y') AS tanggal_masuk, kode_pre_order, nama_supplier, status FROM quality_control WHERE quality_control.kode_gudang = '" + Request.Kode_gudang + "'"

	if Request_filter.Kode_lot != "" {
		statement += " && kode_lot = '" + Request_filter.Kode_lot + "'"
	}

	if Request_filter.Tanggal_awal != "" && Request_filter.Tanggal_akhir != "" {

		date, _ := time.Parse("02-01-2006", Request_filter.Tanggal_awal)
		Request_filter.Tanggal_awal = date.Format("2006-01-02")

		date2, _ := time.Parse("02-01-2006", Request_filter.Tanggal_akhir)
		Request_filter.Tanggal_akhir = date2.Format("2006-01-02")

		statement += " AND (tanggal_masuk >= '" + Request_filter.Tanggal_awal + "' && tanggal_masuk <= '" + Request_filter.Tanggal_akhir + "' )"

	} else if Request_filter.Tanggal_awal != "" {

		date, _ := time.Parse("02-01-2006", Request_filter.Tanggal_awal)
		Request_filter.Tanggal_awal = date.Format("2006-01-02")

		statement += " && tanggal_masuk = '" + Request_filter.Tanggal_awal + "'"

	}

	statement += " ORDER BY quality_control.tanggal_masuk DESC"

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

func Update_Berat_Rill_Quality_Control(Request request.Update_Berat_Barang_Rill_Request, Request_kode request.Update_Quality_Control_Kode_Request) (response.Response, error) {
	var res response.Response

	check := -1
	con_check := db.CreateConGorm().Table("quality_control")

	err := con_check.Select("status").Joins("JOIN barang_quality_control bqc on bqc.kode_quality_control = quality_control.kode_quality_control").Where("kode_barang_quality_control = ?", Request_kode.Kode_barang_quality_control).Scan(&check)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Update Error"
		res.Data = Request
		return res, err.Error
	}
	if check == 0 || check == 2 {

		con := db.CreateConGorm()

		berat_barang := float64(0)

		err := con.Table("barang_quality_control").Select("berat_barang").Where("kode_barang_quality_control = ?", Request_kode.Kode_barang_quality_control).Scan(&berat_barang)

		Request.Penyusutan = berat_barang - Request.Berat_barang_rill - Request.Berat_barang_ditolak

		Request.Persentase = math.Round((Request.Penyusutan/(berat_barang-Request.Berat_barang_ditolak)*100)*1000) / 1000

		err = con.Table("barang_quality_control").Where("kode_barang_quality_control = ?", Request_kode.Kode_barang_quality_control).Select("berat_barang_rill", "berat_barang_ditolak", "penyusutan", "persentase", "kadar_air").Updates(&Request)

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
	} else {
		res.Status = http.StatusNotFound
		res.Message = "Barang Tidak dapat di update"
		res.Data = Request
		return res, err.Error
	}
	return res, nil
}

func Read_Detail_Quality_Control(Request request.Read_Detail_Quality_Control_Request) (response.Response, error) {

	var res response.Response
	var arr_data response.Read_Detail_Quality_Control_Response
	var data []response.Read_Barang_Quality_Control_Response

	con := db.CreateConGorm()

	err := con.Table("quality_control").Select("kode_quality_control", "kode_lot", "tanggal_masuk", "kode_pre_order", "nama_supplier").Where("kode_quality_control = ?", Request.Kode_quality_control).Scan(&arr_data)

	err = con.Table("barang_quality_control").Select("kode_barang_quality_control", "kode_quality_control", "barang_quality_control.kode_barang", "nama_barang", "barang_quality_control.kode_grade_barang", "nama_grade_barang", "berat_barang", "berat_barang_rill", "berat_barang_ditolak", "penyusutan", "persentase", "kadar_air", "harga", "sub_total").Joins("JOIN barang b on b.kode_barang = barang_quality_control.kode_barang").Joins("JOIN grade_barang gb ON gb.kode_grade_barang = barang_quality_control.kode_grade_barang").Where("kode_quality_control = ?", Request.Kode_quality_control).Scan(&data)

	arr_data.Barang_quality_control = data

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

func Update_Status_Quality_Control(Request request.Update_Status_Quality_Control_Request, Request_kode request.Kode_Quality_Control_Request) (response.Response, error) {
	var res response.Response
	con := db.CreateConGorm().Table("quality_control")
	status := -1

	err := con.Select("status").Where("kode_quality_control = ?", Request_kode.Kode_quality_control).Scan(&status)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	if status != 1 {
		if Request.Status == 2 || Request.Status == 0 {

			con := db.CreateConGorm().Table("quality_control")

			err := con.Where("kode_quality_control = ?", Request_kode.Kode_quality_control).Select("status").Updates(&Request)

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
		} else if Request.Status == 1 {
			con := db.CreateConGorm()

			err := con.Table("quality_control").Where("kode_quality_control = ?", Request_kode.Kode_quality_control).Select("status").Updates(&Request)

			if err.Error != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = Request
				return res, err.Error
			}

			var input_SM request.Input_Stock_Masuk_Request

			err = con.Table("quality_control").Select("kode_lot", "kode_quality_control", "nama_supplier", "kode_gudang").Where("kode_quality_control = ?", Request_kode.Kode_quality_control).Scan(&input_SM)

			fmt.Println(input_SM)

			if err.Error != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = Request
				return res, err.Error
			}

			con_qc := db.CreateConGorm().Table("stock_masuk")

			co := 0

			err = con_qc.Select("co").Order("co DESC").Limit(1).Scan(&co)

			input_SM.Co = co + 1
			input_SM.Kode_stock_masuk = "SM-" + strconv.Itoa(input_SM.Co)

			if err.Error != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = Request
				return res, err.Error
			}

			date := time.Now()
			input_SM.Tanggal = date.Format("2006-01-02")

			var input_barang_SM []request.Input_Barang_Stock_Masuk_Request

			err = con.Table("barang_quality_control").Select("kode_barang", "kode_grade_barang", "berat_barang_rill AS berat_barang", "penyusutan", "kadar_air", "harga", "sub_total").Where("kode_quality_control = ?", Request_kode.Kode_quality_control).Order("co ASC").Scan(&input_barang_SM)

			fmt.Println(input_barang_SM)

			co_bSM := 0

			err = con.Table("barang_stock_masuk").Select("co").Order("co DESC").Limit(1).Scan(&co_bSM)

			if err.Error != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = Request
				return res, err.Error
			}

			for i := 0; i < len(input_barang_SM); i++ {

				input_barang_SM[i].Co = co_bSM + 1 + i
				input_barang_SM[i].Kode_barang_stock_masuk = "BSM-" + strconv.Itoa(input_barang_SM[i].Co)
				input_barang_SM[i].Kode_stock_masuk = input_SM.Kode_stock_masuk
				input_barang_SM[i].Kode_lot = input_SM.Kode_lot

			}

			err = con.Table("stock_masuk").Select("co", "kode_stock_masuk", "kode_quality_control", "kode_lot", "tanggal", "nama_supplier", "kode_gudang").Create(&input_SM)

			if err.Error != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = Request
				return res, err.Error
			}

			err = con.Table("barang_stock_masuk").Select("co", "kode_barang_stock_masuk", "kode_stock_masuk", "kode_barang", "kode_grade_barang", "berat_barang", "kadar_air", "penyusutan", "harga", "sub_total").Create(&input_barang_SM)

			if err.Error != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = Request
				return res, err.Error
			}

			co_detail := 0

			err = con.Table("detail_barang").Select("co").Order("co DESC").Limit(1).Scan(&co_detail)

			if err.Error != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = Request
				return res, err.Error
			}

			for i := 0; i < len(input_barang_SM); i++ {
				input_barang_SM[i].Co = co_detail + 1 + i
			}

			err = con.Table("detail_barang").Select("co", "kode_barang_stock_masuk", "kode_stock_masuk", "kode_lot", "kode_barang", "kode_grade_barang", "berat_barang", "kadar_air", "penyusutan").Create(&input_barang_SM)

			if err.Error != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = Request
				return res, err.Error
			}

			res.Status = http.StatusOK
			res.Message = "Suksess"
			res.Data = map[string]int64{
				"rows": err.RowsAffected,
			}

		}
	} else {
		res.Status = http.StatusNotFound
		res.Message = "Tidah dapat di edit diakrenakan sudah sukses"
		res.Data = Request
	}
	return res, nil
}
