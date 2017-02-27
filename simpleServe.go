package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		requestDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			fmt.Fprintf(w, "%v", err)
		}
		fmt.Fprintf(w, string(requestDump))
	})
	http.Handle("/", http.FileServer(http.Dir("./static")))
	fmt.Println("Now serving on port 3033")
	err := http.ListenAndServe(":3033", nil)
	if err != nil {
		panic(err)
	}
}
