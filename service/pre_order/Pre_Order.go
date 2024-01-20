package pre_order

import (
	"POS-SRI/db"
	"POS-SRI/model/request"
	"POS-SRI/model/response"
	"POS-SRI/tools"
	"database/sql"
	"fmt"
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

	err = con.Select("co", "kode_pre_order", "kode_lot", "tanggal", "kode_supplier", "tanggal_etd", "kode_jenis_pembayaran", "kode_gudang", "status").Create(&Request)

	kode_stock := tools.String_Separator_To_String(Request_Barang.Kode_barang)
	Berat_barang := tools.String_Separator_To_float64(Request_Barang.Berat_barang)
	harga_pokok := tools.String_Separator_To_Int64(Request_Barang.Harga)

	for i := 0; i < len(kode_stock); i++ {
		var barang_V2 request.Input_Barang_Pre_Order_V2_Request

		con_barang := db.CreateConGorm().Table("barang_pre_order")

		co := 0

		err := con_barang.Select("co").Order("co DESC").Limit(1).Scan(&co)

		barang_V2.Co = co + 1
		barang_V2.Kode_barang_pre_order = "BPO-" + strconv.Itoa(barang_V2.Co)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = barang_V2
			return res, err.Error
		}

		barang_V2.Kode_pre_order = Request.Kode_pre_order
		barang_V2.Kode_barang = kode_stock[i]
		barang_V2.Berat_barang = Berat_barang[i]
		barang_V2.Harga = harga_pokok[i]

		fmt.Println(barang_V2)

		err = con_barang.Select("co", "kode_barang_pre_order", "kode_barang", "kode_pre_order", "berat_barang", "harga").Create(&barang_V2)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = barang_V2
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

func Read_Pre_Order(Request request.Read_Pre_Order_Request, Request_filter request.Read_Pre_Order_Filter_Request) (response.Response, error) {

	var res response.Response
	var arr_data []response.Read_Pre_Order_Response
	var data response.Read_Pre_Order_Response
	var rows *sql.Rows
	var err error

	con := db.CreateConGorm()

	statement := "SELECT pre_order.kode_pre_order, DATE_FORMAT(tanggal, '%d-%m-%Y') AS tanggal, kode_lot, kode_supplier, nama_supplier, DATE_FORMAT(tanggal_etd, '%d-%m-%Y') AS tanggal_etd, DATE_FORMAT(tanggal_rtd, '%d-%m-%Y') AS tanggal_rtd, sum(berat_barang), pre_order.kode_gudang, nama_gudang FROM pre_order JOIN supplier s ON s.kode_supplier = pre_order.kode_supplier JOIN barang_pre_order bpo ON bpo.kode_pre_order = pre_order.kode_pre_order WHERE pre_order.kode_gudang = '" + Request.Kode_gudang + "'"

	if Request_filter.Kode_supplier != "" {
		statement += " && pre_order.kode_supplier = '" + Request_filter.Kode_supplier + "'"
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

		err = rows.Scan(&data.Kode_pre_order, &data.Tanggal, &data.Kode_lot, &data.Kode_supplier, &data.Nama_supplier, &data.Tanggal_etd, &data.Tanggal_rtd, &data.Total_berat, &data.Kode_satuan, &data.Nama_satuan, &data.Kode_gudang, &data.Nama_gudang)

		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = data
			return res, err
		}

		// con_detail := db.CreateConGorm().Table("barang_pre_order")
		// var detail_data []response.Read_Barang_Pre_Order_Response

		// err = con_detail.Select("kode_barang_pre_order", "nama_barang", "DATE_FORMAT(tanggal_kadaluarsa, '%d-%m-%Y') AS tanggal_kadaluarsa", "jumlah_barang", "harga").Joins("join stock s on barang_pre_order.kode_stock = s.kode_stock").Where("kode_pre_order = ?", data.Kode_pre_order).Scan(&detail_data).Error

		// if err != nil {
		// 	res.Status = http.StatusNotFound
		// 	res.Message = "Status Not Found"
		// 	res.Data = data
		// 	return res, err
		// }

		// data.Detail_stock_masuk = detail_data

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

	statement := "SELECT pre_order.kode_pre_order, DATE_FORMAT(tanggal, '%d-%m-%Y') AS tanggal, kode_lot, kode_supplier, nama_supplier, sum(berat_barang), pre_order.kode_jenis_pembayaran, nama_jenis_pembayaran, status FROM pre_order JOIN supplier s ON s.kode_supplier = pre_order.kode_supplier JOIN barang_pre_order bpo ON bpo.kode_pre_order = pre_order.kode_pre_order JOIN jenis_pembayaran jp on jp.kode_jenis_pembayaran = pre_order.kode_jenis_pembayaran WHERE pre_order.kode_pre_order = '" + Request.Kode_pre_order + "'"

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

		err = rows.Scan(&data.Kode_pre_order, &data.Tanggal_pre_order, &data.Kode_lot, &data.Kode_supplier, &data.Nama_supplier, &data.Total_berat, &data.Kode_jenis_pembayaran, &data.Nama_jenis_pembayaran, &data.Status_pembayaran)

		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = data
			return res, err
		}

		con_detail := db.CreateConGorm().Table("barang_pre_order")
		var detail_data []response.Read_Barang_Pre_Order_Response

		err = con_detail.Select("kode_barang_pre_order", "b.kode_barang", "nama_barang", "berat_barang", "harga").Joins("join barang b on barang_pre_order.kode_barang = s.kode_barang").Where("kode_pre_order = ?", data.Kode_pre_order).Scan(&detail_data).Error

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

		err = con.Where("kode_barang_pre_order = ?", Request_kode.Kode_barang_pre_order).Select("kode_barang", "berat_barang", "harga").Updates(&Request)

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

		con := db.CreateConGorm().Table("pre_order")

		data := ""

		err = con.Select("pre_order.kode_pre_order").Joins("JOIN barang_pre_order bpo ON bpo.kode_pre_order = pre_order.kode_pre_order ").Where("kode_barang_pre_order = ?", Request.Kode_barang_pre_order).Scan(&data)

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

			err = con.Where("kode_pre_order = ?", Request.Kode_barang_pre_order).Delete("")

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
			con := db.CreateConGorm().Table("pre_order")

			err := con.Where("kode_pre_order = ?", Request_kode.Kode_pre_order).Select("status").Updates(&Request)

			if err.Error != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = Request
				return res, err.Error
			}

			var input_qc []request.Input_Quality_Control_Request

			err = con.Select("kode_lot", "bpo.kdoe_pre_order", "kode_supplier", "kode_barang", "berat_barang", "kode_gudang").Joins("JOIN barang_pre_order bpo ON bpo.kode_pre_order = pre_order.kode_pre_order ").Where("kode_pre_order = ?", Request_kode.Kode_pre_order).Scan(input_qc)

			if err.Error != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = Request
				return res, err.Error
			}

			for i := 0; i < len(input_qc); i++ {

				con_qc := db.CreateConGorm().Table("quality_control")

				co := 0

				err := con_qc.Select("co").Order("co DESC").Limit(1).Scan(&co)

				input_qc[i].Co = co + 1
				input_qc[i].Kode_quality_control = "QC-" + strconv.Itoa(input_qc[i].Co)

				if err.Error != nil {
					res.Status = http.StatusNotFound
					res.Message = "Status Not Found"
					res.Data = Request
					return res, err.Error
				}

				date := time.Now()
				input_qc[i].Tanggal_masuk = date.Format("2006-01-02")

				err = con_qc.Select("co", "kode_quality_control", "kode_lot", "tanggal_masuk", "kode_pre_order", "kode_supplier", "kode_barang", "berat_barang", "kode_gudang").Create(&Request)

				if err.Error != nil {
					res.Status = http.StatusNotFound
					res.Message = "Status Not Found"
					res.Data = Request
					return res, err.Error
				}

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
