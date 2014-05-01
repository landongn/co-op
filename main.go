package main

import (
	"api/users"
	"fmt"
	"github.com/eaigner/hood"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func SetupDB() *hood.Hood {
	hd, err := hood.Open("postgres", "host=localhost user=coop_dev dbname=coop_dev password=coop_dev sslmode=disable")
	PanicIf(err)
	return hd
}
func main() {
	m := martini.Classic()

	m.Use(martini.Static("static/deploy"))
	fmt.Println("Setting up DB")
	m.Map(SetupDB())
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

	m.Run()
}
