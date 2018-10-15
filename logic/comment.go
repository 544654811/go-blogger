package logic

import (
	"blogger/dal/db"
	"blogger/model"
	"fmt"
)

func GetCommentListByArticleId(articleId int64) (commentList []*model.Comment, err error) {
	commentList, err = db.GetCommentListByArticleId(articleId)
	return
}

func InstertComment(comment, author string, articleId int64)(err error){
	// 判断文章id是否存在
	detail, err := db.GetArticleDetail(articleId)
	if err != nil {
		fmt.Println("InstertComment GetArticleDetail error: ", err)
		return
	}
	if detail == nil{
		err = fmt.Errorf("article not found, articleId=%d", articleId)
		return
	}

	// 插入评论信息
	c := &model.Comment{
		Content:comment,
		ArticleID:articleId,
		Username:author,
	}
	err = db.InsertComment(c)
	if err == nil {
		// 更新文章评论数
		count := detail.CommentCount + 1
		_, err = db.UpdateArticleForCommentCount(count, articleId)
	}
	return
}
