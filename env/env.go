package env

import "os"
var Verbose bool = ("true" == os.Getenv("VERBOSE"))
