package logic

import (
	"blogger/dal/db"
	"blogger/model"
)

func GetAllCategoryList() (categoryList []*model.Category, err error) {
	categoryList, err = db.GetCategoryList()
	return
}
