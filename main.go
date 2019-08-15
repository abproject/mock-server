package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/abproject/mock-server/internal/parser"
	"github.com/abproject/mock-server/internal/rest"
	"github.com/abproject/mock-server/internal/router"
)

func main() {
	logger := log.New(os.Stdout, "mock ", log.LstdFlags|log.Lshortfile)
	port := flag.Int("port", 8000, "port")
	file := flag.String("file", "", "path to configuration file")
	flag.Parse()
	logger.Printf("Port: %d\n", *port)
	logger.Printf("File: %s\n", *file)

	restStorage := rest.MakeStorage()

	if *file != "" {
		parserContext := parser.Context{
			Logger:      logger,
			RestStorage: &restStorage,
		}
		parser := parser.New(parserContext)
		parser.Parse(*file)
	}

	routerContext := router.Context{
		Logger:      logger,
		RestStorage: &restStorage,
	}
	router := router.New(routerContext)
	http.HandleFunc("/", router.Route)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		logger.Fatalf("Server failed to start: %v", err)
	}
}
