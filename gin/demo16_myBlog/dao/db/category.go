package db

import (
	"demo16_myblog/model"

	"github.com/jmoiron/sqlx"
)

//分类相关操作： 添加，查询一或多或所有

//添加分类
func NewCategory(category *model.Category) (categoryId int64, err error) {
	sqlStr := "insert into category(category_name, category_no) values(?, ?)"
	res, err := DB.Exec(sqlStr, category.CategoryName, category.CategoryNo)
	if err != nil {
		return
	}
	categoryId, err = res.LastInsertId()
	return
}

//获取单个的文章分类
func GetCategoryById(id int64) (category *model.Category, err error) {
	category = &model.Category{}
	sqlStr := "select id, category_name, category_no from category where id=?"
	err = DB.Get(category, sqlStr, id)
	return
}

//获取多个文章分类
func GetCategoryList(categoryIds []int64) (categoryList []*model.Category, err error) {
	//构建Sql
	sqlStr, args, err := sqlx.In("select id, category_name, category_no from category where id in (?)", categoryIds)
	if err != nil {
		return
	}
	//查询
	err = DB.Select(&categoryList, sqlStr, args...)
	return
}

//获取所有文章分类
func GetCategoryAll() (categoryList []*model.Category, err error) {
	sqlStr := "select id, category_name, category_no from category order by category_no asc"
	err = DB.Select(&categoryList, sqlStr)
	return
}
