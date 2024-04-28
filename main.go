package main

import (
	"fmt"
	"log"
	"net/http"

	"capybara-go/config"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func chat() {

}

func main() {
	port := config.GlobalConfig.Server.Port
	http.HandleFunc("/", HelloWorldHandler)
	log.Println(fmt.Sprintf("Server starting on port %d ...", port))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
