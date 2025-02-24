package www

import (
	"fmt"
	httpsvr "myproject/srv/http"
	"net/http"
)
func init(){
	httpsvr.Mux.HandleFunc("/", servehttp)
}

func servehttp(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello World")
}