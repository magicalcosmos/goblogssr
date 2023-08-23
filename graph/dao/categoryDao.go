package dao

import (
	"fmt"
	"time"

	"github.com/magicalcosmos/goblogssr/graph/db"
	"github.com/magicalcosmos/goblogssr/graph/model"
)

// 获取文章目录列表
func GetCategoryList(input model.Q) *model.CategoryWithPage {
	// 获取总数
	var count []int
	err := db.DB.Select(&count, "SELECT COUNT(*) FROM category WHERE status=1")
	if err != nil {
		fmt.Println("GetCategoryList occur error: ", err)
	}

	// 获取具体的值
	categoryWithPage := &model.CategoryWithPage{}
	limitStart := (input.CurrentPage - 1) * input.PageSize
	err = db.DB.Select(&categoryWithPage.Categories, "SELECT * FROM category WHERE status=1 ORDER BY create_time DESC LIMIT ?, ?", limitStart, input.PageSize)
	if err != nil {
		fmt.Println("GetCategoryList occur error: ", err)
	}
	categoryWithPage.Page.CurrentPage = input.CurrentPage
	categoryWithPage.Page.PageSize = input.PageSize
	categoryWithPage.Page.Total = count[0]
	return categoryWithPage
}

// 保存文章目录列表
func SaveCategory(data *model.Category) (category *model.Category) {
	currentTime := time.Now()
	date, _ := time.Parse("2006-01-02 15:04:05", currentTime.Format("2006-01-02 15:04:05"))
	_, err := db.DB.NamedExec(`INSERT INTO category (name,create_time,update_time) VALUES (:name,:createTime,:updateTime)`, map[string]interface{}{
		"name":       data.Name,
		"createTime": date,
		"updateTime": date,
	})
	if err != nil {
		fmt.Println("SaveCategory occur error: ", err)
	}
	return category
}

// 更新文章目录列表
func UpdateCategory(data *model.Category) (category *model.Category) {
	currentTime := time.Now()
	date, _ := time.Parse("2006-01-02 15:04:05", currentTime.Format("2006-01-02 15:04:05"))
	_, err := db.DB.NamedExec(`UPDATE category SET name=:name, update_time=:updateTime WHERE id=:id`, map[string]interface{}{
		"id":         data.ID,
		"name":       data.Name,
		"updateTime": date,
	})
	if err != nil {
		fmt.Println("UpdateCategory occur error: ", err)
	}
	return category
}

// 根据ID删除文章目录
func DeleteCategoryById(data *model.Category) string {
	currentTime := time.Now()
	date, _ := time.Parse("2006-01-02 15:04:05", currentTime.Format("2006-01-02 15:04:05"))
	data.UpdateTime = date
	_, err := db.DB.NamedExec(`UPDATE category SET status=:status, update_time=:updateTime WHERE id=:id`, map[string]interface{}{
		"id":         data.ID,
		"status":     2,
		"updateTime": data.UpdateTime,
	})
	if err != nil {
		fmt.Println("DeleteCategoryById occur error: ", err)
	}
	return "Success"
}
