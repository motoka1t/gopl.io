// server1は、最小限の「echo」サーバです
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handlerは、リクエストされたURLxのPath要素を返します。
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL Path = %q\n", r.URL.Path)
}
