package dao

import (
	"fmt"

	"github.com/magicalcosmos/goblogssr/graph/db"
	"github.com/magicalcosmos/goblogssr/graph/model"
)

// GetUserInfo get user info
func GetArticleList() (article []*model.Article) {
	err := db.DB.Select(&article, "SELECT * FROM article")
	if err != nil {
		fmt.Println("GetArticleList occur error: ", err)
	}
	return article
}

// GetUserInfo get user info
func SaveArticle() (article *model.Article) {
	err := db.DB.Select(&article, "SELECT * FROM category")
	if err != nil {
		fmt.Println("SaveArticle occur error: ", err)
	}
	return article
}
