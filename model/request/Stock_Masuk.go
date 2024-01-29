package request

type Input_Stock_Masuk_Request struct {
	Co                   int    `json:"co"`
	Kode_stock_masuk     string `json:"kode_stock_masuk"`
	Kode_quality_control string `json:"kode_quality_control"`
	Tanggal              string `json:"tanggal"`
	Kode_lot             string `json:"kode_lot"`
	Nama_supplier        string `json:"nama_supplier"`
	Kode_gudang          string `json:"kode_gudang"`
}

type Input_Barang_Stock_Masuk_Request struct {
	Co                      int     `json:"co"`
	Kode_barang_stock_masuk string  `json:"kode_barang_stock_masuk"`
	Kode_stock_masuk        string  `json:"kode_stock_masuk"`
	Kode_lot                string  `json:"kode_lot"`
	Kode_barang             string  `json:"kode_barang"`
	Kode_grade_barang       string  `json:"kode_grade_barang"`
	Berat_barang            float64 `json:"berat_barang"`
	Kadar_air               float64 `json:"kadar_air"`
	Penyusutan              float64 `json:"penyusutan"`
	Harga                   int64   `json:"harga"`
	Sub_total               int64   `json:"sub_total"`
}

type Read_Stock_Masuk_Request struct {
	Kode_gudang string `json:"kode_gudang"`
}

type Read_Stock_Masuk_Filter_Request struct {
	Kode_lot      string `json:"kode_lot"`
	Tanggal_awal  string `json:"tanggal_awal"`
	Tanggal_akhir string `json:"tanggal_akhir"`
}

type Read_Detail_Stock_Masuk_Request struct {
	Kode_stock_masuk string `json:"kode_stock_masuk"`
}

type Input_Stock_Keluar_Masuk_Request struct {
	Co      int    `json:"co"`
	Kode    string `json:"kode"`
	Tanggal string `json:"tanggal"`
}

type Input_Barang_Stock_Keluar_Masuk_Request struct {
	Co                       int     `json:"co"`
	Kode_barang_keluar_masuk string  `json:"kode_barang_keluar_masuk"`
	Kode                     string  `json:"kode"`
	Kode_lot                 string  `json:"kode_lot"`
	Kode_barang              string  `json:"kode_barang"`
	Kode_grade_barang        string  `json:"kode_grade_barang"`
	Berat_barang             float64 `json:"berat_barang"`
	Kadar_air                float64 `json:"kadar_air"`
	Penyusutan               float64 `json:"penyusutan"`
	Keterangan               string  `json:"keterangan"`
}

type Kode_Stock_Masuk_Request struct {
	Kode_stock_masuk string `json:"kode_stock_masuk"`
}
