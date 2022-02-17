package controllers

import (
	"github.com/revel/revel"
	"goblog/app/models"
	"goblog/app/routes"
)

type Post struct {
	GormController
}

func (c Post) Index() revel.Result {
	var posts []models.Post
	c.Txn.Order("created_at desc").Find(&posts)
	return c.Render(posts)
}

func (c Post) New() revel.Result {
	post := models.Post{}
	return c.Render(post)
}

func (c Post) Create(title, body string) revel.Result {
	post := models.Post{Title: title, Body: body}
	c.Txn.Create(&post)
	c.Flash.Success("포스트 작성 완료")
	return c.Redirect(routes.Post.Index())
}

func (c Post) Show(id int) revel.Result {
	var post models.Post
	c.Txn.First(&post, id)
	c.Txn.Where(&models.Comment{PostId: id}).Find(&post.Comments)
	return c.Render(post)
}

func (c Post) Edit(id int) revel.Result {
	var post models.Post
	c.Txn.First(&post, id)
	return c.Render(post)
}

func (c Post) Update(id int, title, body string) revel.Result {
	var post models.Post
	c.Txn.First(&post, id)
	post.Title = title
	post.Body = body
	c.Txn.Save(&post)
	c.Flash.Success("포스트 수정 완료")
	return c.Redirect(routes.Post.Show(id))
}

func (c Post) Destroy(id int) revel.Result {
	c.Txn.Where("post_id = ?", id).Delete(&models.Comment{})
	c.Txn.Where("id = ?", id).Delete(&models.Post{})
	c.Flash.Success("포스트 삭제 완료")
	return c.Redirect(routes.Post.Index())
}
