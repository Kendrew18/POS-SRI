package request

type Input_Stock_Keluar_Request struct {
	Co                 int     `json:"co"`
	Kode_stock_keluar  string  `json:"kode_stock_keluar"`
	Surat_jalan        string  `json:"surat_jalan"`
	Tanggal            string  `json:"tanggal"`
	Tujuan             string  `json:"tujuan"`
	Tipe               string  `json:"tipe"`
	Total_berat_barang float64 `json:"total_berat_barang"`
	Kode_gudang        string  `json:"kode_gudang"`
}

type Input_Barang_Stock_Keluar_Request struct {
	Kode_barang       string `json:"kode_barang"`
	Kode_grade_barang string `json:"kode_grade_barang"`
	Kode_lot          string `json:"kode_lot"`
	Berat_barang      string `json:"berat_barang"`
}

type Input_Barang_Stock_Keluar_V2_Request struct {
	Co                       int     `json:"co"`
	Kode_barang_stock_keluar string  `json:"kode_barang_stock_keluar"`
	Kode_stock_keluar        string  `json:"kode_stock_keluar"`
	Kode_barang              string  `json:"kode_barang"`
	Kode_grade_barang        string  `json:"kode_grade_barang"`
	Kode_lot                 string  `json:"kode_lot"`
	Berat_barang             float64 `json:"berat_barang"`
}

type Read_Stock_Kelaur_Request struct {
	Kode_gudang string `json:"kode_gudang"`
}

type Read_Stock_Keluar_Filter_Request struct {
	Tujuan        string `json:"tujuan"`
	Tanggal_awal  string `json:"tanggal_awal"`
	Tanggal_akhir string `json:"tanggal_akhir"`
}
