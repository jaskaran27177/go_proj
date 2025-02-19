package main

import "os"
var verbose bool = ("true" == os.Getenv("VERBOSE"))
