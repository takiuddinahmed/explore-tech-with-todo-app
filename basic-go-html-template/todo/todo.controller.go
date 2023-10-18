package todo

import (
	"html/template"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, data []Todo) {
	temp, err := template.ParseFiles("todo/todo.template.html")

	if err != nil {
		panic(err)
	}
	err = temp.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

func Router() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()
			task := r.FormValue("task")
			Add(Todo{
				ID:     3,
				Task:   task,
				Status: false,
			})
		}
		renderTemplate(w, List())
	})

	return router
}
