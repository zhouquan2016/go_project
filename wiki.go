package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}


func filename(title string) string {
	return title
}
func (p *Page)savePage() error {
	filename := filename(p.Title)
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page,error) {
	filename := filename(title)
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{
		Title: title,
		Body:  file,
	},nil
}

func handleIndex(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func handleView(w http.ResponseWriter, req *http.Request) {
	title := req.URL.Path[len("/view/"):]
	page,err := loadPage(title);
	if err != nil {
		printErr(w, err)
	}else {
		fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", page.Title, page.Body)
	}

}

func handleEdit(w http.ResponseWriter, req *http.Request) {
	title := req.URL.Path[len("/edit/"):]
	page,err := loadPage(title);
	if err != nil {
		printErr(w, err);
	}else {
		renderTemplate(w, "edit", page)
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

func handleSave(w http.ResponseWriter, req *http.Request) {
	content := req.FormValue("body")
	title := req.URL.Path[len("/edit/"):]
	if content != "" {
		page := Page{
			Title: title,
			Body: []byte(content),
		}
		err := page.savePage()
		if err != nil {
			printErr(w, err)
		}
	}
}
func printErr(w http.ResponseWriter, err error) {
	var errStr string
	if os.IsNotExist(err) {
		errStr = "文件不存在"
	}else {
		errStr = "系统异常"
	}
	fmt.Fprintf(w,"<h1>%s</h1>", errStr)
}
func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/view/", handleView)
	http.HandleFunc("/edit/", handleEdit)
	http.HandleFunc("/save/", handleSave)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
