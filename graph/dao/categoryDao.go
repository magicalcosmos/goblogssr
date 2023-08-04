package dao

import (
	"fmt"

	"github.com/magicalcosmos/goblogssr/graph/db"
	"github.com/magicalcosmos/goblogssr/graph/model"
)

// GetUserInfo get user info
func GetCategoryList() (category []*model.Category) {
	err := db.DB.Select(&category, "SELECT * FROM category")
	if err != nil {
		fmt.Println("GetCategoryList occur error: ", err)
	}
	return category
}

// GetUserInfo get user info
func SaveCategory() (category *model.Category) {
	err := db.DB.Select(&category, "SELECT * FROM category")
	if err != nil {
		fmt.Println("SaveCategory occur error: ", err)
	}
	_, err = db.DB.NamedExec(`INSERT INTO catetory (first_name,last_name,email) VALUES (:first,:last,:email)`, category)
	if err != nil {
		fmt.Println("SaveCategory occur error: ", err)
	}
	return category
}
