package db

import (
	"blogger/model"
	"fmt"
)

func GetAllLeaveList()(list []*model.Leave, err error){
	sql := `select id, content, username, email, create_time from leaves`

	err = DB.Select(&list, sql)
	if err != nil {
		fmt.Println("GetAllLeaveList error, ", err)
	}
	return
}

func InsertLeave(leave *model.Leave)error{
	sql := `insert into leaves(content, username, email) values(?,?,?)`

	_, err := DB.Exec(sql, leave.Content, leave.Username, leave.Email)
	if err != nil {
		fmt.Println("InsertLeave error, ", err)
	}

	return err
}
