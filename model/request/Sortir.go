package request

type Input_Sortir_Request struct {
	Co                     int     `json:"co"`
	Kode_sortir            string  `json:"kode_sortir"`
	Kode_stock_masuk       string  `json:"kode_stock_masuk"`
	Kode_lot               string  `json:"kode_lot"`
	Tanggal                string  `json:"tanggal"`
	Kode_barang            string  `json:"kode_barang"`
	Kode_grade_barang      string  `json:"kode_grade_barang"`
	Berat_barang           float64 `json:"berat_barang"`
	Berat_setelah_disortir float64 `json:"berat_setelah_disortir"`
	Penyusutan             float64 `json:"penyusutan"`
	Kode_gudang            string  `json:"kode_gudang"`
}

type Input_Barang_Sortir_Request struct {
	Co                   int     `json:"co"`
	Kode_barang_sortir   string  `json:"kode_barang_sotir"`
	Kode_sortir          string  `json:"kode_sortir"`
	Kode_grade_barang    string  `json:"kode_grade_barang"`
	Berat_setelah_sortir float64 `json:"berat_setelah_sortir"`
	Kadar_air            float64 `json:"kadar_air"`
	Persentase           float64 `json:"persentase"`
}

type Update_Status_Sortir_Request struct {
	Status_sortir int `json:"status_sortir"`
}

type Read_Sortir_Request struct {
	Kode_gudang string `json:"kode_gudang"`
}

type Read_Sortir_Filter_Request struct {
	Kode_lot      string `json:"kode_lot"`
	Tanggal_awal  string `json:"tanggal_awal"`
	Tanggal_akhir string `json:"tanggal_akhir"`
}

type Read_Detail_Sortir_Request struct {
	Kode_sortir string `json:"kode_sortir"`
}
