package main

import (
	httpsrv "myproject/srv/http"
	"net/http"
)


func main(){
	
	
	port:=tcpPort()
	log("Server is running on port: ", port)
	err:=http.ListenAndServe(":"+port, httpsrv.Mux)
	if err != nil {
		log("Server got an error:", err)
	} 
}