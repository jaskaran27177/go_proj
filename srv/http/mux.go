package httpsrv

import (
	httpmux "myproject/lib/httpmux"
	"net/http"
)

var Mux = &httpmux.HTTPMux{
	        PathHandlers: make(map[string]http.Handler), //Properly initializes the map
	    }