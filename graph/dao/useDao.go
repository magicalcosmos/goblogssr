package dao

import (
	"fmt"

	"github.com/magicalcosmos/goblogssr/graph/db"
	"github.com/magicalcosmos/goblogssr/graph/model"
)

// GetUserInfo get user info
func GetUserInfo() (user []*model.User) {
	db.DB.Select(&user, "SELECT * FROM user")
	fmt.Println(user[0].Password)
	return
}
