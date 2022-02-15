package main

import (
	"github.com/goincremental/negroni-sessions"
	"github.com/goincremental/negroni-sessions/cookiestore"
	"github.com/julienschmidt/httprouter"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
	"gopkg.in/mgo.v2"
	"net/http"
)

var (
	renderer     *render.Render
	mongoSession *mgo.Session
)

func init() {
	renderer = render.New()
	s, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	mongoSession = s
}

const (
	sessionKey    = "simple_chat_session"
	sessionSecret = "simple_chat_session_secret"
)

func main() {
	router := httprouter.New()

	// 파일 실행 시 작업 디렉토리가 맞지 않을경우 동작하지 않음
	router.GET("/", func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		renderer.HTML(w, http.StatusOK, "index", map[string]string{"title": "Simple Chat!"})
	})
	// github.com/codegangsta/negroni 대신 github.com/urfave/negroni 모듈만 사용할 것!!
	router.GET("/login", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		renderer.HTML(w, http.StatusOK, "login", nil)
	})
	router.GET("/logout", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		sessions.GetSession(r).Delete(currentUserKey)
		http.Redirect(w, r, "/login", http.StatusFound)
	})
	router.GET("/auth/:action/:provider", loginHandler)
	router.POST("/rooms", createRoom)
	router.GET("/rooms", retrieveRooms)
	router.GET("/rooms/:id/messages", retrieveMessages)

	n := negroni.Classic()
	store := cookiestore.New([]byte(sessionSecret))
	n.Use(sessions.Sessions(sessionKey, store))
	n.Use(LoginRequired("/login", "/auth"))
	n.UseHandler(router)
	n.Run(":3000")
}
