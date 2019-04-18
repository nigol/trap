package main

import (
	"fmt"
	"log"
	"net/http"
	"bytes"
	"net"
)

func getOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%+v", r)
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	bodyString := buf.String()
	fmt.Printf("\n----------------Body:\n%s", bodyString)
	fmt.Printf("\n!----------------!\n")
}

func printBanner(port string, path string, ip net.IP) {
	fmt.Printf("It's a trap. It grabs every incoming HTTP request and dumps it to the console.\nhttps://github.com/nigol/trap\n\n")
	fmt.Printf("* Listening on http://%+v%s%s\n\n", ip, port, path)
}

func main() {
	ip := getOutboundIP()
	path := "/trap"
	port := ":16385"
	printBanner(port, path, ip)
	http.HandleFunc(path, handler)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
