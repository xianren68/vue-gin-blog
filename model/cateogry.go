package model

import (
	"gin-blog/errormsg"

	"gorm.io/gorm"
)

type Category struct {
	Id   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// 添加分类
func CreateCategory(data *Category) int {
	err := db.Create(data)
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS

}

// 删除分类
func DeleteCategory(id int) int {
	var cate Category
	err := db.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS

}

// 编辑用户
func EditCategory(id int, data *Category) int {
	var cate Category
	updateMap := make(map[string]interface{}, 4)
	updateMap["name"] = data.Name
	err = db.Model(&cate).Where("id = ?", id).Updates(updateMap).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS
}

// 查询分类列表
func GetCategoryList(pageSize int, pageNum int) ([]*Category, int64) {
	// 用来存储列表的切片
	var cate []*Category
	var total int64
	// 从数据库中获取
	// 判断-1的情况
	var offset int
	if pageNum == -1 {
		offset = -1
	} else {
		offset = (pageNum - 1) * pageSize
	}
	err := db.Limit(pageSize).Offset(offset).Find(&cate).Error
	db.Model(&cate).Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return cate, total

}

// 判断分类名是否存在
func CategoryExist(name string) int {
	var cate Category
	// 查询用户名是否存在于数据库中,并将值返回到user中
	db.Select("id").Where("name = ?", name).Find(&cate)
	if cate.Id > 0 {
		return errormsg.ERROR_CATEGORY_USED
	}
	return errormsg.SUCCESS
}
