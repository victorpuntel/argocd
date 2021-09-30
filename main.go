package main

import "net/http"

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello Dev!!!"))
	})
	http.ListenAndServe(":8080", nil)
}
