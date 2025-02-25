package www

import (
	"fmt"
	httpsrv "myproject/srv/http"
	"net/http"
)
func init(){
	httpsrv.Mux.HandleFunc("/", servehttp)
}

func servehttp(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello World")
}