package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type dolors float64

type dateBase map[string]dolors

func (d dolors) String() string {
	return fmt.Sprintf("%.2f$", float64(d))
}

func (db dateBase) list(w http.ResponseWriter, r *http.Request) {
	for i, p := range db {
		fmt.Fprintf(w, "item %s: %s\n", i, p)
	}
}

func (db dateBase) add(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price := r.URL.Query().Get("price")

	if _, ok := db[item]; ok {
		msg := fmt.Sprintf("%s is duplicated", item)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	p, err := strconv.ParseFloat(price, 64)
	if err != nil {
		msg := fmt.Sprintf("%s is invalid", price)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	db[item] = dolors(p)
	fmt.Fprintf(w, "item %s with price %s added\n", item, db[item])

}

func (db dateBase) update(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price := r.URL.Query().Get("price")

	if _, ok := db[item]; !ok {
		msg := fmt.Sprintf("%s is not found", item)
		http.Error(w, msg, http.StatusNotFound)
		return
	}

	p, err := strconv.ParseFloat(price, 64)
	if err != nil {
		msg := fmt.Sprintf("%s is invalid", price)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	db[item] = dolors(p)
	fmt.Fprintf(w, "item %s with price %s updated\n", item, db[item])

}

func (db dateBase) fetch(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")

	if _, ok := db[item]; !ok {
		msg := fmt.Sprintf("%s is not found", item)
		http.Error(w, msg, http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "item %s has %s\n", item, db[item])
}

func (db dateBase) delete(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")

	if _, ok := db[item]; !ok {
		msg := fmt.Sprintf("%s is not found", item)
		http.Error(w, msg, http.StatusNotFound)
		return
	}

	delete(db, item)
	fmt.Fprintf(w, "item %s deleted\n", item)
}

func main() {
	db := dateBase{
		"shoes": 50,
		"socks": 30.1243,
		"shirt": 70}

	http.HandleFunc("/list", db.list)
	http.HandleFunc("/add", db.add)
	http.HandleFunc("/fetch", db.fetch)
	http.HandleFunc("/delete", db.delete)
	http.HandleFunc("/update", db.update)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
