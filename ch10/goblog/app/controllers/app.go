package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"goblog/app/models"
	"golang.org/x/crypto/bcrypt"
)

type App struct {
	GormController
	// 최신 revel 에서 구조체에 선언한 객체의 값이 공유되지 않는것을 확인 아래 내용 주석 처리
	//CurrentUser *models.User
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

func (c App) setCurrentUser() revel.Result {
	// 책 예제에서는 (c *App) 으로 되어있어서 로그아웃이 되지 않음
	// (c App) 으로 변경해야만 로그아웃이 정상 동작함
	// c.CurrentUser 사용 할 수 없음
	existUser := false
	defer func() {
		//if c.CurrentUser != nil {
		//	// RenderArgs -> ViewArgs 사용
		//	c.ViewArgs["currentUser"] = c.CurrentUser
		//} else {
		//	delete(c.ViewArgs, "currentUser")
		//}
		if !existUser {
			delete(c.ViewArgs, "currentUser")
		}
	}()
	username, ok := c.Session["username"]
	usernameStr := fmt.Sprintf("%v", username)
	if !ok || usernameStr == "" {
		return nil
	}
	authKey, ok := c.Session["authKey"]
	authKeyStr := fmt.Sprintf("%v", authKey)
	if !ok || authKey == "" {
		return nil
	}
	if match := revel.Verify(usernameStr, authKeyStr); match {
		var user models.User
		c.Txn.Where(&models.User{Username: usernameStr}).First(&user)
		if &user != nil {
			c.ViewArgs["currentUser"] = &user
			existUser = true
		}
	}
	return nil
}
