
package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Chao dang do, toi muon dit Minh Trang va Luong Nguyen ca Ngan Phung 1!")
}

func main() {
    http.HandleFunc("/", handler)
		fmt.Println("Start http server at port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}