package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	root, _ := os.Getwd()
	port := ":9999"
	prefix := "/"
	log.Print(root)

	http.Handle(prefix, http.StripPrefix(prefix, http.FileServer(http.Dir(root))))

	mux := http.DefaultServeMux.ServeHTTP
	logger := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print(r.RemoteAddr + " " + r.Method + " " + r.URL.String())
		mux(w, r)
	})

	var err error
	err = http.ListenAndServe(port, logger)
	if err != nil {
		log.Fatalln(err)
	}
}
