package main

import (
	"api/auth"
	"api/users"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/jinzhu/gorm"
	"github.com/martini-contrib/render"
)

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func SetupDB() *gorm.DB {
	DB, err := gorm.Open("postgres", "host=localhost user=coop_production dbname=coop_production password=coop_production sslmode=disable")
	PanicIf(err)
	DB.LogMode(true)
	DB.AutoMigrate(users.Profile{})
	return &DB
}

func main() {
	m := martini.Classic()

	m.Use(martini.Static("static/deploy"))
	fmt.Println("Setting up DB")
	m.Map(SetupDB())
	m.Use(martini.Logger())

	m.Use(render.Renderer(render.Options{
		// Specify what path to load the templates from.
		Directory: "templates",
		// Specify a layout template. Layouts can call {{ yield }} to render the current template.
		// Layout: "base",
		// Specify extensions to load for templates.
		Extensions: []string{".hbs", ".html"},
		// Sets delimiters to the specified strings.
		Delims: render.Delims{"{[{", "}]}"},
		// Sets encoding for json and html content-types. Default is "UTF-8".
		Charset: "UTF-8",
		// Output human readable JSON
		IndentJSON: true,
		// Output XHTML content type instead of default "text/html"
		// HTMLContentType: "application/xhtml+xml",
	}))

	m.Get("/", func(r render.Render) {
		r.HTML(200, "base", nil)
	})

	m.Post("/account/signup", users.CreateProfile)
	m.Post("/account/login", login.AttemptLoginForUser)
	m.Run()
}
