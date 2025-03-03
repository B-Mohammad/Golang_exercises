package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type dolors float64

type db struct {
	items map[string]dolors
	mu    sync.Mutex
}

// type db map[string]dolors

func (d dolors) String() string {
	return fmt.Sprintf("%.2f$", d)
}

func (db *db) list(w http.ResponseWriter, r *http.Request) {
	db.mu.Lock()
	defer db.mu.Unlock()

	for i, p := range db.items {
		fmt.Fprintf(w, "item %s: %s\n", i, p)
	}
}

func (db *db) add(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price := r.URL.Query().Get("price")

	db.mu.Lock()
	defer db.mu.Unlock()

	if _, ok := db.items[item]; ok {
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

	db.items[item] = dolors(p)
	fmt.Fprintf(w, "item %s with price %s added\n", item, db.items[item])

}

func (db *db) update(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price := r.URL.Query().Get("price")

	db.mu.Lock()
	defer db.mu.Unlock()

	if _, ok := db.items[item]; !ok {
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

	db.items[item] = dolors(p)
	fmt.Fprintf(w, "item %s with price %s updated\n", item, db.items[item])

}

func (db *db) fetch(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")

	db.mu.Lock()
	defer db.mu.Unlock()

	if _, ok := db.items[item]; !ok {
		msg := fmt.Sprintf("%s is not found", item)
		http.Error(w, msg, http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "item %s has %s\n", item, db.items[item])
}

func (db *db) delete(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")

	db.mu.Lock()
	defer db.mu.Unlock()

	if _, ok := db.items[item]; !ok {
		msg := fmt.Sprintf("%s is not found", item)
		http.Error(w, msg, http.StatusNotFound)
		return
	}

	delete(db.items, item)
	fmt.Fprintf(w, "item %s deleted\n", item)
}

func main() {
	db := db{
		items: map[string]dolors{
			"shoes": 50,
			"socks": 30.1243,
			"shirt": 70},
	}

	http.HandleFunc("/list", db.list)
	http.HandleFunc("/add", db.add)
	http.HandleFunc("/fetch", db.fetch)
	http.HandleFunc("/delete", db.delete)
	http.HandleFunc("/update", db.update)

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
