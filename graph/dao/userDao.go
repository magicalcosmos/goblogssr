package dao

import (
	"fmt"

	"github.com/magicalcosmos/goblogssr/graph/db"
	"github.com/magicalcosmos/goblogssr/graph/model"
)

// GetUserInfo get user info
func GetUserInfo() (user []*model.User) {
	err := db.DB.Select(&user, "SELECT * FROM user")
	if err != nil {
		fmt.Println("GetUserInfo occur error: ", err)
	}
	return user
}
