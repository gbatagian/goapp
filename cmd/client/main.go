package main

import (
	"flag"
	"goapp/internal/client"
)

var n int

func init() {
	nFlag := flag.Int("n", 1, "Number of WS connections to open")
	flag.Parse()
	n = *nFlag
}

func main() {
	done := make(chan struct{})

	for i := 0; i < n; i++ {
		go func() {
			ws := client.NewWSClient("goapp/ws")
			ws.Connect()
		}()
	}

	<-done
}
