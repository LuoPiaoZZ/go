package model

//操作的对应表的model
type User struct {
	ID int `json:"id" form:"id"`
	Uname string `json:"uname" form:"uname"`
	Upassword string `json:"upassword" form:"upassword"`
}

