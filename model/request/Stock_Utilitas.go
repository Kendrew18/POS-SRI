package request

type Input_Stock_Utilitas_Request struct {
	Co                  int    `json:"co"`
	Kode_stock_utilitas string `json:"kode_stock_utilitas"`
	Nama_stock_utilitas string `json:"nama_stock_utilitas"`
	Tanggal             string `json:"tanggal"`
	Jumlah              int    `json:"jumlah"`
	Kode_gudang         string `json:"kode_gudang"`
}

type Read_Stock_Utilitas_Request struct {
	Kode_gudang string `json:"kode_gudang"`
}

type Update_Stock_Utilitas_Request struct {
	Nama_stock_utilitas string `json:"nama_stock_utilitas"`
	Tanggal             string `json:"tanggal"`
	Jumlah              int    `json:"jumlah"`
}

type Update_Stock_Utilitas_Kode_Request struct {
	Kode_stock_utilitas string `json:"kode_stock_utilitas"`
}
