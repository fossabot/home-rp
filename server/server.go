package server

import (
	"github.com/gorilla/mux"
	"github.com/just1689/home-rp/model"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
	"log"
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
	rev := httputil.NewSingleHostReverseProxy(u)
	r.Host(route.Host).HandlerFunc(rev.ServeHTTP)

}
