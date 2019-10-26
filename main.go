package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/abproject/mock-server/internal/file"
	"github.com/abproject/mock-server/internal/parser"
	"github.com/abproject/mock-server/internal/rest"
	"github.com/abproject/mock-server/internal/router"
)

func main() {
	logger := log.New(os.Stdout, "mock ", log.LstdFlags|log.Lshortfile)
	port := flag.Int("port", 8000, "port")
	config := flag.String("file", "", "path to configuration file")
	flag.Parse()
	logger.Printf("Port: %d\n", *port)
	logger.Printf("File: %s\n", *config)

	restStorage := rest.MakeStorage()
	fileStorage := file.MakeStorage()

	path, _ := filepath.Abs(".")
	if *config != "" {
		parserContext := models.AppContext{
			Logger:      logger,
			RestStorage: &restStorage,
			FileStorage: &fileStorage,
			Path:        path,
		}
		parser := parser.New(parserContext)
		parser.Parse(*config)
	}

	routerContext := models.AppContext{
		Logger:      logger,
		RestStorage: &restStorage,
		FileStorage: &fileStorage,
	}
	router := router.New(routerContext)
	http.HandleFunc("/", router.Route)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		logger.Fatalf("Server failed to start: %v", err)
	}
}
