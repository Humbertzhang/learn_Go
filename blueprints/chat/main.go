package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, nil)
}

func main() {
	//templateHandler实现了ServeHTTP,符合http.Handle的接口，所以可以直接在里面用
	//在访问 / 网页时，Handle函数会调用ServeHTTP函数来生成网页
	http.Handle("/", &templateHandler{filename: "chat.html"})
	fmt.Println("Started...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
