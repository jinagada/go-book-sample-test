package controllers

import (
	db "github.com/revel/modules/db/app"
	"github.com/revel/revel"
	"goblog/app/routes"
	"time"
)

type Comment struct {
	*revel.Controller
	db.Transactional
}

func (c Comment) Create(postId int, body, commenter string) revel.Result {
	_, err := c.Txn.Exec("insert into comments(body, commenter, post_id, created_at, updated_at) values (?, ?, ?, ?, ?)", body, commenter, postId, time.Now(), time.Now())
	if err != nil {
		panic(err)
	}
	c.Flash.Success("댓글 작성 완료")
	return c.Redirect(routes.Post.Show(postId))
}

func (c Comment) Destroy(postId, id int) revel.Result {
	if _, err := c.Txn.Exec("delete from comments where id = ?", id); err != nil {
		panic(err)
	}
	c.Flash.Success("댓글 삭제 완료")
	return c.Redirect(routes.Post.Show(postId))
}
