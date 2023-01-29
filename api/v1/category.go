package v1

import (
	"gin-blog/errormsg"
	"gin-blog/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 定义一个结构体
type CategoryApi struct {
}

// 添加分类
func (ca *CategoryApi) AddCate(c *gin.Context) {
	var data model.Category
	// 将用户数据绑定到结构体
	c.ShouldBindJSON(&data)
	code = model.CategoryExist(data.Name)
	// 如果不存在，则添加到数据库
	if code == errormsg.SUCCESS {
		model.CreateCategory(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errormsg.GetMsg(code),
	})
}

// 删除分类
func (ca *CategoryApi) DeleteCate(c *gin.Context) {
	// 获取id
	id, _ := strconv.Atoi(c.Param("id"))
	// 数据库操作
	code = model.DeleteCategory(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errormsg.GetMsg(code),
	})

}

// 编辑分类
func (ca *CategoryApi) EditCate(c *gin.Context) {
	var cate model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&cate)
	// 判断要修改的用户名是否已存在
	code = model.CategoryExist(cate.Name)
	if code == errormsg.SUCCESS {
		code = model.EditCategory(id, &cate)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errormsg.GetMsg(code),
	})
}

// 查询分类列表
func (ca *CategoryApi) ListCate(c *gin.Context) {
	// 转换为int类型
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	// 如果没有传参设置为-1，返回所有数据
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	// 从数据库中获取
	data, total := model.GetCategoryList(pageSize, pageNum)
	code = errormsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"msg":    errormsg.GetMsg(code),
		"total":  total,
	})

}
