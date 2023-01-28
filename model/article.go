package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Category Category `gorm:"foreignkey:Cid"` // 外键
	Title    string   `gorm:"type:varchar(100);not null" json:"title"`
	Cid      int      `gorm:"type:int;not null" json:"cid"`          // 类型id
	Desc     string   `gorm:"type:varchar(200)" json:"desc"`         // 文章简介
	Content  string   `gorm:"type:longtext;not null" json:"content"` // 文章主体
	Img      string   `gorm:"type:varchar(100)" json:"img"`          // 文章图片
}
