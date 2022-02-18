package controllers

import (
	"github.com/revel/revel"
)

func init() {
	revel.OnAppStart(InitDB)
	revel.InterceptMethod((*GormController).Begin, revel.BEFORE)
	revel.InterceptMethod((*GormController).Commit, revel.AFTER)
	revel.InterceptMethod((*GormController).Rollback, revel.FINALLY)
	revel.InterceptMethod((*App).setCurrentUser, revel.BEFORE)
	revel.InterceptMethod(Post.CheckUser, revel.BEFORE)
	revel.InterceptMethod(Comment.CheckUser, revel.BEFORE)
}
