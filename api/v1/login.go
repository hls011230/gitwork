package v1

import (
	"A11Smile/db"
	"A11Smile/db/model"
)

func Login(user *model.LoginUser) int {
	cli := db.Get()
	var count int
	cli.Raw("select count(1) from users where email= ? and passwd = ?", user.Email, user.Passwd).Find(&count)
	return count
}
