package client

import (
	"fmt"
	"log"
	"net/url"

	"goapp/internal/pkg/watcher"

	"github.com/gorilla/websocket"
)

type WSClient struct {
	URL url.URL
}

func NewWSClient(path string) *WSClient {
	ws := WSClient{URL: url.URL{
		Scheme: "ws",
		Host:   "0.0.0.0:8080",
		Path:   path,
	},
	}
	return &ws
}

func (ws *WSClient) Connect() {
	conn, _, err := websocket.DefaultDialer.Dial(ws.URL.String(), nil)
	if err != nil {
		log.Printf("Error connection to websocket: %v\n", err)
	}
	defer conn.Close()

	for {
		message := watcher.Counter{}
		conn.ReadJSON(&message)

		fmt.Println(message)
	}
}
