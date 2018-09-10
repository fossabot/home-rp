package server

import (
	"github.com/gorilla/mux"
	"github.com/just1689/home-rp/model"
	"github.com/koding/websocketproxy"
	"log"
	"net/http"
	"net/url"
	"time"
)

func SetupServer(listenAddr string, routes []model.Route) *http.Server {

	if len(routes) == 0 {
		log.Fatalln(".. 0 rules found. Need at least one rule")
	}

	r := mux.NewRouter()
	for _, route := range routes {
		reverseProxyRouteHandler(r, route)
	}

	srv := &http.Server{
		Handler:      r,
		Addr:         listenAddr,
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  60 * time.Second,
	}
	return srv

}

func reverseProxyRouteHandler(r *mux.Router, route model.Route) {
	u, _ := url.Parse(route.RawUrl)
	rev := websocketproxy.NewProxy(u)
	r.Host(route.Host).HandlerFunc(rev.ServeHTTP)

}
