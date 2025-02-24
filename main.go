package main

import (
	httpsvr "myproject/srv/http"
	"net/http"
)


func main(){
	
	
	port:=tcpPort()
	log("Server is running on port: ", port)
	err:=http.ListenAndServe(":"+port, httpsvr.Mux)
	if err != nil {
		log("Server got an error:", err)
	} 
}