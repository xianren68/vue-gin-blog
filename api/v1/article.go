package v1

import (
	"gin-blog/errormsg"
	"gin-blog/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArticleApi struct {
}

// 添加文章
func (a *ArticleApi) AddArt(c *gin.Context) {
	var data model.Article
	// 将文章数据绑定到结构体
	c.ShouldBindJSON(&data)
	code = model.CreateArticle(&data)

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errormsg.GetMsg(code),
	})
}

// 删除文章
func (a *ArticleApi) DeleteArt(c *gin.Context) {
	// 获取id
	id, _ := strconv.Atoi(c.Param("id"))
	// 数据库操作
	code = model.DeleteArticle(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errormsg.GetMsg(code),
	})

}

// 编辑文章
func (a *ArticleApi) EditArt(c *gin.Context) {
	var art model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&art)
	code = model.EditArticle(id, &art)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errormsg.GetMsg(code),
	})
}

// 获取单个文章
func (a *ArticleApi) OnlyArt(c *gin.Context) {
	var art model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	art, code = model.OnlyArt(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errormsg.GetMsg(code),
		"data":   art,
	})

}

// 获取分类下的文章列表
func (a *ArticleApi) CateList(c *gin.Context) {
	var artlist []model.Article
	cid, _ := strconv.Atoi(c.Param("cid"))
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	artlist, code, total := model.CateList(cid, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errormsg.GetMsg(code),
		"data":   artlist,
		"total":  total,
	})

}

// 获取文章列表
func (a *ArticleApi) ListArt(c *gin.Context) {
	var artlist []model.Article
	param := c.Query("param")
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	artlist, code, total := model.ListArt(param, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errormsg.GetMsg(code),
		"data":   artlist,
		"total":  total,
	})
}
