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

type Input_Detail_Stock_Request struct {
	Co                      int     `json:"co"`
	Kode_barang_stock_masuk string  `json:"kode_barang_stock_masuk"`
	Kode_lot                string  `json:"kode_lot"`
	Kode_stock_masuk        string  `json:"kode_stock_masuk"`
	Kode_barang             string  `json:"kode_barang"`
	Kode_grade_barang       string  `json:"kode_grade_barang"`
	Berat_barang            float64 `json:"berat_barang"`
	Kadar_air               float64 `json:"kadar_air"`
	Penyusutan              float64 `json:"penyusutan"`
}
