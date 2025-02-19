package main

import "os"

func tcpPort() string {
	env_port:=os.Getenv("PORT")
	const default_port string = "8080"
	port:=default_port
	if env_port != "" {
		port = string(env_port)
	}
	return port

}