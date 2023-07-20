package service

import (
	"demo16_myblog/dao/db"
	"demo16_myblog/model"
)

//获取文章和对应的分类
func GetArticleRecordList(pageNum, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {
	//获取文章列表
	articleInfoList, err := db.GetArticleList(pageNum, pageSize)
	if err != nil || len(articleInfoList) <= 0 {
		return
	}
	//聚合category

	// //存在性能问题，每个文章都要有一次db，效率较低
	// for _, articleInfo := range articleInfoList {
	// 	category, err := db.GetCategoryById(articleInfo.CategoryId)
	// 	if err != nil {
	// 		continue
	// 	}
	// 	articleRecord := &model.ArticleRecord{}
	// 	articleRecord.ArticleInfo = *articleInfo
	// 	articleRecord.Category = *category
	// 	articleRecordList = append(articleRecordList, articleRecord)
	// }

	categoryIds := getCategoryIds(articleInfoList)
	//根据制定的categoryId列表获取到对应的category列表
	categoryList, err := db.GetCategoryList(categoryIds)
	if err != nil {
		return
	}

	//遍历所有文章，聚合article和category
	for _, article := range articleInfoList {
		//根据当前文章生成结构体
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *article,
		}
		//取出文章的分类id
		articleCategoryID := article.CategoryId
		//遍历分类列表
		for _, category := range categoryList {
			if articleCategoryID == category.Id {
				articleRecord.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList, articleRecord)
	}
	return
}

//根据多个文章id获取多个分类id的集合
func getCategoryIds(articleInfoList []*model.ArticleInfo) (ids []int64) {
	//遍历文章list，取出分类ID
	categorySet := make(map[int64]bool, 16)
	for _, article := range articleInfoList {
		categoryId := article.CategoryId
		//加ids前去重
		_, ok := categorySet[categoryId]
		if !ok {
			categorySet[categoryId] = true
		}
	}
	for k := range categorySet {
		ids = append(ids, k)
	}
	return
}

//根据分类id，获取该类文章和对应的分类信息
func GetArticleRecordListByID(categoryId, pageNum, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {
	//获取文章列表
	articleList, err := db.GetArticleListByCategoryId(categoryId, pageNum, pageSize)
	if err != nil || len(articleList) <= 0 {
		return
	}
	//获取文章对应分类
	categoryIds := getCategoryIds(articleList)
	//根据categoryIDs获取到对应的分类列表
	categoryList, err := db.GetCategoryList(categoryIds)

	//遍历文章列表，聚合分类
	for _, article := range articleList {
		//创建articleRecord结构体
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *article,
		}
		//匹配分类
		articleCategoryId := article.CategoryId
		for _, category := range categoryList {
			if articleCategoryId == category.Id {
				articleRecord.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList, articleRecord)
	}
	return
}
