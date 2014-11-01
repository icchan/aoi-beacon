package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

// A simple web app which response with a message you set at startup
func main() {
	message := flag.String("msg", "Hello! I'm Alive", "A cute message to broadcast.")
	listen := flag.String("listen", ":8080", "Address to listen to")
	flag.Parse()

	http.HandleFunc("/", makeHandler(*message))

	log.Println("listening at", *listen)
	http.ListenAndServe(*listen, nil)
}

func makeHandler(msg string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s\n", msg) // just return the message
	}
}
