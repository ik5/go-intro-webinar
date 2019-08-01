package rest

import (
	"context"
	"net/http"
	"sync"
)

type routeAndMethod struct {
	Route  string
	Method string
}

// REST holds content to control the rest server
type REST struct {
	address string
	port    uint16

	routing  map[routeAndMethod]http.HandlerFunc
	rwRouter *sync.RWMutex

	mux        *http.ServeMux
	ctx        context.Context
	cancelFunc context.CancelFunc
	srv        *http.Server
}
