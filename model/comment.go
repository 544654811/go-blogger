package model

import "time"

type Comment struct {
	ID         int64     `db:"id"`
	Content    string    `db:"content"`
	Username   string    `db:"username"`
	CreateTime time.Time `db:"create_time"`
	Status     int       `db:"status"`
	ArticleID  int64     `db:"article_id"`
}
