package service

import (
	"demo16_myblog/dao/db"
	"demo16_myblog/model"
)

//获取所有分类
func GetAllCategoryList() (categoryList []*model.Category, err error) {
	categoryList, err = db.GetCategoryAll()
	if err != nil {
		return
	}
	return
}
