package model

import (
	"gin-blog/errormsg"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Category Category `gorm:"foreignkey:Cid"` // 外键
	Title    string   `gorm:"type:varchar(100);not null" json:"title"`
	Cid      int      `gorm:"type:int;not null" json:"cid"`          // 类型id
	Desc     string   `gorm:"type:varchar(200)" json:"desc"`         // 文章简介
	Content  string   `gorm:"type:longtext;not null" json:"content"` // 文章主体
	Img      string   `gorm:"type:varchar(100)" json:"img"`          // 文章图片
}

// 添加文章
func CreateArticle(data *Article) int {
	err := db.Create(data).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS

}

// 删除文章
func DeleteArticle(id int) int {
	var art Article
	err := db.Where("id = ?", id).Delete(&art).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS

}

// 编辑文章
func EditArticle(id int, data *Article) int {
	var art Article
	updateMap := make(map[string]interface{}, 8)
	updateMap["title"] = data.Title
	updateMap["cid"] = data.Cid
	updateMap["desc"] = data.Desc
	updateMap["content"] = data.Content
	updateMap["img"] = data.Img
	err = db.Model(&art).Where("id = ?", id).Updates(updateMap).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS
}

// 获取单个文章
func OnlyArt(id int) (Article, int) {
	var art Article
	err := db.Preload("Category").Where("id = ?", id).First(&art).Error
	if err != nil {
		return art, errormsg.ERROR_ART_NOT_EXIST
	}
	return art, errormsg.SUCCESS
}

// 获取分类下的文章列表
func CateList(cid int, pageSize int, pageNum int) ([]Article, int, int64) {
	var artlist []Article
	var total int64
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid = ?", cid).Find(&artlist).Error
	// 计数
	db.Model(&artlist).Where("cid =?", cid).Count(&total)
	if err != nil {
		return artlist, errormsg.ERROR_CATE_NOT_EXIST, 0
	}
	return artlist, errormsg.SUCCESS, total
}

// 获取文章列表
func ListArt(param string, pageSize int, pageNum int) ([]Article, int, int64) {
	var artlist []Article
	var total int64
	err := db.Preload("Category").Where("title like ?", param+"%").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&artlist).Error
	db.Model(&artlist).Where("title like ?", param+"%").Count(&total)
	if err != nil {
		return artlist, errormsg.ERROR_CATE_NOT_EXIST, 0
	}
	return artlist, errormsg.SUCCESS, total
}
