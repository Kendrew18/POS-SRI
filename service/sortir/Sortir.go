package sortir

import (
	"POS-SRI/db"
	"POS-SRI/model/request"
	"POS-SRI/model/response"
	"math"
	"net/http"
	"strconv"
	"time"
)

func Update_Status_Sortir(Request request.Update_Status_Sortir_Request, Request_kode request.Kode_Stock_Masuk_Request) (response.Response, error) {
	var res response.Response
	con := db.CreateConGorm().Table("stock_masuk")
	status_sortir := -1

	err := con.Select("status_sortir").Where("kode_stock_masuk = ?", Request_kode.Kode_stock_masuk).Scan(&status_sortir)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	if status_sortir != 2 {
		if Request.Status_sortir == 2 {

			con := db.CreateConGorm().Table("stock_masuk")

			err := con.Where("kode_stock_masuk = ?", Request_kode.Kode_stock_masuk).Select("status_sortir").Updates(&Request)

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

		} else if Request.Status_sortir == 1 && status_sortir != 1 {
			con := db.CreateConGorm()

			err := con.Table("stock_masuk").Where("kode_stock_masuk = ?", Request_kode.Kode_stock_masuk).Select("status_sortir").Updates(&Request)

			if err.Error != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = Request
				return res, err.Error
			}

			var input_sortir []request.Input_Sortir_Request

			err = con.Table("stock_masuk").Select("stock_masuk.kode_stock_masuk", "kode_lot", "kode_barang", "kode_grade_barang", "berat_barang", "kode_gudang").Joins("JOIN barang_stock_masuk bsm ON bsm.kode_stock_masuk = stock_masuk.kode_stock_masuk").Where("stock_masuk.kode_stock_masuk = ?", Request_kode.Kode_stock_masuk).Scan(&input_sortir)

			if err.Error != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = Request
				return res, err.Error
			}

			con_qc := db.CreateConGorm().Table("sortir")

			co := 0

			err = con_qc.Select("co").Order("co DESC").Limit(1).Scan(&co)

			if err.Error != nil {
				res.Status = http.StatusNotFound
				res.Message = "Status Not Found"
				res.Data = Request
				return res, err.Error
			}

			date := time.Now()
			Tanggal_sortir := date.Format("2006-01-02")

			for i := 0; i < len(input_sortir); i++ {

				input_sortir[i].Co = co + 1 + i
				input_sortir[i].Kode_sortir = "S-" + strconv.Itoa(input_sortir[i].Co)
				input_sortir[i].Tanggal = Tanggal_sortir

			}

			err = con.Table("sortir").Select("co", "kode_sortir", "kode_stock_masuk", "kode_lot", "tanggal", "kode_barang", "kode_grade_barang", "berat_barang", "penyusutan", "kode_gudang").Create(&input_sortir)

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

func Read_Sortir(Request request.Read_Sortir_Request, Request_filter request.Read_Sortir_Filter_Request) (response.Response, error) {

	var res response.Response
	var arr_data []response.Read_Sortir_Response
	var header_sortir response.Read_header_sortir_response
	var err error

	con := db.CreateConGorm()

	err = con.Table("sortir").Select("SUM(penyusutan) AS penyusutan_global").Where("kode_gudang = ? AND sortir.status = 0", Request.Kode_gudang).Scan(&header_sortir).Error

	statement := "SELECT kode_sortir, kode_lot, DATE_FORMAT(tanggal, '%d-%m-%Y') AS tanggal, sortir.kode_barang nama_barang, sortir.kode_grade_barang, nama_grade_barang, berat_barang, berat_setelah_sortir, penyusutan,kode_stock_masuk FROM sortir JOIN barang b on b.kode_barang = sortir.kode_barang JOIN grade_barang gb ON gb.kode_grade_barang = sortir.kode_grade_barang WHERE sortir.kode_gudang = '" + Request.Kode_gudang + "' && sortir.status = 0"

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

	statement += " ORDER BY sortir.tanggal DESC"

	err = con.Raw(statement).Scan(&arr_data).Error

	header_sortir.Sortir = arr_data

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = header_sortir
		return res, err
	}

	if arr_data == nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = header_sortir

	} else {
		res.Status = http.StatusOK
		res.Message = "Suksess"
		res.Data = header_sortir
	}

	return res, nil
}

func Read_Detail_Sortir(Request request.Read_Detail_Sortir_Request) (response.Response, error) {

	var res response.Response
	var arr_data response.Read_Detail_Sortir_Response
	var data []response.Read_Barang_Sortir_Response

	con := db.CreateConGorm()

	err := con.Table("sortir").Select("kode_sortir", "kode_lot", "DATE_FORMAT(tanggal, '%d-%m-%Y') AS tanggal", "sortir.kode_barang", "nama_barang", "berat_barang", "berat_setelah_sortir", "penyusutan").Joins("join barang b on b.kode_barang = sortir.kode_barang").Where("kode_sortir = ?", Request.Kode_sortir).Scan(&arr_data)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Update Error"
		res.Data = Request
		return res, err.Error
	}

	kode := ""

	err = con.Table("barang_sortir").Select("kode_barang_sortir").Where("kode_sortir = ?", Request.Kode_sortir).Scan(&kode)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Update Error"
		res.Data = Request
		return res, err.Error
	}

	if kode == "" {
		err = con.Table("grade_barang").Select("kode_grade_barang", "nama_grade_barang", "satuan").Where("kode_barang = ?", arr_data.Kode_barang).Scan(&data)

		arr_data.Barang_sortir = data
	} else {
		//var data_2 []response.Read_Barang_Sortir_Response
		err = con.Table("barang_sortir").Select("COALESCE(IF(kode_barang_sortir = NULL,'',kode_barang_sortir),'') AS kode_barang_sortir", "gb.kode_grade_barang", "nama_grade_barang", "berat_setelah_sortir", "kadar_air", "persentase", "satuan").Joins("Right join grade_barang gb on gb.kode_grade_barang = barang_sortir.kode_grade_barang AND kode_sortir = ?", Request.Kode_sortir).Order("barang_sortir.co ASC").Scan(&data)

		//fmt.Println(err.Statement)

		// co := 0

		// err = con.Table("grade_barang").Select("COUNT(kode_grade_barang)").Where("kode_barang = ?", arr_data.Kode_barang).Scan(&co)

		// if co != len(data) {

		// 	where_statement := "kode_barang NOT IN ("

		// 	for i := 0; i < len(data); i++ {
		// 		if i < len(data)-1 {
		// 			where_statement += data[i].Kode_grade_barang + ", "
		// 		} else if i == len(data)-1 {
		// 			where_statement += data[i].Kode_grade_barang + ")"
		// 		}
		// 	}

		// 	err = con.Table("grade_barang").Select("kode_grade_barang", "nama_grade_barang").Where(where_statement, arr_data.Kode_barang).Scan(&data_2)

		// 	data = append(data, data_2...)
		// }

		arr_data.Barang_sortir = data
	}

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

func Update_Berat_Barang_Setelah_Sortir(Request request.Input_Barang_Sortir_Request) (response.Response, error) {
	var res response.Response

	con := db.CreateConGorm()

	if Request.Kode_barang_sortir == "" {

		co := 0

		err := con.Table("barang_sortir").Select("co").Order("co DESC").Limit(1).Scan(&co)

		Request.Co = co + 1
		Request.Kode_barang_sortir = "BS-" + strconv.Itoa(Request.Co)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}

		err = con.Table("barang_sortir").Select("IFNULL(SUM(berat_setelah_sortir) , 0) AS persentase").Where("kode_sortir=?", Request.Kode_sortir).Scan(&Request.Persentase)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}

		Request.Persentase = math.Round((Request.Berat_setelah_sortir/(Request.Persentase+Request.Berat_setelah_sortir))*100*1000) / 1000

		err = con.Table("barang_sortir").Select("co", "kode_barang_sortir", "kode_sortir", "kode_grade_barang", "berat_setelah_sortir", "kadar_air", "persentase").Create(&Request)

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

	} else if Request.Kode_barang_sortir != "" {

		err := con.Table("barang_sortir").Select("IFNULL(SUM(berat_setelah_sortir) , 0) AS persentase").Where("kode_sortir=?", Request.Kode_sortir).Scan(&Request.Persentase)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}

		Request.Persentase = math.Round((Request.Berat_setelah_sortir/(Request.Persentase+Request.Berat_setelah_sortir))*100*1000) / 1000

		err = con.Table("barang_sortir").Where("kode_barang_sortir = ?", Request.Kode_barang_sortir).Select("berat_setelah_sortir", "kadar_air", "persentase").Updates(&Request)

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

	}

	return res, nil
}
