package rest

/*
	User REST holds pages that are statically routed without any pattern matching
	and does not require authentication.
*/

import (
	"net/http"
)

// RegisterUserRoute write/rewrite a new route with it's handler
func (rest *REST) RegisterUserRoute(route, method string, handler http.HandlerFunc) {
	defer rest.rwRouter.Unlock()
	rest.rwRouter.Lock()

	routeAndMethodStruct := routeAndMethod{
		Route:  route,
		Method: method,
	}

	rest.routing[routeAndMethodStruct] = handler
}

// SetUserRouting execute all registered routing callback
func (rest *REST) SetUserRouting() {
	defer rest.rwRouter.RUnlock()
	rest.rwRouter.RLock()

	for route, handler := range rest.routing {
		rest.mux.HandleFunc(route.Route, func(w http.ResponseWriter, r *http.Request) {
			if r.Method == "" {
				handler(w, r)
				return
			}
			if r.Method != route.Method {
				http.NotFound(w, r)
				return
			}
			if r.URL.Path != route.Route {
				http.NotFound(w, r)
				return
			}
			handler(w, r)
		})
	}
}
