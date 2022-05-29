package service

import (
	"blog-go/config"
	"blog-go/dao"
	"blog-go/models"
	"html/template"
)

func GetAllIndexInfo(page, pageSize int) (*models.HomeResponse, error) {

	// 页面上的数据必须定义
	categorys, err := dao.GetAllCategory()
	if err != nil {
		return nil, err
	}

	posts, err := dao.GetPostPage(page, pageSize)
	if err != nil {
		return nil, err
	}
	var postMores []models.PostMore
	for _, post := range posts {
		categoryName := dao.GetCategoryNameById(post.CategoryId)
		userName := dao.GetUserNameById(post.UserId)
		content := []rune(post.Content)
		if len(content) > 100 {
			content = content[0:100]
		}

		postMore := models.PostMore{
			post.Pid,
			post.Title,
			post.Slug,
			template.HTML(content),
			post.CategoryId,
			categoryName,
			post.UserId,
			userName,
			post.ViewCount,
			post.Type,
			post.CreateAt.Format("2006-01-02 15:04:05"),
			post.UpdateAt.Format("2006-01-02 15:04:05"),
		}
		postMores = append(postMores, postMore)
	}

	hr := &models.HomeResponse{
		config.Cfg.Viewer,
		categorys,
		postMores,
		1,
		1,
		[]int{1},
		true,
	}

	return hr, nil
}
