package main

import (
	"io"
	"log"
	"net/http"

	"github.com/cowaar/weather-app/pkg"
	"github.com/cowaar/weather-app/util"
)

func main() {
	// Hello world, the web server

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	weatherHandler := func(writer http.ResponseWriter, req *http.Request){
		io.WriteString(writer, pkg.GetWeather("709966"))
	}

	pokemonHandler := func(writer http.ResponseWriter, req *http.Request){
	}

	log.Println(util.ReadConfig())

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/weather", weatherHandler)
	http.HandleFunc("/pokemon", pokemonHandler)
    log.Println("Listing for requests at http://localhost:8000/hello")
	log.Fatal(http.ListenAndServe(":8000", nil))
}