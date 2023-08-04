package model

/** 搜索的/分页 **/
type Q struct {
	/** 搜索的内容 **/
	Content int `json:"content"`
	/** 当前页共多少条记录 **/
	PageSize int `json:"status"`
	/** 总共多少页 **/
	Total int `json:"total"`
	/** 当前页 **/
	CurrentPage int `json:"currentPage"`
	/** 根据什么排序 **/
	OrderBy string `json:"orderBy"`
}
