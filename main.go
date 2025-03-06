package main

import (
	httpsrv "myproject/srv/http"
	logsrv "myproject/srv/log"
	"net/http"
)


func main(){
	
	
	port:=tcpPort()
	logsrv.Logger.Log("Server is running on port: ", port)
	err:=http.ListenAndServe(":"+port, &httpsrv.Mux)
	if err != nil {
		logsrv.Logger.Log("Server got an error:", err)
	} 
}