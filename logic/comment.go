package logic

import (
	"blogger/dal/db"
	"blogger/model"
)

func GetCommentListByArticleId(articleId int64) (commentList []*model.Comment, err error) {
	commentList, err = db.GetCommentListByArticleId(articleId)
	return
}
