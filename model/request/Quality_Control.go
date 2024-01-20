package request

type Input_Quality_Control_Request struct {
	Co                   int     `json:"co"`
	Kode_quality_control string  `json:"kode_quality_control"`
	Kode_lot             string  `json:"kode_lot"`
	Tanggal_masuk        string  `json:"tanggal_masuk"`
	Kode_pre_order       string  `json:"kode_pre_order"`
	Kode_supplier        string  `json:"kode_supplier"`
	Kode_barang          string  `json:"kode_barang"`
	Berat_barang         float64 `json:"berat_barang"`
	Kode_gudang          string  `json:"kode_gudang"`
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
	Berat_barang_rill float64 `json:"berat_barang_rill"`
}

type Update_Quality_Control_Kode_Request struct {
	Kode_quality_control string `json:"kode_quality_control"`
}
