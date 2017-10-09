package main

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
)

func main() {
	// Show request information if the echo endpoint is requested
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		requestDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			fmt.Fprintf(w, "%v", err)
		}
		fmt.Fprintf(w, string(requestDump))
	})
	// Serve files from the folder "static"
	http.Handle("/", http.FileServer(http.Dir("./static")))

	// Get and print server's local address
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Println("Local IP unknown.")
	} else {
		localAddr := conn.LocalAddr().(*net.UDPAddr).IP
		fmt.Println("Local Address:", localAddr)
	}
	defer conn.Close()

	// Start serving
	fmt.Println("Now serving on port 3033")
	err = http.ListenAndServe(":3033", nil)
	if err != nil {
		panic(err)
	}
}
