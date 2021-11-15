package main

import (
	"fmt"
	"github.com/Elahi-Zahra/bookRoom/pkg/config"
	"github.com/Elahi-Zahra/bookRoom/pkg/handlers"
	"github.com/Elahi-Zahra/bookRoom/pkg/render"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

const pageNum = ":8080"
var app config.AppConfig
var session *scs.SessionManager
func main()  {

	//TODO change this to true when in production
	app.InProduction =false

	session = scs.New()
	session.Lifetime = 24*time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	//session.Cookie.Secure To be false for Port 8080
	session.Cookie.Secure =app.InProduction
	app.Session =session

	ts ,err :=render.CreatTemplateCash()
	if err !=nil{
		log.Fatal("can cot creat template cash:",err)
	}
	app.TemplateCash=ts
	app.UseCash=false
	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)
	render.NewTemplate(&app)
	fmt.Println(fmt.Sprintf("Page Number is : %s",pageNum))

	srv:=&http.Server{
		Addr: pageNum,
		Handler: routes(&app),
	}
	err =srv.ListenAndServe()
	log.Fatal(err)
}
