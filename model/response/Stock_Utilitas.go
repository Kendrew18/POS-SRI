package response

type Stock_Utilitas_Response struct {
	Kode_stock_utilitas string `json:"kode_stock_utilitas"`
	Tanggal             string `json:"tanggal"`
	Nama_stock_utilitas string `json:"nama_stock_utilitas"`
	Jumlah              int    `json:"jumlah"`
}
