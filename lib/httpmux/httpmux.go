package httpmux

import (
	"errors"
	"net/http"
	"sync"
)



type HTTPMux struct {
    PathHandlers map[string]http.Handler
	mu sync.Mutex
}

func (receiver *HTTPMux) ServeHTTP(respwriter http.ResponseWriter, req *http.Request) {
    // TODO
	receiver.mu.Lock()
    defer receiver.mu.Unlock()
	if receiver.PathHandlers == nil {
		receiver.PathHandlers = make(map[string]http.Handler)
	}
	if handler, ok := receiver.PathHandlers[req.URL.Path]; ok {
		handler.ServeHTTP(respwriter, req)
		return
	}
	http.NotFound(respwriter, req)
}

func (receiver *HTTPMux) HandlePath(handler http.Handler, path string) error {
    // TODO
	receiver.mu.Lock()
    defer receiver.mu.Unlock()
	if receiver.PathHandlers == nil {
		receiver.PathHandlers = make(map[string]http.Handler)
	}
	if receiver.PathHandlers[path] != nil {
		return errors.New("path already exists")
	}
	receiver.PathHandlers[path] = handler
	return nil
}
func (receiver *HTTPMux) HandlePathFunc(fn func(http.ResponseWriter, *http.Request),path string) {
    // TODO
	receiver.HandlePath(http.HandlerFunc(fn), path)
}
// func NewHTTPMux() *HTTPMux {
//     return &HTTPMux{
//         pathHandlers: make(map[string]http.Handler), //Properly initializes the map
//     }
// }