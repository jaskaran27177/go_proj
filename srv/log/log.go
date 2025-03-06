package logsrv

import (
	"io"
	env "myproject/env"
	"os"

	log "github.com/jaskaran27177/go-log"
)
var Logger log.Logger
func init() {
	if env.Verbose {
		Logger=log.CreateLogger(os.Stdout)
	}else{
		Logger=log.CreateLogger(io.Discard)
	}

}