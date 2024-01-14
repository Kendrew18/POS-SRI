package user

import (
	"POS-SRI/db"
	"POS-SRI/model/request"
	"POS-SRI/model/response"
	"fmt"
	"net/http"
)

func Login(user request.User_Request) (response.Response, error) {

	var res response.Response
	var us response.User_Response
	con := db.CreateConGorm().Table("user")

	err := con.Select("kode_user", "status", "kode_gudang").Where("username =? AND password =?", user.Username, user.Password).Scan(&us).Error

	fmt.Println(err)

	if err != nil || us.Kode_user == "" {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		us.Kode_user = ""
		res.Data = us

	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = us
	}

	fmt.Println()

	return res, nil

}
