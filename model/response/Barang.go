package response

type Read_Jenis_Barang_Response struct {
	Kode_barang  string                       `json:"kode_barang"`
	Nama_barang  string                       `json:"nama_barang"`
	Grade_barang []Read_Grade_Barang_Response `json:"grade_barang"`
}

type Read_Barang_Stock_Response struct {
	Kode_barang string  `json:"kode_barang"`
	Nama_barang string  `json:"nama_barang"`
	Total_berat float64 `json:"Total_berat"`
	Nama_satuan string  `json:"nama_satuan"`
}

type Barang struct {
	Kode_barang        string `json:"kode_barang"`
	Nama_barang        string `json:"nama_barang"`
	Nama_satuan_barang string `json:"nama_satuan_barang"`
}
