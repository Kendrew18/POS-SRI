package request

type Input_Jenis_Barang_Request struct {
	Co          int    `json:"co"`
	Kode_barang string `json:"kode_barang"`
	Nama_barang string `json:"nama_barang"`
	Kode_gudang string `json:"kode_gudang"`
}

type Read_Jenis_Barang_Response struct {
	Kode_gudang string `json:"kode_gudang"`
}
