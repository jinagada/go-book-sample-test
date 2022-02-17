package controllers

import (
	"github.com/revel/revel"
	"goblog/app/models"
	"golang.org/x/crypto/bcrypt"
)

type App struct {
	GormController
	CurrentUser *models.User
}

func (c App) Login() revel.Result {
	return c.Render()
}

func (c App) CreateSession(username, password string) revel.Result {
	var user models.User
	c.Txn.Where(&models.User{Username: username}).First(&user)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err == nil {
		authKey := revel.Sign(user.Username)
		c.Session["authKey"] = authKey
		c.Session["username"] = user.Username
		c.Flash.Success("Welcome, " + user.Name)
		return c.Redirect(Post.Index)
	}
	for k := range c.Session {
		delete(c.Session, k)
	}
	c.Flash.Out["username"] = username
	c.Flash.Error("Login failed")
	return c.Redirect(Home.Index)
}

func (c App) DestorySession() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}
	return c.Redirect(Home.Index)
}

func (c *App) setCurrentUser() revel.Result {
	defer func() {
		if c.CurrentUser != nil {
			// RenderArgs -> ViewArgs 사용
			c.ViewArgs["currentUser"] = c.CurrentUser
		} else {
			delete(c.ViewArgs, "currentUser")
		}
	}()
	return nil
}
