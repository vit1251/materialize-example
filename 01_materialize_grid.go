package main

import "log"

import "github.com/vit1251/materialize"


func main() {

	doc, err1 := materialize.NewDocument()
	log.Printf("err1 = %v", err1)

	s := materialize.NewStyle("https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0-beta/css/materialize.min.css")
	log.Printf("s = %v", s)
	doc.Head.AddStyle(*s)

        js := materialize.NewJavaScript("https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0-beta/js/materialize.min.js")
	doc.Head.AddJavaScript(*js)

	body, err2 := doc.Content()
	log.Printf("err2 = %v", err2)

	log.Printf("body = %v", body)

}
