package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var gittag = "SED_GIT_TAG"
	var githash = "SED_GIT_HASH"

	fmt.Fprintf(w,"<!DOCTYPE html>")
	fmt.Fprintf(w,"<html>")
	fmt.Fprintf(w,"<head><title>Golang service</title><style>body { background-color: cyan; }</style></head>")
	fmt.Fprintf(w,"<body>")
	fmt.Fprintf(w,"<h1 style=\"text-align:center\">This is a simple Golang service!!</h1>")
	fmt.Fprintf(w,"<h1 style=\"text-align:center\">Git tag: "+ gittag +"</h1>")
	fmt.Fprintf(w,"<h1 style=\"text-align:center\">Git commit ID: "+ githash +"</h1>")
	fmt.Fprintf(w,"</body>")
	fmt.Fprintf(w,"</html>")
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8083", nil))
}