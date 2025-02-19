package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type todo struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var form = `<h1>Todo#{{.Id}}</h1>
<div>{{printf "%s (completed: %t)" .Title .Completed}}</div>`

func handler(w http.ResponseWriter, r *http.Request) {
	const url = "https://jsonplaceholder.typicode.com/"

	res, err := http.Get(url + r.URL.Path[1:])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("error:", err)
		return
	}

	defer res.Body.Close()

	var todo todo

	if err := json.NewDecoder(res.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("error:", err)
		return
	}

	fmt.Println("data:", todo)

	temp := template.New("todo")
	_, err = temp.Parse(form)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("error:", err)
		return
	}
	err = temp.Execute(w, todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("error:", err)
		return
	}
}

func main() {

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
