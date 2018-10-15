package controller

import (
	"blogger/logic"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ToIndex(c *gin.Context) {
	atricleRecordList, err := logic.GetArticleRecordList(1, 15)
	if err != nil {
		fmt.Println("get articleRecord list failed, ", err)
		return
	}

	allCategoryList, err := logic.GetAllCategoryList()
	if err != nil {
		fmt.Println("get allCategory list failed, ", err)
		return
	}

	var data map[string]interface{} = make(map[string]interface{}, 15)
	data["article_list"] = atricleRecordList
	data["category_list"] = allCategoryList

	c.HTML(http.StatusOK, "views/index.html", data)
}

func ToCategory(c *gin.Context) {
	categoryIdStr := c.Query("category_id")
	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
	if err != nil{
		fmt.Println("ToCategory-invalid parameter, ", err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	atricleRecordList, err := logic.GetArticleRecordListByCategoryId(categoryId,1, 15)
	if err != nil {
		fmt.Println("get articleRecord list failed, ", err)
		return
	}

	allCategoryList, err := logic.GetAllCategoryList()
	if err != nil {
		fmt.Println("get allCategory list failed, ", err)
		return
	}

	var data map[string]interface{} = make(map[string]interface{}, 15)
	data["article_list"] = atricleRecordList
	data["category_list"] = allCategoryList

	c.HTML(http.StatusOK, "views/index.html", data)
}

func ToArticleDetail(c *gin.Context) {
	articleIDStr := c.Query("article_id")
	articleID, err := strconv.ParseInt(articleIDStr, 10, 64)
	if err != nil {
		fmt.Println("articleIdStr parseInt failed, ", err)
		return
	}

	// 获取文章详情
	articleDetail, err := logic.GetArticleDetail(articleID)
	if err != nil {
		fmt.Println("ToArticleDetail-GetArticleDetail failed, ", err)
		return
	}
	// 获取评论信息
	commentList, err := logic.GetCommentListByArticleId(articleID)
	if err != nil {
		fmt.Println("ToArticleDetail-GetCommentListByArticleId failed, ", err)
		return
	}
	// 获取上下篇文章信息
	preAtricle, nextAtricle := logic.GetPreAndNextArticleByArticleId(articleID)

	// 获取文章的相关文章
	articleRelativeList, err := logic.GetArticleRelativeList(articleDetail.ArticleInfo.CategoryId)
	if err != nil {
		fmt.Println("ToArticleDetail-GetArticleRelativeList failed, ", err)
		return
	}
	// 文章查看数增加
	count := articleDetail.ArticleInfo.ViewCount + 1
	articleDetail.ArticleInfo.ViewCount = count
	flag, err := logic.UpdateArticleForViewCount(count, articleID)
	if err != nil || !flag {
		fmt.Println("ToArticleDetail-UpdateArticleForViewCount failed, ", err)
		return
	}
	// 栏目导航
	categorys, err := logic.GetAllCategoryList()
	if err != nil || !flag {
		fmt.Println("ToArticleDetail-GetAllCategoryList failed, ", err)
		return
	}

	data := make(map[string]interface{})
	data["detail"] = articleDetail
	data["comment_list"] = commentList
	data["prev"] = preAtricle
	data["next"] = nextAtricle
	data["relative_article"] = articleRelativeList
	data["categorys"] = categorys
	data["articleId"] = articleID

	c.HTML(http.StatusOK, "views/detail.html", data)
}

func ToPostArticle(c *gin.Context) {
	categoryList, err := logic.GetAllCategoryList()
	if err != nil {
		fmt.Println("ToPostArticle-GetAllCategoryList failed, ", err)
		return
	}
	c.HTML(http.StatusOK, "views/post_article.html", categoryList)
}

func PostArticleHandle(c *gin.Context) {
	username := c.PostForm("author")
	title := c.PostForm("title")
	categoryIdStr := c.PostForm("category_id")
	content := c.PostForm("content")
	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
	if err != nil {
		fmt.Println("ToPostArticle-invalid parameter, ", err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}

	err = logic.AddArticleDetail(content, title, username, categoryId)
	if err != nil {
		fmt.Println("PostArticleHandle-AddArticleDetail failed, ", err)
		return
	}

	c.Redirect(http.StatusMovedPermanently, "/")
}

func PostCommentHandle(c *gin.Context){
	articleIdStr := c.PostForm("article_id")
	comment := c.PostForm("comment")
	author := c.PostForm("author")
	//fmt.Printf("articleIdStr: %s, comment: %s, author: %s /n", articleIdStr, comment, author)

	articleId, err := strconv.ParseInt(articleIdStr, 10, 64)
	if err != nil {
		fmt.Printf("articleIdStr parse articleId error: %v", err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
	}
	
	logic.InstertComment(comment, author, articleId)
	url := fmt.Sprintf("/article/detail?article_id=%d", articleId)
	c.Redirect(http.StatusMovedPermanently, url)
}

func ToLeave(c *gin.Context){
	list, err := logic.GetAllLeave()
	if err != nil{
		fmt.Printf("ToLeave GetAllLeave error, %v", err)
	}
	c.HTML(http.StatusOK, "views/gbook.html", list)
}

func PostLeaveHandle(c *gin.Context){
	comment := c.PostForm("comment")
	author := c.PostForm("author")
	email := c.PostForm("email")
	err := logic.InsertLeave(comment, author, email)
	if err != nil{
		fmt.Printf("PostLeaveHandle InsertLeave error, %v", err)
	}
	c.Redirect(http.StatusMovedPermanently, "/leave/new/")
}
