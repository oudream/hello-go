package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	PageTime  float64
	Todos     []Todo
}

func main() {
	helloHttpResp1()
	//hello2file1()
}

func hello2file1() {
	t, err := template.ParseFiles("layout.html")
	if err != nil {
		log.Print(err)
		return
	}
	p := time.Now().Format("2006-01-02T15-04-05") + ".html"
	f, err := os.Create(p)
	if err != nil {
		log.Println("create file: ", err)
		return
	}
	data := TodoPageData{
		PageTitle: "My TODO list",
		PageTime:  1.23456677,
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}
	err = t.Execute(f, data)
	if err != nil {
		log.Println("execute: ", err)
		return
	}
	log.Println("out put path: ", p)
}

func helloHttpResp1() {
	tmpl := template.Must(template.ParseFiles("layout.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			PageTime:  1.23456677,
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":80", nil)
}