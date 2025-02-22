package main

import (
	"fmt"
	srvhttp "myproject/srv/http"
	"net/http"
	"strconv"
)


func init(){
	srvhttp.Mux.HandleFunc("/", response)
	srvhttp.Mux.HandleFunc("/v1/sum", sum)
}
func response(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello World")
}

func sum(w http.ResponseWriter, r *http.Request){
	x, x_err := strconv.ParseInt(r.FormValue("x"), 10, 64)
	y, y_err:= strconv.ParseInt(r.FormValue("y"), 10, 64)
	if x_err != nil || y_err != nil {
		http.Error(w, "Invalid parameters", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Sum of %d and %d is %d", x, y, x+y)
	
	
}


