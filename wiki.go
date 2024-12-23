package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler)
	// log.Fatal(http.ListenAndServe(":8080", nil))
	http.HandleFunc("/viev/", vievHandler)
	http.HandleFunc(("/edit/"), editHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func renderTemplate(w http.ResponseWriter, temp string, page *Page) {
	t, _ := template.ParseFiles(temp)
	t.Execute(w, page)
}

func vievHandler(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Path[len("/viev/"):]
	page, err := loadPage(filename)
	if err != nil {
		log.Fatal(err)
	}
	renderTemplate(w, "viev", page)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("edit/"):]
	page, err := loadPage(title)
	if err == nil {
		log.Fatal(err)
	}
	renderTemplate(w, "edit", page)
}

type Page struct {
	Title string
	Body  []byte
}

// почему мы используем указатель
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600) //возвращаем error потому что его возвращает этот метод
	//0600 - read/write
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{title, body}, nil
}
