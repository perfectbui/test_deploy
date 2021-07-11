
package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love Minh Trang, Bao Tran and Ngan Phung!")
}

func main() {
    http.HandleFunc("/", handler)
		fmt.Println("Start http server at port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}