package db

import (
	"blogger/model"
	"fmt"
)

func GetCommentListByArticleId(articleId int64) (commentList []*model.Comment, err error) {
	if articleId == 0 {
		fmt.Println("invalid parameter, articleId=", articleId)
		return
	}

	sql := `select id, content, username, create_time, status, article_id
					from comment
					where 
						status = 1
					and article_id = ?`
	err = DB.Select(&commentList, sql, articleId)
	if err != nil {
		fmt.Println("GetCommentListByArticleId failed, ", err)
		return
	}
	return
}
