package model

import "time"

// User User
type Article struct {
	ID         string    `json:"id" db:"id"`
	AuthorId   string    `json:"authorId" db:"author_id"`
	ParentId   string    `json:"parentId" db:"parent_id"`
	Title      string    `json:"title" db:"title"`
	Brief      string    `json:"brief" db:"brief"`
	Content    string    `json:"content" db:"content"`
	Published  int       `json:"published" db:"published"`
	Tag        string    `json:"tag" db:"tag"`
	Pv         int       `json:"pv" db:"pv"`
	Review     int       `json:"review" db:"review"`
	Recommend  int       `json:"recommend" db:"recommend"`
	Like       int       `json:"like" db:"like"`
	Status     int       `json:"status" db:"status"`
	CreateTime time.Time `json:"createAt" db:"create_time"`
	UpdateTime time.Time `json:"updateAt" db:"update_time"`
}

type ArticleCategory struct {
	ArticleId  string `json:"articleId" db:"article_id"`
	CategoryId string `json:"categoryId" db:"category_id"`
}

type ArticleTag struct {
	ArticleId string `json:"articleId" db:"article_id"`
	TagId     string `json:"tagId" db:"tag_id"`
}
