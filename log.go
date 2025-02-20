package main

import (
	"fmt"
	"io"
	"myproject/env"
	"os"
)
var output io.Writer
func log(a... any){
	fmt.Fprintln(output, a...)
}
func init(){
	if env.Verbose {
		output = os.Stdout
	} else{
		output = io.Discard
	}

}