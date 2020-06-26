package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("welcome max  system")

	http.Handle("/", http.HandlerFunc(hellowold))

	http.ListenAndServe(":1314", nil)
}

func hellowold(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HelloWorld"))
}
