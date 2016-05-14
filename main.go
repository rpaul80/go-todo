package main

import (
	"fmt"
	"net/http"
	"runtime"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world, I'm running on %s with an %s CPU again", runtime.GOOS, runtime.GOARCH)

}

func main() {
        fmt.Println("Hello world")
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}
