package response

type Read_Stock_Keluar_Response struct {
	Kode_stock_keluar  string  `json:"kode_stock_keluar"`
	Surat_jalan        string  `json:"surat_jalan"`
	Tanggal            string  `json:"tanggal"`
	Total_berat_barang float64 `json:"total_berat_barang"`
	Satuan             string  `json:"satuan"`
	Tujuan             string  `json:"tujuan"`
	Tipe               string  `json:"tipe"`
}

type Read_Detail_Stock_Keluar_Response struct {
}
