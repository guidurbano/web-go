package main

import (
	"net/http"
	"web-go/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe("localhost:8000", nil)
}
