package request

type Input_Quality_Control_Request struct {
	Co                   int    `json:"co"`
	Kode_quality_control string `json:"kode_quality_control"`
	Kode_lot             string `json:"kode_lot"`
	Kode_pre_order       string `json:"kode_pre_order"`
	Tanggal_masuk        string `json:"tanggal_masuk"`
	Nama_supplier        string `json:"nama_supplier"`
	Kode_gudang          string `json:"kode_gudang"`
}

type Input_Barang_Quality_Control_Request struct {
	Co                          int     `json:"co"`
	Kode_barang_quality_control string  `json:"kode_barang_quality_control"`
	Kode_quality_control        string  `json:"kode_quality_control"`
	Kode_barang                 string  `json:"kode_barang"`
	Kode_grade_barang           string  `json:"kode_grade_barang"`
	Berat_barang                float64 `json:"berat_barang"`
	Harga                       int64   `json:"harga"`
	Sub_total                   int64   `json:"sub_total"`
}

type Read_quality_control_Request struct {
	Kode_gudang string `json:"kode_gudang"`
}

type Read_Quality_Control_Filter_Request struct {
	Kode_lot      string `json:"kode_lot"`
	Tanggal_awal  string `json:"tanggal_awal"`
	Tanggal_akhir string `json:"tanggal_akhir"`
}

type Update_Berat_Barang_Rill_Request struct {
	Berat_barang_rill    float64 `json:"berat_barang_rill"`
	Berat_barang_ditolak float64 `json:"berat_barang_ditolak"`
	Penyusutan           float64 `json:"penyusutan"`
	Persentase           float64 `json:"persentase"`
	Kadar_air            float64 `json:"kadar_air"`
}

type Update_Quality_Control_Kode_Request struct {
	Kode_barang_quality_control string `json:"kode_barang_quality_control"`
}

type Read_Detail_Quality_Control_Request struct {
	Kode_quality_control string `json:"kode_quality_control"`
}

type Update_Status_Quality_Control_Request struct {
	Status int `json:"status"`
}

type Kode_Quality_Control_Request struct {
	Kode_quality_control string `json:"kode_quality_control"`
}
