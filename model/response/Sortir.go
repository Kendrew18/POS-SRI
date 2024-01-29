package response

type Read_header_sortir_response struct {
	Penyusutan_global float64                `json:"penyusutan_global"`
	Sortir            []Read_Sortir_Response `gorm:"-"`
}

type Read_Sortir_Response struct {
	Kode_sortir          string  `json:"kode_sortir"`
	Kode_lot             string  `json:"kode_lot"`
	Tanggal              string  `json:"tanggal"`
	Kode_barang          string  `json:"kode_barang"`
	Nama_barang          string  `json:"nama_barang"`
	Kode_grade_barang    string  `json:"kode_grade_barang"`
	Nama_grade_barang    string  `json:"nama_grade_barang"`
	Berat_barang         float64 `json:"berat_barang"`
	Berat_setelah_sortir float64 `json:"berat_setelah_sortir"`
	Penyusutan           float64 `json:"penyusutan"`
	Kode_stock_masuk     string  `json:"kdoe_stock_masuk"`
	Satuan               string  `json:"satuan"`
}

type Read_Detail_Sortir_Response struct {
	Kode_sortir          string                        `json:"kode_sortir"`
	Kode_lot             string                        `json:"kode_lot"`
	Tanggal              string                        `json:"tanggal"`
	Kode_barang          string                        `json:"kdoe_barang"`
	Nama_barang          string                        `json:"nama_barang"`
	Berat_barang         float64                       `json:"berat_barang"`
	Berat_setelah_sortir float64                       `json:"berat_setelah_sortir"`
	Penyusutan           float64                       `json:"penyusutan"`
	Barang_sortir        []Read_Barang_Sortir_Response `gorm:"-"`
}

type Read_Barang_Sortir_Response struct {
	Kode_barang_sortir   string  `json:"kode_barang_sotir"`
	Kode_grade_barang    string  `json:"kode_grade_barang"`
	Nama_grade_barang    string  `json:"nama_grade_barang"`
	Berat_setelah_sortir float64 `json:"berat_setelah_sortir"`
	Kadar_air            float64 `json:"kadar_air"`
	Peresentase          float64 `json:"persentase"`
	Satuan               string  `json:"satuan"`
}
