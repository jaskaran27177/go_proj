package main

import (
	"fmt"
	"io"
	"os"
)
var output io.Writer
func log(a... any){
	fmt.Fprintln(output, a...)
}
func init(){
	if verbose {
		output = os.Stdout
	} else{
		output = io.Discard
	}

}