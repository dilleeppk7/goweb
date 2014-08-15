package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"log"
)


func main() {
	PORT :="3000"
	m := martini.Classic()


	/****** Render Engine *******/
	m.Use(render.Renderer(render.Options{	
  Directory: "templates", // Specify what path to load the templates from.
 // Layout: "layout", // Specify a layout template. Layouts can call {{ yield }} to render the current template.
  Extensions: []string{".ejs", ".html"}, // Specify extensions to load for templates.
 // Funcs: []template.FuncMap{AppHelpers}, // Specify helper function maps for templates to access.
  Delims: render.Delims{"{{", "}}"}, // Sets delimiters to the specified strings.
  Charset: "UTF-8", // Sets encoding for json and html content-types. Default is "UTF-8".
  IndentJSON: true, // Output human readable JSON
  IndentXML: true, // Output human readable XML
  //HTMLContentType: "application/xhtml+xml", // Output XHTML content type instead of default "text/html"
	}));





  /********** Routing ****************/
  m.Get("/",home)

  m.Get("/contact",contactus)

  

  /**********SERVER INIT**************/
  // pass control to martini
	http.Handle("/", m)
	fmt.Println("Server Listening on PORT "+PORT)
	//Listen Serve is blocking Call
	err := http.ListenAndServe(":"+PORT, nil)
	if err != nil {
		log.Fatal("Server initialize: ",err)
	}
}


func home(r render.Render){
	r.HTML(200, "home", "heelo")
}


func contactus() string{
	return "contact us at"
}
