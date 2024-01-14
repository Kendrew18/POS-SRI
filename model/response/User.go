package response

type User_Response struct {
	Kode_user   string `json:"kode_user"`
	Status      int    `json:"status"`
	Kode_gudang string `json:"kode_gudang"`
}
