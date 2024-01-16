package request

type Input_List_Formula_Request struct {
	Co                int    `json:"co"`
	Kode_list_formula string `json:"kode_list_formula"`
	Nama_barang_jadi  string `json:"nama_barang_jadi"`
	Berat_estimasi    string `json:"berat_estimasi"`
	Kode_gudang       string `json:"kode_gudang"`
}

type Input_Barang_List_Formula_Response struct {
	Kode_barang string `json:"kode_barang"`
}
