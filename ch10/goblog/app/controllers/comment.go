package controllers

import (
	"github.com/revel/revel"
	"goblog/app/models"
	"goblog/app/routes"
)

type Comment struct {
	GormController
}

func (c Comment) Create(postId int, body, commenter string) revel.Result {
	comment := models.Comment{PostId: postId, Body: body, Commenter: commenter}
	c.Txn.Create(&comment)
	c.Flash.Success("댓글 작성 완료")
	return c.Redirect(routes.Post.Show(postId))
}

func (c Comment) Destroy(postId, id int) revel.Result {
	c.Txn.Where("id = ?", id).Delete(&models.Comment{})
	c.Flash.Success("댓글 삭제 완료")
	return c.Redirect(routes.Post.Show(postId))
}
