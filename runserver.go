package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func RunTech() {
	r := mux.NewRouter()
	r.HandleFunc("/admin/redirects", DataControl)
	r.HandleFunc("/admin/redirects/{id:[0-9]+}", DataControl)
	r.HandleFunc("/redirects", Redirect)
	err := http.ListenAndServe(":3334", r)

	if err != nil {
		panic(err.Error())
	}
}
