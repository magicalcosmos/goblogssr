package dao

import (
	"fmt"
	"strconv"
	"time"

	"github.com/magicalcosmos/goblogssr/graph/db"
	"github.com/magicalcosmos/goblogssr/graph/model"
)

// Get article list get user info
func GetArticleList(input model.Q) *model.ArticleWithPage {
	// 获取总数
	var count []int
	err := db.DB.Select(&count, "SELECT COUNT(*) FROM article WHERE status=1")
	if err != nil {
		fmt.Println("GetArticleList occur error: ", err)
	}

	// 获取具体的值
	limitStart := (input.CurrentPage - 1) * input.PageSize
	articleWithPage := &model.ArticleWithPage{}
	err = db.DB.Select(&articleWithPage.Articles, `
		SELECT 
			* 
		FROM 
			article 
		WHERE 
			status=1 
		ORDER BY 
			create_time 
		DESC LIMIT ?, ?
	`, limitStart, input.PageSize)
	if err != nil {
		fmt.Println("GetArticleList occur error: ", err)
	}
	articleWithPage.Page.CurrentPage = input.CurrentPage
	articleWithPage.Page.PageSize = input.PageSize
	articleWithPage.Page.Total = count[0]
	return articleWithPage
}

// Save article get user info
// TODO: userId 需要根据登录用户来传
func SaveArticle(data *model.Article) (article *model.Article) {
	currentTime := time.Now()
	date, _ := time.Parse("2006-01-02 15:04:05", currentTime.Format("2006-01-02 15:04:05"))
	var categoryData []model.Category
	err := db.DB.Select(&categoryData, "SELECT * FROM category WHERE status=1 ORDER BY create_time DESC LIMIT 1")
	if err != nil {
		fmt.Println("GetArticleList occur error: ", err)
	}
	rows, err := db.DB.NamedExec(`
		INSERT INTO 
			article (
				title,
				create_time,
				update_time, 
				publish_time,
				user_id, 
				category_id, 
				brief,
				content, 
				published
			) VALUES (
				:title,
				:createTime,
				:updateTime, 
				:publishTime,
				:userId, 
				:categoryId, 
				:brief,
				:content, 
				:published
			)`,
		map[string]interface{}{
			"title":       data.Title,
			"createTime":  date,
			"updateTime":  date,
			"publishTime": date,
			"brief":       "",
			"content":     "",
			"published":   0,
			"userId":      1,
			"categoryId":  categoryData[0].ID,
		})

	if err != nil {
		fmt.Println("SaveArticle occur error: ", err)
	}
	insertId, _ := rows.LastInsertId()

	s := strconv.FormatInt(insertId, 10)

	id, _ := strconv.Atoi(s)

	return GetArticleById(&model.Article{
		ID: id,
	})
}

// Update article get user info
func UpdateArticle(data *model.Article) (article *model.Article) {
	currentTime := time.Now()
	date, _ := time.Parse("2006-01-02 15:04:05", currentTime.Format("2006-01-02 15:04:05"))
	publishAt, _ := time.Parse("2006-01-02 15:04:05", data.PublishAt.Format("2006-01-02 15:04:05"))
	_, err := db.DB.NamedExec(`
		UPDATE
			article 
		SET 
		  title=:title, 
			brief=:brief, 
			published=:published, 
			user_id=:userId, 
			category_id=:categoryId, 
			content=:content, 
			update_time=:updateAt,
			publish_time=:publishAt
		WHERE 
			id=:id
	`, map[string]interface{}{
		"id":         data.ID,
		"title":      data.Title,
		"brief":      data.Brief,
		"content":    data.Content,
		"published":  data.Published,
		"userId":     1,
		"publishAt":  publishAt,
		"categoryId": data.CategoryId,
		"updateAt":   date,
	})
	if err != nil {
		fmt.Println("UpdateArticle occur error: ", err)
	}
	return article
}

// delete article by id
func DeleteArticleById(data *model.Article) string {
	currentTime := time.Now()
	date, _ := time.Parse("2006-01-02 15:04:05", currentTime.Format("2006-01-02 15:04:05"))
	_, err := db.DB.NamedExec(`
		UPDATE 
			article 
		SET 
			status=:status, 
			update_time=:updateTime 
		WHERE 
			id=:id
	`, map[string]interface{}{
		"id":         data.ID,
		"status":     2,
		"updateTime": date,
	})
	if err != nil {
		fmt.Println("DeleteArticleById occur error: ", err)
	}
	return "Success"
}

// delete article by id
func GetArticleById(data *model.Article) *model.Article {
	articleData := model.Article{}
	err := db.DB.Get(&articleData, "SELECT * FROM article WHERE id=?", data.ID)
	if err != nil {
		fmt.Println("GetArticleById occur error: ", err)
	}
	return &articleData
}
