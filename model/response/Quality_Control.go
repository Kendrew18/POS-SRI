package response

type Read_Quality_Control_Response struct {
	Kode_quality_control string  `json:"kode_quality_control"`
	Kode_lot             string  `json:"kode_lot"`
	Tanggal_masuk        string  `json:"tanggal_masuk"`
	Kode_pre_order       string  `json:"kode_pre_order"`
	Kode_supplier        string  `json:"kode_supplier"`
	Nama_supplier        string  `json:"nama_supplier"`
	Kode_barang          string  `json:"kode_barang"`
	Nama_barang          string  `json:"nama_barang"`
	Berat_barang         float64 `json:"berat_barang"`
	Berat_barang_rill    float64 `json:"berat_barang_rill"`
	Nama_satuan          string  `json:"nama_satuan"`
	Status               int     `json:"status"`
}
