package main

import (
	"basic-go-html-template/todo"
	"net/http"
)



func main() {
	router := todo.Router()
	http.ListenAndServe(":3000",router);
}
