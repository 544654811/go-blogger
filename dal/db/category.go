package db

import (
	"blogger/model"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func GetCategoryListByCategoryIds(categoryIds []int64) (categoryList []*model.Category, err error) {
	if len(categoryIds) == 0 {
		err = fmt.Errorf("invalid parameter, catcategoryIds = %v", categoryIds)
		return
	}

	sql := `select id, category_name, category_no from category where id in (?)`
	sqlStr, args, err := sqlx.In(sql, categoryIds)
	if err != nil {
		fmt.Println("GetCategoryListByCategoryIds failed, ", err)
		return
	}

	err = DB.Select(&categoryList, sqlStr, args...)
	if err != nil {
		fmt.Println("GetCategoryListByCategoryIds failed, ", err)
		return
	}
	return
}

func GetCategoryList() (categoryList []*model.Category, err error) {
	sql := `select id, category_name, category_no from category`

	err = DB.Select(&categoryList, sql)
	if err != nil {
		fmt.Println("GetAllCategoryList failed, ", err)
		return
	}
	return
}
