package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Todo struct {
	Title string
	Done  bool
	ID    int
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

var todos []Todo = []Todo{
	{Title: "Task 1", Done: false, ID: 1},
	{Title: "Task 2", Done: true, ID: 2},
	{Title: "Task 3", Done: true, ID: 3},
}
var todoFormTmpl *template.Template = template.Must(template.ParseFiles("./src/components/todo.html", "./src/components/tick.html", "./src/components/cross.html"))
var layoutTmpl *template.Template = template.Must(template.ParseFiles("./src/pages/layout.html", "./src/components/todo.html", "./src/components/tick.html", "./src/components/cross.html"))

func doneHandler(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Path
	id, err := strconv.Atoi(idString)

	if err == nil {
		index := -1
		for k, s := range todos {
			if s.ID == id {
				index = k
				break
			}
		}
		if index != -1 {
			todos[index].Done = true
		}
	}

	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos:     todos,
	}
	todoFormTmpl.ExecuteTemplate(w, "todos", data)
}

func main() {

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos:     todos,
		}
		layoutTmpl.ExecuteTemplate(w, "base", data)
	})

	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {

		todos = append(todos, Todo{
			Title: "hello",
			Done:  false,
			ID:    len(todos) + 1,
		})
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos:     todos,
		}
		todoFormTmpl.ExecuteTemplate(w, "todos", data)
	})

	http.HandleFunc("/create-new-todo", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		item := r.Form.Get("new-item")

		fmt.Printf("%s", item)
		todos = append(todos, Todo{
			Title: item,
			Done:  false,
			ID:    len(todos) + 1,
		})
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos:     todos,
		}
		todoFormTmpl.ExecuteTemplate(w, "todos", data)
	})

	http.Handle("/done/", http.StripPrefix("/done/", http.HandlerFunc(doneHandler)))

	http.ListenAndServe(":80", nil)
}
