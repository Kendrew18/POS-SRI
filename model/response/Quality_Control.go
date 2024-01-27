package response

type Read_Quality_Control_Response struct {
	Kode_quality_control string `json:"kode_quality_control"`
	Kode_lot             string `json:"kode_lot"`
	Tanggal_masuk        string `json:"tanggal_masuk"`
	Kode_pre_order       string `json:"kode_pre_order"`
	Nama_supplier        string `json:"nama_supplier"`
	Status               int    `json:"status"`
}

type Read_Detail_Quality_Control_Response struct {
	Kode_quality_control   string                                 `json:"kode_quality_control"`
	Kode_lot               string                                 `json:"kode_lot"`
	Tanggal_masuk          string                                 `json:"tanggal_masuk"`
	Kode_pre_order         string                                 `json:"kode_pre_order"`
	Nama_supplier          string                                 `json:"nama_supplier"`
	Barang_quality_control []Read_Barang_Quality_Control_Response `gorm:"-"`
}

type Read_Barang_Quality_Control_Response struct {
	Kode_barang_quality_control string  `json:"kode_barang_quality_control"`
	Kode_quality_control        string  `json:"kode_quality_control"`
	Kode_barang                 string  `json:"kode_barang"`
	Nama_barang                 string  `json:"nama_barang"`
	Kode_grade_barang           string  `json:"kode_grade_barang"`
	Nama_grade_barang           string  `json:"nama_grade_barang"`
	Berat_barang                float64 `json:"berat_barang"`
	Berat_barang_rill           float64 `json:"berat_barang_rill"`
	Berat_barang_ditolak        float64 `json:"berat_barang_ditolak"`
	Penyusutan                  float64 `json:"penyusutan"`
	Persentase                  float64 `json:"persentase"`
	Kadar_air                   float64 `json:"kadar_air"`
	Harga                       int64   `json:"harga"`
	Sub_total                   int64   `json:"sub_total"`
}
