package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {
	localIP := getOutboundIP()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "nihao from %s", localIP)
	})

	fs := http.FileServer(http.Dir("static/"))
	http.ListenAndServe(":8080", nil)
}

func getOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
