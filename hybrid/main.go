// +build !appengine
package main

import (
	"net/http"

	_ "github.com/andrii-stasiuk/go-cloud/hybrid/hybridapplib"
)

func main() {
	http.ListenAndServe("localhost:8080", nil)
}
