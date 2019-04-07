package main

import (
	"flag"
	"fmt"
	.  "github.com/abproject/mock-server/config"
	.  "github.com/abproject/mock-server/router"
	"log"
	"net/http"
)

func main() {
	port := flag.Int("port", 8000, "port")
	file := flag.String("file", "", "path to configuration file")
	flag.Parse()
	log.Printf("Port: %d\n", *port)
	log.Printf("File: %s\n", *file)

	config := ParseConfig(*file)
	router := NewRouter(*config)

	http.HandleFunc("/", router.Request)
	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}


