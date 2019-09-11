package main

import (
	"Gophercises/Ex2/urlshort"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	yamlFile := os.Args[1]

	//ServeMux is an HTTP request multiplexer.
	//It matches the URL of each incoming request against a list of registered patterns and calls
	//the handler for the pattern that most closely matches the URL.
	mux := defaultMux()

	//Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	mapHandler := urlshort.HttpHandler(urlshort.MapHandler(pathsToUrls), mux)

//		yaml := `
//- path: /urlshort
//  url: https://github.com/gophercises/urlshort
//- path: /urlshort-final
//  url: https://github.com/gophercises/urlshort/tree/solution
//`
	yaml, err := ioutil.ReadFile(yamlFile)

	if err != nil {
		log.Println(err)
	}

	yamlMapHandler, err := urlshort.YAMLHandler([]byte(yaml))
	if err != nil {
		panic(err)
	}

	//Build the YAMLHandler using the mapHandler as the fallback
	yamlHandler := urlshort.HttpHandler(yamlMapHandler, mapHandler)
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
