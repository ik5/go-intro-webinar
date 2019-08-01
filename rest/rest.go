package rest

import (
	"context"
	"fmt"
	"net/http"
	"sync"
)

// InitREST initialize a new REST service
func InitREST(address string, port uint16) *REST {
	ctx, cancelFunc := context.WithCancel(context.Background())
	return &REST{
		address:    address,
		port:       port,
		routing:    make(map[routeAndMethod]http.HandlerFunc),
		rwRouter:   &sync.RWMutex{},
		mux:        http.NewServeMux(),
		ctx:        ctx,
		cancelFunc: cancelFunc,
		srv:        &http.Server{},
	}
}

// GetContext returns the REST context for external usage
func (rest *REST) GetContext() context.Context {
	return rest.ctx
}

// Stop the HTTP server
func (rest *REST) Stop() error {
	rest.cancelFunc()
	return rest.srv.Shutdown(rest.ctx)
}

// Serve an HTTP server
func (rest *REST) Serve() error {
	rest.srv.Addr = fmt.Sprintf("%s:%d", rest.address, rest.port)
	rest.srv.Handler = rest.mux
	return rest.srv.ListenAndServe()
}

// ServeTLS start a TLS server
func (rest *REST) ServeTLS(certFile, keyFile string) error {
	rest.srv.Addr = fmt.Sprintf("%s:%d", rest.address, rest.port)
	rest.srv.Handler = rest.mux
	return rest.srv.ListenAndServeTLS(certFile, keyFile)
}
