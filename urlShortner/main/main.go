package main

import (
	"fmt"
	"net/http"
	"os"
	"urlShortner/urlshort"
	"path/filepath"
)

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func main() {
	mux := defaultMux()
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc": "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	yamlContent, err := os.ReadFile(filepath.Join(cwd, "paths.yaml"))
	if err != nil {
		panic(err)
	}
	yamlHandler, err := urlshort.YAMLHandler(yamlContent, mapHandler)
	if err != nil {
		panic(err)
	}

	jsonContent, err := os.ReadFile(filepath.Join(cwd, "paths.json"))
	if err != nil {
		panic(err)
	}
	jsonHandler, err := urlshort.JSONHandler(jsonContent, yamlHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :1234")
	http.ListenAndServe(":1234", jsonHandler)
}

