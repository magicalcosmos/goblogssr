package dao

import (
	"errors"
	"fmt"

	"github.com/magicalcosmos/goblogssr/graph/db"
	"github.com/magicalcosmos/goblogssr/graph/model"
)

// CreateCategory create category
func CreateCategory(input model.NewCategory) (*model.Category, error) {
	data := map[string]interface{}{
		"pid":    0,
		"sort":   1,
		"status": 1,
		"name":   input.Name,
	}
	category := model.Category{}
	err := db.DB.Get(&category, `SELECT * FROM category WHERE name=?`, input.Name)
	if category.Name != "" {
		err = errors.New(fmt.Sprintf("Category name [%s] exist", input.Name))
		return nil, err
	}

	tx, err := db.DB.NamedExec("INSERT INTO category (name, pid, sort, status) VALUES (:name, :pid, :sort, :status)", data)
	if err != nil {
		return nil, err
	}
	id, _ := tx.LastInsertId()
	result := &model.Category{
		ID:     id,
		Name:   input.Name,
		Pid:    0,
		Sort:   1,
		Status: 1,
	}
	return result, nil
}

// GetCategory get all category
func GetCategory() (category []*model.Category, err error) {
	err = db.DB.Select(&category, `SELECT * FROM category`)
	return
}

// DeleteCategory delete  category by id
func DeleteCategory(input model.NewCategory) (message string, err error) {
	 result, err := db.DB.Exec(`DELETE FROM category WHERE id=?`, input.ID)
	 if result != nil {
	 	message = "Delete success"
	 }
	return
}
