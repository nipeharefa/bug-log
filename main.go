package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	bb := NewBugsnag()

	go bb.Watch()
	r := chi.NewRouter()
	r.Use(bb.Handler)
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		u, ok := r.Context().Value("bugsnag").(ReportFunc)
		fmt.Println(u, ok)
		if ok {
			query := r.URL.Query()
			count, _ := strconv.Atoi(query.Get("count"))
			for i := 0; i < count; i++ {
				u(errors.New("error"))
			}
		}
		w.Write([]byte("welcome"))
	})
	r.Get("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	http.ListenAndServe(":3000", r)
}
