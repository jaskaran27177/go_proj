package main

import (
	srvhttp "myproject/srv/http"
	"net/http"
)


func main(){
	
	srvhttp.Mux.HandleFunc("/", response)
	srvhttp.Mux.HandleFunc("/v1/sum", sum)
	port:=tcpPort()
	log("Server is running on port: ", port)
	err:=http.ListenAndServe(":"+port, srvhttp.Mux)
	if err != nil {
		log("Server got an error:", err)
	} 
}