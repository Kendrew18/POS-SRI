package response

type Read_Pre_Order_Response struct {
	Kode_pre_order string  `json:"kode_pre_order"`
	Kode_lot       string  `json:"kode_lot"`
	Tanggal        string  `json:"tanggal"`
	Kode_supplier  string  `json:"kode_supplier"`
	Nama_supplier  string  `json:"nama_supplier"`
	Tanggal_etd    string  `json:"tanggal_etd"`
	Tanggal_rtd    string  `json:"tanggal_rtd"`
	Total_berat    float64 `json:"total_berat"`
	Kode_satuan    string  `json:"kode_satuan"`
	Nama_satuan    string  `json:"nama_satuan"`
	Kode_gudang    string  `json:"kode_gudang"`
	Nama_gudang    string  `json:"nama_gudang"`
}

type Detail_Pre_Order_Response struct {
	Kode_pre_order          string                           `json:"kode_pre_order"`
	Kode_lot                string                           `json:"kode_lot"`
	Tanggal_pre_order       string                           `json:"tanggal_pre_order"`
	Kode_supplier           string                           `json:"kode_supplier"`
	Nama_supplier           string                           `json:"nama_supplier"`
	Total_berat             float64                          `json:"total_berat"`
	Kode_jenis_pembayaran   string                           `json:"kode_jenis_pembayaran"`
	Nama_jenis_pembayaran   string                           `json:"nama_jenis_pembayaran"`
	Status_pembayaran       string                           `json:"status_pembayaran"`
	Detail_barang_pre_order []Read_Barang_Pre_Order_Response `json:"detail_barang_pre_order"`
}

type Read_Barang_Pre_Order_Response struct {
	Kode_barang_pre_order string  `json:"kode_barang_pre_order"`
	Kode_barang           string  `json:"kode_barang"`
	Nama_barang           string  `json:"nama_barang"`
	Berat_barang          float64 `json:"berat_barang"`
	Harga                 int64   `json:"harga"`
}
