// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type Page struct {
	Title string
	Body  []byte
}

const (
	dataPath = "data/"
	tmplPath = "tmpl/"
)

func (p *Page) save() error {
	filename := dataPath + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0644)
}

func loadPage(title string) (*Page, error) {
	filename := dataPath + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fileInfos, err := ioutil.ReadDir(dataPath)
	if err != nil {
		return
	}

	// 写入Cookie
	expiration := time.Now()
	expiration = expiration.AddDate(0, 1, 1)
	cookie := http.Cookie{Name: "username",
		Value:  "tiny",
		Raw:    "Raw value",
		Path:   "/",
		Domain: "localhost",
		// Unparsed: []string{"1", "a", "3"},
		HttpOnly: true,
		// Secure:   true,
		Expires: expiration}
	http.SetCookie(w, &cookie)

	locals := make(map[string]interface{})
	pages := []string{}
	for _, fi := range fileInfos {
		name := fi.Name()
		name = name[:len(name)-4]
		pages = append(pages, name)
	}
	locals["pages"] = pages
	templates.ExecuteTemplate(w, "home.html", locals)
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)

	// 读取Cookie
	if cookie, err := r.Cookie("username"); err == nil {
		log.Println(cookie)
	}

	for _, cookie := range r.Cookies() {
		log.Println(cookie)
	}

}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

var templates = template.Must(template.ParseFiles(tmplPath+"edit.html", tmplPath+"view.html", tmplPath+"home.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)") // 测试querys

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}

		r.ParseForm()
		log.Println("r.Form", r.Form) //这些信息是输出到服务器端的打印信息
		log.Println("path", r.URL.Path)
		log.Println("scheme", r.URL.Scheme)
		for k, v := range r.Form {
			log.Println("key:", k)
			log.Println("val:", strings.Join(v, ","))
		}

		fn(w, r, m[2])
	}
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
