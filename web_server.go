package main

import (
	"fmt"
	"net/http"
)



func response(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello World")
}


