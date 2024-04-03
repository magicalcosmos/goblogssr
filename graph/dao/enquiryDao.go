package dao

import (
	"fmt"
	"strconv"

	"github.com/magicalcosmos/goblogssr/graph/db"
	"github.com/magicalcosmos/goblogssr/graph/model"
	"github.com/magicalcosmos/goblogssr/graph/utils"
)

// Get enquiry list
func GetEnquiryList(input model.Q) *model.EnquiryWithPage {
	// 获取总数
	var count []int
	err := db.DB.Select(&count, "SELECT COUNT(*) FROM enquiry")
	if err != nil {
		fmt.Println("GetEnquiryList occur error: ", err)
	}

	// 获取具体的值
	limitStart := (input.CurrentPage - 1) * input.PageSize
	enquiryWithPage := &model.EnquiryWithPage{}
	err = db.DB.Select(&enquiryWithPage.Enquiry, `
		SELECT 
			* 
		FROM 
			enquiry 
		WHERE 
			1=1 
		ORDER BY 
			create_time 
		DESC LIMIT ?, ?
	`, limitStart, input.PageSize)
	if err != nil {
		fmt.Println("GetArticleList occur error: ", err)
	}
	enquiryWithPage.Page.CurrentPage = input.CurrentPage
	enquiryWithPage.Page.PageSize = input.PageSize
	enquiryWithPage.Page.Total = count[0]
	return enquiryWithPage
}

// Save enquiry
// TODO: userId 需要根据登录用户来传
func SaveEnquiry(data *model.Enquiry) (enquiry *model.Enquiry) {
	date := utils.GetCurrentTime()
	rows, err := db.DB.NamedExec(`
		INSERT INTO 
			enquiry (
				name,
				email,
				phone, 
				type,
				message,  
				create_time
			) VALUES (
				:name,
				:email,
				:phone, 
				:type, 
				:message, 
				:createTime
			)`,
		map[string]interface{}{
			"name":       data.Name,
			"email":      data.Email,
			"phone":      data.Phone,
			"type":       data.Type,
			"message":    data.Message,
			"createTime": date,
		})

	if err != nil {
		fmt.Println("SaveEnquiry occur error: ", err)
	}
	insertId, _ := rows.LastInsertId()

	s := strconv.FormatInt(insertId, 10)

	id, _ := strconv.Atoi(s)

	return GetEnquiryById(&model.Enquiry{
		ID: id,
	})
}

func GetEnquiryById(data *model.Enquiry) *model.Enquiry {
	enquiry := model.Enquiry{}
	err := db.DB.Get(&enquiry, "SELECT * FROM enquiry WHERE id=?", data.ID)
	if err != nil {
		fmt.Println("GetEnquiryById occur error: ", err)
	}
	return &enquiry
}
