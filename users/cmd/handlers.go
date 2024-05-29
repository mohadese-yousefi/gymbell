package main

import "net/http"


func RootHandler(r http.ResponseWriter, w *http.Request) {
	r.Write([]byte("This is the root page"))
}