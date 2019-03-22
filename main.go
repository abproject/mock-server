package main

import (
	"fmt"
	.  "github.com/abproject/mock-server/internal/init"
	"net/http"
)


func main() {
	parameters := GetArguments()

	var config Config
	config.Parse(parameters.File)

	var router Router
	router.Init(config)

	http.HandleFunc("/", router.Request)
	http.ListenAndServe(fmt.Sprintf(":%d", parameters.Port), nil)
}


