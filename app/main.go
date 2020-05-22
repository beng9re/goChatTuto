package main

import (
	"net/http"

	"github.com/codegangsta/negroni"
	sessions "github.com/goincremental/negroni-sessions"
	"github.com/goincremental/negroni-sessions/cookiestore"

	"github.com/julienschmidt/httprouter"
	"github.com/unrolled/render"
)

var renderer *render.Render

const (
	sessionKey    = "simple_chat_session"
	sessionSecret = "imple_chat_session_secret"
)

func init() {
	renderer = render.New()
}

func main() {

	// Http 전달 객체를 받는 라우터
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		renderer.HTML(w, http.StatusOK, "index", map[string]string{"title": "hellow"})
	})

	n := negroni.Classic()
	store := cookiestore.New([]byte(sessionKey))
	n.Use(sessions.Sessions(sessionKey, store))
	n.UseHandler(router)

	n.Run(":3000")

}
