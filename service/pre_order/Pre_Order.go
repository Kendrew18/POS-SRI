package pre_order

import (
	"POS-SRI/db"
	"POS-SRI/model/request"
	"POS-SRI/model/response"
	"POS-SRI/tools"
	"database/sql"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"
)

func Input_Pre_Order(Request request.Input_Pre_Order_Request, Request_Barang request.Input_Barang_Pre_Order_Request) (response.Response, error) {

	var res response.Response

	con := db.CreateConGorm().Table("pre_order")

	co := 0

	err := con.Select("co").Order("co DESC").Limit(1).Scan(&co)

	Request.Co = co + 1
	Request.Kode_pre_order = "PO-" + strconv.Itoa(Request.Co)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	//0 = pending
	//1 = success
	//2 = ditolak / Cancel

	date, _ := time.Parse("02-01-2006", Request.Tanggal)
	Request.Tanggal = date.Format("2006-01-02")

	date, _ = time.Parse("02-01-2006", Request.Tanggal_ETD)
	Request.Tanggal_ETD = date.Format("2006-01-02")

	Request.Status = 0

	err = con.Select("co", "kode_pre_order", "kode_lot", "tanggal", "nama_supplier", "tanggal_etd", "kode_jenis_pembayaran", "kode_gudang", "status").Create(&Request)

	kode_stock := tools.String_Separator_To_String(Request_Barang.Kode_barang)
	Berat_barang := tools.String_Separator_To_float64(Request_Barang.Berat_barang)
	harga_pokok := tools.String_Separator_To_Int64(Request_Barang.Harga)
	kode_grade_barang := tools.String_Separator_To_String(Request_Barang.Kode_grade_barang)

	var barang_V2 []request.Input_Barang_Pre_Order_V2_Request

	con_barang := db.CreateConGorm().Table("barang_pre_order")

	co_barang := 0

	err = con_barang.Select("co").Order("co DESC").Limit(1).Scan(&co_barang)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = barang_V2
		return res, err.Error
	}

	for i := 0; i < len(kode_stock); i++ {

		barang_V2[i].Co = co_barang + 1 + i
		barang_V2[i].Kode_barang_pre_order = "BPO-" + strconv.Itoa(barang_V2[i].Co)

		barang_V2[i].Kode_pre_order = Request.Kode_pre_order
		barang_V2[i].Kode_barang = kode_stock[i]
		barang_V2[i].Berat_barang = Berat_barang[i]
		barang_V2[i].Harga = harga_pokok[i]
		barang_V2[i].Kode_grade_barang = kode_grade_barang[i]
		barang_V2[i].Sub_total = int64(math.Round(float64(harga_pokok[i]) * Berat_barang[i]))

	}

	err = con_barang.Select("co", "kode_barang_pre_order", "kode_barang", "kode_grade_barang", "kode_pre_order", "berat_barang", "harga", "sub_total").Create(&barang_V2)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = barang_V2
		return res, err.Error
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

func Read_Pre_Order(Request request.Read_Pre_Order_Request, Request_filter request.Read_Pre_Order_Filter_Request) (response.Response, error) {

	var res response.Response
	var arr_data []response.Read_Pre_Order_Response
	var data response.Read_Pre_Order_Response
	var rows *sql.Rows
	var err error

	con := db.CreateConGorm()

	statement := "SELECT pre_order.kode_pre_order, kode_lot, DATE_FORMAT(tanggal, '%d-%m-%Y') AS tanggal, nama_supplier, DATE_FORMAT(tanggal_etd, '%d-%m-%Y') AS tanggal_etd,  IF(tanggal_rtd = '0001-01-01','-', DATE_FORMAT(tanggal_rtd, '%d-%m-%Y')) AS tanggal_rtd, sum(berat_barang), pre_order.kode_gudang, nama_gudang, sum(sub_total) FROM pre_order JOIN gudang on gudang.kode_gudang = pre_order.kode_gudang JOIN barang_pre_order bpo ON bpo.kode_pre_order = pre_order.kode_pre_order WHERE pre_order.kode_gudang = '" + Request.Kode_gudang + "'"

	if Request_filter.Tanggal_etd != "" {
		date, _ := time.Parse("02-01-2006", Request_filter.Tanggal_etd)
		Request_filter.Tanggal_etd = date.Format("2006-01-02")

		statement += " && pre_order.tanggal_etd = '" + Request_filter.Tanggal_etd + "'"
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

	statement += " GROUP BY pre_order.kode_pre_order ORDER BY pre_order.co DESC"

	rows, err = con.Raw(statement).Rows()

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = arr_data
		return res, err
	}

	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(&data.Kode_pre_order, &data.Kode_lot, &data.Tanggal, &data.Nama_supplier, &data.Tanggal_etd, &data.Tanggal_rtd, &data.Total_berat, &data.Kode_gudang, &data.Nama_gudang, &data.Total_harga)

		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = data
			return res, err
		}

		arr_data = append(arr_data, data)

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

func Read_Detail_Pre_Order(Request request.Read_Detail_Pre_Order_Request) (response.Response, error) {

	var res response.Response
	var arr_data []response.Detail_Pre_Order_Response
	var data response.Detail_Pre_Order_Response
	var rows *sql.Rows
	var err error

	con := db.CreateConGorm()

	statement := "SELECT pre_order.kode_pre_order, kode_lot, DATE_FORMAT(tanggal, '%d-%m-%Y') as tanggal, nama_supplier, sum(berat_barang), status FROM pre_order JOIN barang_pre_order bpo ON bpo.kode_pre_order = pre_order.kode_pre_order WHERE pre_order.kode_pre_order = '" + Request.Kode_pre_order + "'"

	statement += " GROUP BY pre_order.kode_pre_order ORDER BY pre_order.tanggal DESC"

	rows, err = con.Raw(statement).Rows()

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = arr_data
		return res, err
	}

	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(&data.Kode_pre_order, &data.Kode_lot, &data.Tanggal, &data.Nama_supplier, &data.Total_berat, &data.Status)

		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = data
			return res, err
		}

		con_detail := db.CreateConGorm().Table("barang_pre_order")
		var detail_data []response.Read_Barang_Pre_Order_Response

		err = con_detail.Select("kode_barang_pre_order", "b.kode_barang", "nama_barang", "gb.kode_grade_barang", "nama_grade_barang", "berat_barang", "gb.satuan", "harga", "sub_total").Joins("join barang b on barang_pre_order.kode_barang = b.kode_barang").Joins("join grade_barang gb on barang_pre_order.kode_grade_barang = gb.kode_grade_barang").Where("kode_pre_order = ?", data.Kode_pre_order).Scan(&detail_data).Error

		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = data
			return res, err
		}

		data.Detail_barang_pre_order = detail_data

		arr_data = append(arr_data, data)

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

func Update_Pre_Order(Request request.Update_Pre_order_Request, Request_kode request.Update_Pre_Order_Kode_Request) (response.Response, error) {
	var res response.Response

	check := -1
	con_check := db.CreateConGorm().Table("pre_order")

	err := con_check.Select("status").Joins("JOIN barang_pre_order bpo ON bpo.kode_pre_order = pre_order.kode_pre_order ").Where("kode_barang_pre_order = ?", Request_kode.Kode_barang_pre_order).Scan(&check)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Update Error"
		res.Data = Request
		return res, err.Error
	}
	if check == 0 || check == 2 {

		con := db.CreateConGorm().Table("barang_pre_order")

		Request.Sub_total = int64(math.Round(float64(Request.Harga) * Request.Berat_barang))

		fmt.Println(Request)

		err = con.Where("kode_barang_pre_order = ?", Request_kode.Kode_barang_pre_order).Select("kode_barang", "kode_grade_barang", "berat_barang", "harga", "sub_total").Updates(&Request)

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

func Delete_Pre_Order(Request request.Update_Pre_Order_Kode_Request) (response.Response, error) {
	var res response.Response

	check := -1
	con_check := db.CreateConGorm().Table("pre_order")

	err := con_check.Select("status").Joins("JOIN barang_pre_order bpo ON bpo.kode_pre_order = pre_order.kode_pre_order ").Where("kode_barang_pre_order = ?", Request.Kode_barang_pre_order).Scan(&check)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Update Error"
		res.Data = Request
		return res, err.Error
	}
	if check == 0 || check == 2 {

		con := db.CreateConGorm()

		data := ""

		err = con.Table("pre_order").Select("pre_order.kode_pre_order").Joins("JOIN barang_pre_order bpo ON bpo.kode_pre_order = pre_order.kode_pre_order ").Where("kode_barang_pre_order = ?", Request.Kode_barang_pre_order).Scan(&data)

		fmt.Println(data)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Update Error"
			res.Data = Request
			return res, err.Error
		}

		con_barang := db.CreateConGorm().Table("barang_pre_order")

		err = con_barang.Where("kode_barang_pre_order = ?", Request.Kode_barang_pre_order).Delete("")

		kode_barang := ""

		con_barang_check := db.CreateConGorm().Table("barang_pre_order")

		err = con_barang_check.Select("kode_barang_pre_order").Where("kode_pre_order=?", data).Limit(1).Scan(&kode_barang)

		fmt.Println(kode_barang)

		if kode_barang == "" {
			fmt.Println("masuk")

			err = con.Table("pre_order").Where("kode_pre_order = ?", data).Delete("")

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
	} else {
		res.Status = http.StatusNotFound
		res.Message = "Barang Tidak dapat di update"
		res.Data = Request
		return res, err.Error
	}
	return res, nil
}

func Update_Status_Pre_Order(Request request.Update_Status_Pre_Order_Request, Request_kode request.Kode_Pre_Order_Request) (response.Response, error) {
	var res response.Response
	con := db.CreateConGorm().Table("pre_order")
	status := -1

	err := con.Select("status").Where("kode_pre_order = ?", Request_kode.Kode_pre_order).Scan(&status)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	if status != 1 {
		if Request.Status == 2 || Request.Status == 0 {

			con := db.CreateConGorm().Table("pre_order")

			err := con.Where("kode_pre_order = ?", Request_kode.Kode_pre_order).Select("status").Updates(&Request)

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

			err := con.Table("pre_order").Where("kode_pre_order = ?", Request_kode.Kode_pre_order).Select("status").Updates(&Request)

			if err.Error != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = Request
				return res, err.Error
			}

			var input_qc request.Input_Quality_Control_Request

			err = con.Table("pre_order").Select("kode_lot", "kode_pre_order", "nama_supplier", "kode_gudang").Where("kode_pre_order = ?", Request_kode.Kode_pre_order).Scan(&input_qc)

			if err.Error != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = Request
				return res, err.Error
			}

			con_qc := db.CreateConGorm().Table("quality_control")

			co := 0

			err = con_qc.Select("co").Order("co DESC").Limit(1).Scan(&co)

			input_qc.Co = co + 1
			input_qc.Kode_quality_control = "QC-" + strconv.Itoa(input_qc.Co)

			if err.Error != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = Request
				return res, err.Error
			}

			date := time.Now()
			input_qc.Tanggal_masuk = date.Format("2006-01-02")

			var input_barang_QC []request.Input_Barang_Quality_Control_Request

			err = con.Table("barang_pre_order").Select("kode_barang", "kode_grade_barang", "berat_barang", "harga", "sub_total").Where("kode_pre_order = ?", Request_kode.Kode_pre_order).Scan(&input_barang_QC)

			fmt.Println(input_barang_QC)

			co_bqc := 0

			err = con.Table("barang_quality_control").Select("co").Order("co DESC").Limit(1).Scan(&co_bqc)

			if err.Error != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = Request
				return res, err.Error
			}

			for i := 0; i < len(input_barang_QC); i++ {

				input_barang_QC[i].Co = co_bqc + 1 + i
				input_barang_QC[i].Kode_barang_quality_control = "BQC-" + strconv.Itoa(input_barang_QC[i].Co)
				input_barang_QC[i].Kode_quality_control = input_qc.Kode_quality_control

			}

			err = con.Table("quality_control").Select("co", "kode_quality_control", "kode_lot", "kode_pre_order", "tanggal_masuk", "nama_supplier", "kode_gudang").Create(&input_qc)

			if err.Error != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = Request
				return res, err.Error
			}

			err = con.Table("barang_quality_control").Select("co", "kode_barang_quality_control", "kode_quality_control", "kode_barang", "kode_grade_barang", "berat_barang", "harga", "sub_total").Create(&input_barang_QC)

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
