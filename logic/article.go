package logic

import (
	"blogger/dal/db"
	"blogger/model"
	"fmt"
	"math"
)

func AddArticleDetail(content, title, username string, categoryId int64) (err error) {
	articleDetail := &model.ArticleDetail{}
	articleDetail.Content = content
	articleDetail.ArticleInfo.CategoryId = categoryId
	articleDetail.ArticleInfo.CommentCount = 0
	articleDetail.ArticleInfo.Title = title
	articleDetail.ArticleInfo.Username = username
	articleDetail.ArticleInfo.ViewCount = 0
	c := []rune(content)
	length := int(math.Min(float64(len(c)), 128.0))
	articleDetail.ArticleInfo.Summary = string(c[:length])
	id, err := db.InsertArticle(articleDetail)
	if err != nil {
		fmt.Println("AddArticleDetail faild, ", err)
		return
	}
	fmt.Printf("add articleDetail success, id = %d \n", id)
	return
}

func UpdateArticleForViewCount(count uint32, articleId int64) (flag bool, err error) {
	retsult, err := db.UpdateArticleForViewCount(count, articleId)
	if retsult < 0 {
		fmt.Println("UpdateArticleForViewCount faild, articleId:", articleId)
		return
	}
	flag = true
	return
}

func GetArticleRelativeList(categoryId int64) (list []*model.ArticleRelative, err error) {
	list, err = db.GetArticleRelativeList(categoryId)
	return
}

func GetPreAndNextArticleByArticleId(articleId int64) (preArticle *model.ArticleRelative, nextArticle *model.ArticleRelative) {
	preArticle, err1 := db.GetPreArticleByArticleId(articleId)
	if err1 != nil {
		preArticle.ArticleID = -1
	}
	nextArticle, err2 := db.GetNextArticleByArticleId(articleId)
	if err2 != nil {
		nextArticle.ArticleID = -1
	}
	return
}

func GetArticleDetail(articleId int64) (detail *model.ArticleDetail, err error) {
	detail, err = db.GetArticleDetail(articleId)
	return
}

func GetArticleRecordList(pageNo, pageSize int) (list []*model.ArticleRecord, err error) {
	// 获取所有文章列表
	articleList, err := db.GetArticleList(pageNo, pageSize)
	if err != nil {
		fmt.Println("get articleList err, ", err)
		return
	}

	if articleList == nil || len(articleList) == 0 {
		return
	}
	// 得到文章列表中存在的分类id（不重复）
	categoryIds := getCategoryIds(articleList)
	// 获取所需的分类信息
	categoryList, err := db.GetCategoryListByCategoryIds(categoryIds)
	if err != nil {
		fmt.Println("get categoryList err, ", err)
		return
	}
	// 组装 articleRecord 信息
	list = buildArticleRecord(articleList, categoryList)
	return
}

func GetArticleRecordListByCategoryId(categroyId int64, pageNo, pageSize int) (list []*model.ArticleRecord, err error) {
	// 获取所有文章列表
	articleList, err := db.GetArticleListByCategoryId(categroyId, pageNo, pageSize)
	if err != nil {
		fmt.Println("get articleList err, ", err)
		return
	}

	if articleList == nil || len(articleList) == 0 {
		return
	}
	// 得到文章列表中存在的分类id（不重复）
	categoryIds := getCategoryIds(articleList)
	// 获取所需的分类信息
	categoryList, err := db.GetCategoryListByCategoryIds(categoryIds)
	if err != nil {
		fmt.Println("get categoryList err, ", err)
		return
	}
	// 组装 articleRecord 信息
	list = buildArticleRecord(articleList, categoryList)
	return
}

func buildArticleRecord(articleList []*model.ArticleInfo,
	categoryList []*model.Category) (list []*model.ArticleRecord) {
	for _, article := range articleList {
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *article,
		}
		for _, category := range categoryList {
			if article.CategoryId == category.CategoryId {
				articleRecord.Category = *category
			}
		}

		list = append(list, articleRecord)
	}
	return
}

func getCategoryIds(list []*model.ArticleInfo) (ids []int64) {
LABLE:
	for _, articleInfo := range list {
		categoryId := articleInfo.CategoryId
		for _, id := range ids {
			if id == categoryId {
				continue LABLE
			}
		}
		ids = append(ids, categoryId)
	}
	return
}
