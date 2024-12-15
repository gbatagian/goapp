package client

import (
	"fmt"
	"log"
	"net/url"

	"goapp/internal/pkg/watcher"

	"github.com/gorilla/websocket"
)

type WSClient struct {
	URL      url.URL
	clientID int
}

func NewWSClient(path string, id int) *WSClient {
	ws := WSClient{
		URL: url.URL{
			Scheme: "ws",
			Host:   "localhost:8080",
			Path:   path,
		},
		clientID: id,
	}
	return &ws
}

func (ws *WSClient) Connect() error {
	conn, _, err := websocket.DefaultDialer.Dial(ws.URL.String(), nil)

	if err != nil {
		log.Printf("Error connection to websocket: %v\n", err)
		return err
	}
	defer conn.Close()

	for {
		message := watcher.Counter{}
		conn.ReadJSON(&message)

		fmt.Printf(
			"[conn #%d] iteration: %d, value: %s\n",
			ws.clientID,
			message.Iteration,
			message.Value,
		)
	}
}
