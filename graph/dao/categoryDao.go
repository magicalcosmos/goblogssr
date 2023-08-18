package dao

import (
	"fmt"
	"time"

	"github.com/magicalcosmos/goblogssr/graph/db"
	"github.com/magicalcosmos/goblogssr/graph/model"
)

// 获取文章目录列表
func GetCategoryList() (category []*model.Category) {
	err := db.DB.Select(&category, "SELECT * FROM category where status=1 order by create_time desc")
	if err != nil {
		fmt.Println("GetCategoryList occur error: ", err)
	}
	return category
}

// 保存文章目录列表
func SaveCategory(ca *model.Category) (category *model.Category) {
	currentTime := time.Now()
	date, _ := time.Parse("2006-01-02 15:04:05", currentTime.Format("2006-01-02 15:04:05"))
	ca.CreateTime = date
	ca.UpdateTime = date
	_, err := db.DB.NamedExec(`INSERT INTO category (name,create_time,update_time) VALUES (:name,:createTime,:updateTime)`, map[string]interface{}{
		"name":       ca.Name,
		"createTime": ca.CreateTime,
		"updateTime": ca.UpdateTime,
	})
	if err != nil {
		fmt.Println("SaveCategory occur error: ", err)
	}
	return category
}

// 更新文章目录列表
func UpdateCategory(ca *model.Category) (category *model.Category) {
	currentTime := time.Now()
	date, _ := time.Parse("2006-01-02 15:04:05", currentTime.Format("2006-01-02 15:04:05"))
	ca.UpdateTime = date
	_, err := db.DB.NamedExec(`UPDATE category SET name=:name, update_time=:updateTime WHERE id=:id`, map[string]interface{}{
		"id":         ca.ID,
		"name":       ca.Name,
		"updateTime": ca.UpdateTime,
	})
	if err != nil {
		fmt.Println("SaveCategory occur error: ", err)
	}
	return category
}

// 根据ID删除文章目录
func DeleteCategoryById(ca *model.Category) string {
	currentTime := time.Now()
	date, _ := time.Parse("2006-01-02 15:04:05", currentTime.Format("2006-01-02 15:04:05"))
	ca.UpdateTime = date
	_, err := db.DB.NamedExec(`UPDATE category SET status=:status, update_time=:updateTime WHERE id=:id`, map[string]interface{}{
		"id":         ca.ID,
		"status":     2,
		"updateTime": ca.UpdateTime,
	})
	if err != nil {
		fmt.Println("SaveCategory occur error: ", err)
	}
	return "Success"
}
