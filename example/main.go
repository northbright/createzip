package main

import (
	"fmt"
	"github.com/northbright/createzip"
	"io"
	"net/http"
)

type MyHandler struct{}

var (
	mux = make(map[string]func(http.ResponseWriter, *http.Request))
)

func hello(w http.ResponseWriter, r *http.Request) {
	// Create zip file for HTTP server
	zf := createzip.NewForHTTP(w, "new.zip")
	// Add file into zip file. 1st param is the real file path, 2nd param is the file name in the zip file.
	// If file name in the zip file is empty, it will put the file in the root dir of the zip.
	zf.AddFile("./README.md", "")
	zf.AddFile("./main.go", "example/main.go")
	zf.Close()
}

func (*MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := r.URL.String()
	if h, ok := mux[url]; ok {
		h(w, r)
		return
	}
	io.WriteString(w, "URL can not found: "+r.URL.String())
}

func main() {
	mux["/"] = hello

	server := http.Server{
		Addr:    ":80",
		Handler: &MyHandler{},
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("ListenAndServe: %v", err)
	}

	// Output:
}
