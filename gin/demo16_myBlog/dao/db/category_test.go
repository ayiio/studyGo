package db

import "testing"

func init() {
	//parseTime=true 将mysql中的时间类型自动解析为go结构体中的时间类型
	dns := "xxx:xxx@tcp(localhost:3306)/xxx?parseTime=true"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

//获取单个的分类信息
func TestGetCategoryById(t *testing.T) {
	category, err := GetCategoryById(1)
	if err != nil {
		panic(err)
	}
	t.Logf("category: %#v\n", category)
}

//获取多个分类信息
func TestGetCategoryList(t *testing.T) {
	var categoryIds []int64 = []int64{1, 2, 3}
	categoryList, err := GetCategoryList(categoryIds)
	if err != nil {
		panic(err)
	}
	for _, v := range categoryList {
		t.Logf("category id:%d, category:%#v\n", v.Id, v)
	}
}

//获取所有分类信息
func TestGetCategoryAll(t *testing.T) {
	categoryList, err := GetCategoryAll()
	if err != nil {
		panic(err)
	}
	for _, v := range categoryList {
		t.Logf("category id:%d, category:%#v\n", v.Id, v)
	}
}
