package main

import (
    "fmt"
    "net/http"
)

func main() {
    helloServer()
}

func initial(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "init")
}

func halo(w http.ResponseWriter, req *http.Request){
	fmt.Fprint(w, "halo~!")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func helloServer() {
	http.HandleFunc("/", initial)
	http.HandleFunc("/halo", halo)
	http.HandleFunc("/headers", headers)
    http.ListenAndServe("localhost:4000", nil)
}
