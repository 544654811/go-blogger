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

func InsertComment(comment *model.Comment)error{
	sql := `insert into comment (content, username, article_id) values (?,?,?)`

	_, err := DB.Exec(sql, comment.Content, comment.Username, comment.ArticleID)
	if err != nil {
		fmt.Println("insert comment failed, ", err)
		return err
	}
	return nil
}
