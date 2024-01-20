package request

type Input_Pre_Order_Request struct {
	Co                    int     `json:"co"`
	Kode_pre_order        string  `json:"kode_pre_order"`
	Kode_lot              string  `json:"kode_lot"`
	Tanggal               string  `json:"tanggal"`
	Kode_supplier         string  `json:"kode_supplier"`
	Tanggal_ETD           string  `json:"tanggal_etd"`
	Tanggal_RTD           string  `json:"tanggal_rtd"`
	Total_berat           float64 `json:"total_berat"`
	Total_harga           int64   `json:"total_harga"`
	Kode_jenis_pembayaran string  `json:"kode_jenis_pembayaran"`
	Kode_gudang           string  `json:"kode_gudang"`
	Status                int     `json:"status"`
}

type Input_Barang_Pre_Order_Request struct {
	Kode_barang  string `json:"kode_barang"`
	Berat_barang string `json:"berat_barang"`
	Harga        string `json:"harga"`
}

type Input_Barang_Pre_Order_V2_Request struct {
	Co                    int     `json:"co"`
	Kode_pre_order        string  `json:"kode_pre_order"`
	Kode_barang_pre_order string  `json:"kode_barang_pre_order"`
	Kode_barang           string  `json:"kode_barang"`
	Berat_barang          float64 `json:"berat_barang"`
	Harga                 int64   `json:"harga"`
}

type Read_Pre_Order_Request struct {
	Kode_gudang string `json:"kode_gudang"`
}

type Read_Pre_Order_Filter_Request struct {
	Kode_supplier string `json:"kode_supplier"`
	Tanggal_awal  string `json:"tanggal_awal"`
	Tanggal_akhir string `json:"tanggal_akhir"`
}

type Read_Detail_Pre_Order_Request struct {
	Kode_pre_order string `json:"kode_pre_order"`
}

type Update_Pre_order_Request struct {
	Kode_barang  string  `json:"kode_barang"`
	Berat_barang float64 `json:"berat_barang"`
	Harga        int64   `json:"harga"`
}

type Update_Pre_Order_Kode_Request struct {
	Kode_barang_pre_order string `json:"kode_barang_pre_order"`
}

type Update_Status_Pre_Order_Request struct {
	Status int `json:"status"`
}

type Kode_Pre_Order_Request struct {
	Kode_pre_order string `json:"kode_pre_order"`
}
