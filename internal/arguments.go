package internal

import (
	"flag"
	"log"
)

type Arguments struct {
	Port int
	File string
}

func GetArguments() Arguments {
	port := flag.Int("port", 8000, "port")
	file := flag.String("file", "", "path to configuration file")
	flag.Parse()
	log.Printf("Port: %d\n", *port)
	log.Printf("File: %s\n", *file)
	return Arguments{Port: *port, File: *file}
}
