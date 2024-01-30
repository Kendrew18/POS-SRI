package stok_keluar

import (
	"POS-SRI/db"
	"POS-SRI/model/request"
	"POS-SRI/model/response"
	"POS-SRI/tools"
	"net/http"
	"strconv"
	"time"
)

func Input_Stock_Keluar(Request request.Input_Stock_Keluar_Request, Request_barang request.Input_Barang_Stock_Keluar_Request) (response.Response, error) {
	var res response.Response
	var bskm []request.Input_Barang_Stock_Keluar_Masuk_Request
	var data_bskm request.Input_Barang_Stock_Keluar_Masuk_Request

	//Masukkan ke tabel stock_keluar
	con := db.CreateConGorm()

	co := 0

	err := con.Select("co").Order("co DESC").Limit(1).Scan(&co)

	Request.Co = co + 1
	Request.Kode_stock_keluar = "SK-" + strconv.Itoa(Request.Co)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	date, _ := time.Parse("02-01-2006", Request.Tanggal)
	Request.Tanggal = date.Format("2006-01-02")

	err = con.Select("co", "kode_stock_keluar", "tanggal", "surat_jalan", "tujuan", "tipe", "kode_gudang").Create(&Request)

	//Masukkan ke tabel barang stock_keluar
	kode_stock := tools.String_Separator_To_String(Request_barang.Kode_barang)
	Berat_barang := tools.String_Separator_To_float64(Request_barang.Berat_barang)
	kode_lot := tools.String_Separator_To_String(Request_barang.Kode_lot)
	kode_grade_barang := tools.String_Separator_To_String(Request_barang.Kode_grade_barang)

	var temp request.Input_Barang_Stock_Keluar_V2_Request
	var barang_V2 []request.Input_Barang_Stock_Keluar_V2_Request

	con_barang := db.CreateConGorm().Table("barang_stock_keluar")

	co_barang := 0

	err = con_barang.Select("co").Order("co DESC").Limit(1).Scan(&co_barang)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = barang_V2
		return res, err.Error
	}

	Total_berat_barang := 0.0

	for i := 0; i < len(kode_stock); i++ {

		temp.Co = co_barang + 1 + i
		temp.Kode_barang_stock_keluar = "BSK-" + strconv.Itoa(temp.Co)

		temp.Kode_stock_keluar = Request.Kode_stock_keluar
		temp.Kode_barang = kode_stock[i]
		temp.Berat_barang = Berat_barang[i]
		temp.Kode_lot = kode_lot[i]
		temp.Kode_grade_barang = kode_grade_barang[i]

		data_bskm.Kode_barang_keluar_masuk = temp.Kode_barang_stock_keluar
		data_bskm.Kode = temp.Kode_stock_keluar
		data_bskm.Berat_barang = temp.Berat_barang
		data_bskm.Kode_lot = temp.Kode_lot
		data_bskm.Kode_barang = temp.Kode_barang
		data_bskm.Kode_grade_barang = temp.Kode_grade_barang
		data_bskm.Keterangan = "KELUAR"

		Total_berat_barang += temp.Berat_barang

		barang_V2 = append(barang_V2, temp)
		bskm = append(bskm, data_bskm)

	}

	err = con_barang.Select("co", "kode_barang_stock_keluar", "kode_barang", "kode_grade_barang", "kode_stock_keluar", "berat_barang", "kode_lot").Create(&barang_V2)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = barang_V2
		return res, err.Error
	}

	//Masukkan ke tabel stock keluar masuk
	var I_SKM request.Input_Stock_Keluar_Masuk_Request

	co_kode_stock_keluar_masuk := 0

	err = con.Table("stock_keluar_masuk").Select("co").Order("co DESC").Limit(1).Scan(&co_kode_stock_keluar_masuk)

	I_SKM.Co = co_kode_stock_keluar_masuk + 1
	I_SKM.Kode = Request.Kode_stock_keluar
	I_SKM.Tanggal = Request.Tanggal

	err = con.Table("stock_keluar_masuk").Select("co", "kode", "tanggal").Create(&I_SKM)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	//Masukkan ke tabel barang stock keluar masuk
	co_detail := 0

	err = con.Table("barang_stock_keluar_masuk").Select("co").Order("co DESC").Limit(1).Scan(&co_detail)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	for i := 0; i < len(bskm); i++ {
		bskm[i].Co = co_detail + 1 + i
	}

	err = con.Table("barang_stock_keluar_masuk").Select("co", "kode_barang_keluar_masuk", "kode", "kode_lot", "kode_barang", "kode_grade_barang", "berat_barang", "kadar_air", "penyusutan", "keterangan").Create(&bskm)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
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

func Read_Stock_Keluar(Request request.Read_Stock_Kelaur_Request, Request_filter request.Read_Stock_Keluar_Filter_Request) (response.Response, error) {

	var res response.Response
	var arr_data []response.Read_Stock_Keluar_Response
	var err error

	con := db.CreateConGorm()

	statement := "SELECT stock_keluar.kode_stock_keluar, surat_jalan, DATE_FORMAT(tanggal, '%d-%m-%Y') AS tanggal, SUM(berat_barang) AS total_berat_barang, tujuan, tipe, satuan FROM stock_keluar_masuk JOIN barang_stock_keluar bsk ON bsk.kode_stock_keluar = stock_keluar.kode_stock_keluar JOIN Barang b ON b.kode_barang = bsk.kode_barang WHERE stock_keluar.kode_gudang = '" + Request.Kode_gudang + "'"

	if Request_filter.Tujuan != "" {
		statement += " && tujuan = '" + Request_filter.Tujuan + "'"
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
