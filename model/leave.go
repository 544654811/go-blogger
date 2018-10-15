package model

import "time"

type Leave struct{
	Id           int64     `db:"id"`
	Content   	 string     `db:"content"`
	Username     string    `db:"username"`
	Email        string      `db:"email"`
	CreateTime   time.Time `db:"create_time"`
}
