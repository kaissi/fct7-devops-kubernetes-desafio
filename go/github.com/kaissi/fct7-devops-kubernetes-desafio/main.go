package main

import (
	"fmt"
	"log"
	"net/http"
)

func greeting(msg string) string {
	return fmt.Sprintf("<b>%s</b>", msg)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {	
		fmt.Fprintf(w, "%s", greeting("Code.education Rocks!"))
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}