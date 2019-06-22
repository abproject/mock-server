package main

import (
	"flag"
	"fmt"
	. "github.com/abproject/mock-server/internal_/config"
	"github.com/abproject/mock-server/internal_/rest"
	. "github.com/abproject/mock-server/internal_/router"
	"github.com/abproject/mock-server/internal_/websocket"
	"log"
	"net/http"
)

func main() {
	port := flag.Int("port", 8000, "port")
	file := flag.String("file", "", "path to configuration file")
	flag.Parse()
	log.Printf("Port: %d\n", *port)
	log.Printf("file: %s\n", *file)

	config := ParseConfig(*file)
	rest.FileRest(config.Rest)
	websocket.FileWebsocket(config.Websocket)

	http.HandleFunc("/", Router)
	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}
