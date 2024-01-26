package response

type Read_Pre_Order_Response struct {
	Kode_pre_order string  `json:"kode_pre_order"`
	Kode_lot       string  `json:"kode_lot"`
	Tanggal        string  `json:"tanggal"`
	Nama_supplier  string  `json:"nama_supplier"`
	Tanggal_etd    string  `json:"tanggal_etd"`
	Tanggal_rtd    string  `json:"tanggal_rtd"`
	Total_berat    float64 `json:"total_berat"`
	Total_harga    int64   `json:"total_harga"`
	Kode_gudang    string  `json:"kode_gudang"`
	Nama_gudang    string  `json:"nama_gudang"`
}

type Detail_Pre_Order_Response struct {
	Kode_pre_order          string                           `json:"kode_pre_order"`
	Kode_lot                string                           `json:"kode_lot"`
	Tanggal                 string                           `json:"tanggal"`
	Nama_supplier           string                           `json:"nama_supplier"`
	Total_berat             float64                          `json:"total_berat"`
	Status                  int                              `json:"status"`
	Detail_barang_pre_order []Read_Barang_Pre_Order_Response `json:"detail_barang_pre_order"`
}

type Read_Barang_Pre_Order_Response struct {
	Kode_barang_pre_order string  `json:"kode_barang_pre_order"`
	Kode_barang           string  `json:"kode_barang"`
	Nama_barang           string  `json:"nama_barang"`
	Kode_grade_barang     string  `json:"kode_grade_barang"`
	Nama_grade_barang     string  `json:"nama_grade_barang"`
	Satuan                string  `json:"satuan"`
	Berat_barang          float64 `json:"berat_barang"`
	Harga                 int64   `json:"harga"`
	Sub_total             int64   `json:"sub_total"`
}
