package define

import "errors"

//用户数据模型
type User struct {
	Id        int    `json:"id" gorm:"id" form:"id"`
	Uname     string `json:"uname" gorm:"uname" form:"uname" binding:"required"`
	Upassword string `json:"upassword" gorm:"upassword" form:"upassword" binding:"required"`
}

func (c *User) Validate() error {
	if len(c.Uname) > 20 || len(c.Uname) < 1 {
		return errors.New("uname不符合规范")
	}
	if len(c.Upassword) > 20 || len(c.Upassword) < 1 {
		return errors.New("upassword不符合规范")
	}
	return nil
}
