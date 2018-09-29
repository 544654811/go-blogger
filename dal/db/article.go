package db

import (
	"blogger/model"
	"fmt"
)

func InsertArticle(articleDetail *model.ArticleDetail) (articleId int64, err error) {
	sql := `insert article
						(category_id, title, view_count, comment_count, username, summary, content)
					values (?, ?, ?, ?, ?, ?, ?)`

	res, err := DB.Exec(sql, articleDetail.ArticleInfo.CategoryId, articleDetail.Title, articleDetail.ViewCount,
		articleDetail.CommentCount, articleDetail.Username, articleDetail.Summary, articleDetail.Content)

	if err != nil {
		return
	}
	articleId, err = res.LastInsertId()
	return
}

func GetArticleList(pageNo, pageSize int) (list []*model.ArticleInfo, err error) {
	if pageNo < 0 || pageSize < 0 {
		err = fmt.Errorf("invalid parameter, pageNo = %d, pageSize = %d", pageNo, pageSize)
		return
	}

	if pageNo == 0 {
		pageNo = 1
	}

	sql := `select id, category_id, title, view_count, comment_count, username, status, summary, create_time
					from article
					where 
						status = 1
					order by create_time desc
					limit ?,?`

	err = DB.Select(&list, sql, pageNo-1, pageSize)
	if err != nil {
		return
	}
	return
}

func GetArticleDetail(articleId int64) (detail *model.ArticleDetail, err error) {
	detail = &model.ArticleDetail{}
	sql := `select id, category_id, title, view_count, comment_count, username, status, summary, create_time, content
					from article
					where 
						status = 1
						and id = ?
					`
	err = DB.Get(detail, sql, articleId)
	return
}

func GetPreArticleByArticleId(articleId int64) (articleRelative *model.ArticleRelative, err error) {
	articleRelative = &model.ArticleRelative{}

	sql := `select id, title from article where id < ? order by id desc limit 1`
	err = DB.Get(articleRelative, sql, articleId)
	if err != nil {
		fmt.Println("GetPreArticleByArticleId sql failed, ", err)
		return
	}
	return
}

func GetNextArticleByArticleId(articleId int64) (articleRelative *model.ArticleRelative, err error) {
	articleRelative = &model.ArticleRelative{}

	sql := `select id, title from article where id > ? order by id asc limit 1`
	err = DB.Get(articleRelative, sql, articleId)
	if err != nil {
		fmt.Println("GetPreArticleByArticleId sql failed, ", err)
		return
	}
	return
}

func GetArticleRelativeList(categoryId int64) (list []*model.ArticleRelative, err error) {
	sql := `select id, title from article where status = 1 and category_id = ? order by create_time desc`

	err = DB.Select(&list, sql, categoryId)
	if err != nil {
		fmt.Println("GetArticleRelativeList sql failed, ", err)
		return
	}
	return
}

func UpdateArticleForViewCount(count uint32, articleId int64) (res int64, err error) {
	sql := `update article set view_count = ? where id = ?`
	result, err := DB.Exec(sql, count, articleId)
	if err != nil {
		fmt.Println("UpdateArticleForViewCount sql failed, ", err)
		return
	}
	res, err = result.RowsAffected()
	if err != nil {
		fmt.Println("UpdateArticleForViewCount result failed, ", err)
		return
	}
	fmt.Println("articleId:", articleId)
	fmt.Println("count:", count)
	return
}
