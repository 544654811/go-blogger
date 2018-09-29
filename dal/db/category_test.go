package db

import (
	"testing"
)

func init() {
	dns := "root:123@tcp(localhost:3306)/blogger?parseTime=true"
	err := InitDb(dns)
	if err != nil {
		panic(err)
	}
}

func TestGetCategoryListByCategoryIds(t *testing.T) {
	ids := []int64{
		1, 2, 3,
	}
	list, _ := GetCategoryListByCategoryIds(ids)

	t.Logf("list length%v \n", len(list))
}
