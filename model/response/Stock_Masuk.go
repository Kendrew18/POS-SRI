package response

type Read_Stock_Masuk_Response struct {
	Kode_stock_masuk     string `json:"kode_stock_masuk"`
	Tanggal              string `json:"tanggal"`
	Kode_lot             string `json:"kode_lot"`
	Kode_quality_control string `json:"kode_quality_control"`
	Nama_supplier        string `json:"nama_supplier"`
}

type Read_Detail_Stock_Masuk_Response struct {
	Kode_stock_masuk     string                             `json:"kode_stock_masuk"`
	Kode_lot             string                             `json:"kode_lot"`
	Kode_quality_control string                             `json:"kode_quality_control"`
	Tanggal              string                             `json:"tanggal"`
	Nama_supplier        string                             `json:"nama_supplier"`
	Detail_barang        []Read_Barang_Stock_Masuk_Response `gorm:"-"`
}

type Read_Barang_Stock_Masuk_Response struct {
	Kode_barang_stock_masuk string  `json:"kode_barang_stock_masuk"`
	Kode_stock_masuk        string  `json:"kode_stock_masuk"`
	Kode_barang             string  `json:"kode_barang"`
	Nama_barang             string  `json:"nama_barang"`
	Kode_grade_barang       string  `json:"kode_grade_barang"`
	Nama_grade_barang       string  `json:"nama_grade_barang"`
	Berat_barang            float64 `json:"berat_barang"`
	Penyusutan              float64 `json:"penyusutan"`
	Kadar_air               float64 `json:"kadar_air"`
	Harga                   int64   `json:"harga"`
	Sub_total               int64   `json:"sub_total"`
}
