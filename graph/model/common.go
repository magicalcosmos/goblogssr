package model

/** 搜索的/分页 **/
type Q struct {
	Content     string `json:"content"`
	PageSize    int    `json:"status"`
	Total       int    `json:"total"`
	CurrentPage int    `json:"currentPage"`
	OrderBy     string `json:"orderBy"`
}
