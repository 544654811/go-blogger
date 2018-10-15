package logic

import (
	"blogger/model"
	"blogger/dal/db"
)

func GetAllLeave()(list []*model.Leave, err error){
	return db.GetAllLeaveList()
}

func InsertLeave(content, username, email string) error{
	l := &model.Leave{
		Content:content,
		Username:username,
		Email:email,
	}
	return db.InsertLeave(l)
}
