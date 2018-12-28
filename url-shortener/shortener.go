package main

import (
	"fmt"
	"net/http"
)

func mapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return nil
}

func yamlHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	return nil, nil
}

func main() {
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	m := mapHandler(pathsToUrls, mux)

	yaml :=
		`
		- path: /urlshort
		url: https://github.com/gophercises/urlshort
		- path: /urlshort-final
		url: https://github.com/gophercises/urlshort/tree/solution
		`
	y, err := yamlHandler([]byte(yaml), m)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", y)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
