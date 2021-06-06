package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type product struct {
	Name        string
	Price       float32
	Description string
}

var catalogue []product

func init() {
	for i := 1; i < 41; i++ {
		catalogue = append(catalogue, product{
			Name:        fmt.Sprintf("candle %d", i),
			Price:       40 + float32(i),
			Description: fmt.Sprintf("Nice candle %d", i),
		})
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/products", ProductsHandler)
	http.Handle("/", r)

	srv := &http.Server{
		Handler: r,
		Addr:    ":8000",
	}

	log.Fatal(srv.ListenAndServe())
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	if page := q.Get("page"); page != "" {
		p, err := strconv.Atoi(page)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, `{"message": "page should be integer"}`)
			return
		}
		c := GetDataPage(p)
		if c == nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, `{"message": "page query param should be between 1-%.0f"}`, math.Floor(float64(len(catalogue)/itemsPerPage)))
			return
		}
		b, err := json.Marshal(c)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, `{"message": "Cannot handle this request"}`)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, string(b))
		return
	}
	b, err := json.Marshal(catalogue)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"message": "Cannot handle this request"}`)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(b))

}

const (
	itemsPerPage = 5
)

func GetDataPage(page int) []product {
	start := (page - 1) * itemsPerPage
	stop := start + itemsPerPage
	if start > len(catalogue) {
		return nil
	}
	if stop > len(catalogue) {
		stop = len(catalogue)
	}
	return catalogue[start:stop]
}
