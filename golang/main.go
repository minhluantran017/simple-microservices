// +build ignore

package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var githash = "${GIT_HASH}"
	var background = "blue"
	if ((githash % 2) == 1) {
		background = "green"
	}
	fmt.Fprintf(w,"<!DOCTYPE html>")
	fmt.Fprintf(w,"<html>")
	fmt.Fprintf(w,"<head><title>Sample NodeJS service</title><style>body { background-color: "+ background +"; }</style></head>")
	fmt.Fprintf(w,"<body>")
	fmt.Fprintf(w,"<h1 style=\"text-align:center\">This is a simple NodeJS services!!</h1>")
	fmt.Fprintf(w,"<h1 style=\"text-align:center\">Git commit ID:"+ githash +"</h1>")
	fmt.Fprintf(w,"</body>")
	fmt.Fprintf(w,"</html>")
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8081", nil))
}