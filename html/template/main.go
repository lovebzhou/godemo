package main

import (
	"bytes"
	"fmt"
	"html"
	"html/template"
)

func main() {
	t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	if err != nil {
		fmt.Println(err)
	}

	b := bytes.NewBuffer(make([]byte, 1024))
	err = t.ExecuteTemplate(b, "T", "<script>alert('you have been pwned')</script>")
	s := b.String()
	fmt.Println(s)
	fmt.Println(html.UnescapeString(b.String()))

}
