package main

import "net/http"


func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", response)
	port:=tcpPort()
	log("Server is running on port: ", port)
	err:=http.ListenAndServe(":"+port, mux)
	if err != nil {
		log("Server got an error:", err)
	} 
}