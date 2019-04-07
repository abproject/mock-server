package websocket

import (
	"errors"
	.  "github.com/abproject/mock-server/config"
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
	"net/http"
	"sort"
	"time"
)

type Websocket struct {
	endpoints []Endpoint
}

func NewWebsocket(config WebsocketConfig) *Websocket {
	var endpoints = make([]Endpoint, len(config.Endpoints))

	for index, endpointConfig := range config.Endpoints {
		endpoint := NewEndpoint(endpointConfig)
		endpoints[index] = *endpoint
	}
	return &Websocket{
		endpoints: endpoints,
	}
}

func (ws *Websocket) FindEndpoint(r *http.Request) (error, Endpoint) {
	for _, endpoint := range ws.endpoints {
		var path = endpoint.Url
		if path[0] != '/' {
			path = "/" + path
		}
		if path == r.RequestURI {
			return nil, endpoint
		}
	}
	return errors.New("NO WEBSOCKET FOUND"), Endpoint{}
}

func (ws *Websocket) Subscribe(w http.ResponseWriter, r *http.Request, endpoint Endpoint) {
	log.Printf("Websocket found: %#v", endpoint)
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Connection\n%s", err)
		return
	}
	ws.emitMessages(conn, endpoint, r.Header.Get("Sec-Websocket-Key"))
}

func (ws *Websocket) emitMessages(conn *websocket.Conn, endpoint Endpoint, client string) {
	for i := int64(0); i < endpoint.Repeat; i++ {
		var messages = make([]string, len(endpoint.Messages))
		copy(messages, endpoint.Messages)
		switch endpoint.Order {
		case "end":
			sort.Sort(sort.Reverse(sort.StringSlice(endpoint.Messages)))
		case "random":
			r := rand.New(rand.NewSource(time.Now().Unix()))
			for n := len(messages); n > 0; n-- {
				randIndex := r.Intn(n)
				messages[n-1], messages[randIndex] = messages[randIndex], messages[n-1]
			}
		}
		for _, message := range messages {
			if err := conn.WriteMessage(1, []byte(message)); err != nil {
				log.Println("Write:", err)
				return
			}
			log.Printf("Client: %s. Iteration: %d. Sent message: '%#v'", client, i, message)
			time.Sleep(time.Duration(endpoint.Interval) * time.Millisecond)
		}
	}
}
