package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	KB = 1024
	MB = 1024 * KB
)

var (
	STEP     int
	MAX      int
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1 * KB,
		WriteBufferSize: 0, // overwritten by flag.Parse
	}
)

func init() {
	flag.IntVar(&STEP, "step", 10*MB, "size of datagram")
	flag.IntVar(&MAX, "max", 150*MB, "max data to send")
	flag.IntVar(&upgrader.WriteBufferSize, "wb", 1*MB, "size of write buffer")
}

func serveWS(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s\n", r.Method, r.URL)

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("ws upgrade err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	start := time.Now()
	for i := 0; i < MAX; i += STEP {
		writer, err := conn.NextWriter(websocket.TextMessage)
		if err != nil {
			log.Println("no ws writer:", err)
			return
		}
		writeBytes(writer, STEP)
		writer.Close()
	}
	log.Printf("Delivered %d bytes in %s", MAX, time.Now().Sub(start))
}

func writeBytes(w io.Writer, numBytes int) {
	data := []byte(`A`)
	for i := 0; i < numBytes; i++ {
		w.Write(data)
	}
}
