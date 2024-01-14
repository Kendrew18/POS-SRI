package request

type Input_Grade_Barang_Request struct {
	Kode_barang       string `json:"kode_barang"`
	Nama_grade_barang string `json:"nama_grade_barang"`
}

type Input_Grade_Barang_Request_V2 struct {
	Co                int    `json:"co"`
	Kode_barang       string `json:"kode_barang"`
	Kode_grade_barang string `json:"kode_grade_barang"`
	Nama_grade_barang string `json:"nama_grade_barang"`
}
