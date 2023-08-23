package dao

import (
	"fmt"
	"time"

	"github.com/magicalcosmos/goblogssr/graph/db"
	"github.com/magicalcosmos/goblogssr/graph/model"
)

// 获取用户列表
func GetUserList(input model.Q) *model.UserWithPage {
	// 获取总数
	var count []int
	err := db.DB.Select(&count, "SELECT COUNT(*) FROM User WHERE status=1")
	if err != nil {
		fmt.Println("GetUserList occur error: ", err)
	}

	// 获取具体的值
	userWithPage := &model.UserWithPage{}
	limitStart := (input.CurrentPage - 1) * input.PageSize
	err = db.DB.Select(&userWithPage.Users, "SELECT * FROM user WHERE status=1 ORDER BY create_time DESC LIMIT ?, ?", limitStart, input.PageSize)
	if err != nil {
		fmt.Println("GetUserList occur error: ", err)
	}
	userWithPage.Page.CurrentPage = input.CurrentPage
	userWithPage.Page.PageSize = input.PageSize
	userWithPage.Page.Total = count[0]
	return userWithPage
}

// 保存用户
func SaveUser(data *model.User) (user *model.User) {
	currentTime := time.Now()
	date, _ := time.Parse("2006-01-02 15:04:05", currentTime.Format("2006-01-02 15:04:05"))
	_, err := db.DB.NamedExec(`INSERT INTO user (username,email,password,create_time,update_time) VALUES (:username,:email,:password,:createTime,:updateTime)`, map[string]interface{}{
		"username":   data.Username,
		"password":   data.Password,
		"email":      data.Email,
		"createTime": date,
		"updateTime": date,
	})
	if err != nil {
		fmt.Println("SaveUser occur error: ", err)
	}
	return user
}

// 更新用户
func UpdateUser(data *model.User) (user *model.User) {
	currentTime := time.Now()
	date, _ := time.Parse("2006-01-02 15:04:05", currentTime.Format("2006-01-02 15:04:05"))
	_, err := db.DB.NamedExec(`UPDATE user SET email=:email, password=:password update_time=:updateTime WHERE id=:id`, map[string]interface{}{
		"id":         data.ID,
		"email":      data.Email,
		"password":   data.Password,
		"updateTime": date,
	})
	if err != nil {
		fmt.Println("UpdateUser occur error: ", err)
	}
	return user
}

// 根据ID删除文章目录
func DeleteUserById(data *model.User) string {
	currentTime := time.Now()
	date, _ := time.Parse("2006-01-02 15:04:05", currentTime.Format("2006-01-02 15:04:05"))
	data.UpdateTime = date
	_, err := db.DB.NamedExec(`UPDATE user SET status=:status, update_time=:updateTime WHERE id=:id`, map[string]interface{}{
		"id":         data.ID,
		"status":     2,
		"updateTime": data.UpdateTime,
	})
	if err != nil {
		fmt.Println("DeleteUserById occur error: ", err)
	}
	return "Success"
}
