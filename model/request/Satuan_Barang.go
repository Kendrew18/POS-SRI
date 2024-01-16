package request

type Input_Satuan_Barang_Request struct {
	Co                 int    `json:"co"`
	Kode_satuan_barang string `json:"kode_satuan_barang"`
	Nama_satuan_barang string `json:"nama_satuan_barang"`
	Kode_gudang        string `json:"kode_gudang"`
}

type Read_Satuan_Barang_Request struct {
	Kode_gudang string `json:"kode_gudang"`
}
