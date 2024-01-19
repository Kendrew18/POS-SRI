package request

type Input_Jenis_Pembayaran_Request struct {
	Co                    int    `json:"co"`
	Kode_jenis_pembayaran string `json:"kode_jenis_pembayaran"`
	Nama_jenis_pembayaran string `json:"nama_jenis_pembayaran"`
	Kode_gudang           string `json:"kode_gudang"`
}

type Read_Jenis_Pembayaran_Request struct {
	Kode_gudang string `json:"kode_gudang"`
}
