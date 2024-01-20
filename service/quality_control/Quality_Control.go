package quality_control

import (
	"POS-SRI/db"
	"POS-SRI/model/request"
	"POS-SRI/model/response"
	"database/sql"
	"net/http"
	"time"
)

func Read_Quality_Control(Request request.Read_quality_control_Request, Request_filter request.Read_Quality_Control_Filter_Request) (response.Response, error) {

	var res response.Response
	var arr_data []response.Read_Quality_Control_Response
	var data response.Read_Quality_Control_Response
	var rows *sql.Rows
	var err error

	con := db.CreateConGorm()

	statement := "SELECT kode_quality_control, kode_lot, DATE_FORMAT(tanggal_masuk, '%d-%m-%Y') AS tanggal_masuk, kode_pre_order, kode_supplier, nama_supplier, kode_barang, nama_barang, berat_barang, berat_rill, satuan.nama_satuan, status FROM quality_control JOIN supplier s ON s.kode_supplier = quality_control.kode_supplier JOIN barang b ON b.kode_barang = quality_control.kode_barang JOIN satuan st ON st.kode_satuan = b.kode_satuan WHERE pre_order.kode_gudang = '" + Request.Kode_gudang + "'"

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

	statement += " ORDER BY quality_control.tanggal_masuk DESC"

	rows, err = con.Raw(statement).Rows()

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = arr_data
		return res, err
	}

	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(&data.Kode_quality_control, &data.Kode_lot, &data.Tanggal_masuk, &data.Kode_supplier, &data.Nama_supplier, &data.Kode_barang, &data.Nama_barang, &data.Berat_barang, &data.Berat_barang_rill, &data.Nama_satuan, &data.Status)

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

func Update_Pre_Order(Request request.Update_Berat_Barang_Rill_Request, Request_kode request.Update_Quality_Control_Kode_Request) (response.Response, error) {
	var res response.Response

	check := -1
	con_check := db.CreateConGorm().Table("quality_control")

	err := con_check.Select("status").Where("kode_quality_control = ?", Request_kode.Kode_quality_control).Scan(&check)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Update Error"
		res.Data = Request
		return res, err.Error
	}
	if check == 0 || check == 2 {

		con := db.CreateConGorm().Table("quality_control")

		err = con.Where("kode_quality_control = ?", Request_kode.Kode_quality_control).Select("berat_barang_rill").Updates(&Request)

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
