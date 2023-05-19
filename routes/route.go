package routes

import (
	"fmt"
	"net/http"
)


func MuxRoute(mux *http.ServeMux, method string, path string, handler http.Handler, queryStr ...string) {
	
	if len(queryStr) > 0 {
		fmt.Printf("[%s]: %s %v \n", method, path, queryStr)
	} else {
		fmt.Printf("[%s]: %s \n", method, path)
	}

	mux.Handle(path, handler)
}