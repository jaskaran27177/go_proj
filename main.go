package main

import (
	httpsrv "myproject/srv/http"
	"net/http"

	logger "github.com/jaskaran27177/go-log"
)


func main(){
	
	
	port:=tcpPort()
	logger.Log("Server is running on port: ", port)
	err:=http.ListenAndServe(":"+port, &httpsrv.Mux)
	if err != nil {
		logger.Log("Server got an error:", err)
	} 
}