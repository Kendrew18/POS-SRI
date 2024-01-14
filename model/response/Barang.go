package response

type Read_Jenis_Barang_Response struct {
	Kode_barang  string                       `json:"kode_barang"`
	Nama_barang  string                       `json:"nama_barang"`
	Grade_barang []Read_Grade_Barang_Response `json:"grade_barang"`
}
