package main

import (
	logger "myproject/lib/log"
	httpsrv "myproject/srv/http"
	"net/http"
)


func main(){
	
	
	port:=tcpPort()
	logger.DefaultLogger.Log("Server is running on port: ", port)
	err:=http.ListenAndServe(":"+port, &httpsrv.Mux)
	if err != nil {
		logger.DefaultLogger.Log("Server got an error:", err)
	} 
}