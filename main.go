package main

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/codegangsta/martini-contrib/sessions"

	"log"
	"os"
)

var conf Conf

func init() {
	if err := Flags(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	err := Config() // setups conf var
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func main() {
	//println(conf.servername)
	m := martini.Classic()
	store := sessions.NewCookieStore([]byte("secret123"))
	m.Use(sessions.Sessions("my_session", store))
	m.Use(render.Renderer())

	m.Get("/", func(session sessions.Session) string {
		isLogin := IsLogin(session)
		if isLogin {
			return "login"
		} else {
			return "no login"
		}
	})

	m.Get("/sign_in", SignIn)
	m.Post("/login", LogIn)

	m.Run()
}
