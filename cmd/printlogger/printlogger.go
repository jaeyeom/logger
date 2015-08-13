// Binary printlogger is an example of a very simple logger which
// calls log.Println.
package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/jaeyeom/logger/handler"
)

var addr = flag.String(
	"addr",
	":8888",
	"logger HTTP server address",
)

func main() {
	flag.Parse()
	log.Fatal(http.ListenAndServe(*addr, handler.New(func(content string) {
		log.Println(content)
	})))
}
