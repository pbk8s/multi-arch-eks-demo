package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from NODE:%s, POD:%s, CPU PLATFORM:%s/%s",
		os.Getenv("NODE_NAME"), os.Getenv("POD_NAME"), runtime.GOOS, runtime.GOARCH)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
