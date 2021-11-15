package render

import (
	"bytes"
	"fmt"
	"github.com/Elahi-Zahra/bookRoom/pkg/config"
	"github.com/Elahi-Zahra/bookRoom/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var fanctions =template.FuncMap{

}
var app *config.AppConfig

func NewTemplate(a *config.AppConfig)  {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}
//RenderTemplate renders template using html/template
func RenderTemplate(w http.ResponseWriter,tmpName string,td *models.TemplateData)  {

	// get the template cash from the app config
	var tc map[string]*template.Template
	if app.UseCash{
		tc =app.TemplateCash
	}else {
		tc,_=CreatTemplateCash()
	}
	t,ok := tc[tmpName]
	if !ok{
		log.Fatal("can not get template from template cash")
	}

	buf :=new(bytes.Buffer)
	td = AddDefaultData(td)
	_=t.Execute(buf,td)
	_,err := buf.WriteTo(w)
	if err != nil{
		fmt.Println("Error witting template to browser:",err)
	}
}
func CreatTemplateCash() (map[string]*template.Template,error) {
	myCash := map[string]*template.Template{}
	pages , err :=filepath.Glob("./templates/*.page.tmpl")
	fmt.Println("page",pages,err)
	if err !=nil{
		return myCash,err
	}
	for _,page := range pages{
		name := filepath.Base(page)
		fmt.Println("Page is Currently",name)
		ts,err := template.New(name).Funcs(fanctions).ParseFiles(page)
		if err!=nil{
			return myCash,err
		}

		maches, err :=filepath.Glob("./templates/*.layout.tmpl")
		if err !=nil{
			return myCash,err
		}

		if len(maches)>0{
			ts ,err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err !=nil{
				return myCash,err
			}
		}
		myCash[name] = ts
	}
	return myCash,err
}