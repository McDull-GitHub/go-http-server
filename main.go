package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var address string
var dir string

func init() {
	flag.StringVar(&address, "l", ":80", "listen address")
	flag.StringVar(&dir, "d", ".", "web root dir")
}

func main() {
	flag.Parse()
	fmt.Println("Serving HTTP on", address)
	log.Println(http.ListenAndServe(address, http.FileServer(http.Dir(dir))))
}