/**
 * File: angularIoRouting.go
 * Project: Angular Golang Api Server/QuickStart
 * File Created: Sunday, 26th November 2017 9:33:42 pm
 * Author: Aurélien Castellarnau (castellarnau.a@gmail.com)
 * -----
 * Last Modified: Monday, 27th November 2017 12:25:31 am
 * Modified By: Aurélien Castellarnau (castellarnau.a@gmail.com>)
 * -----
 * Copyright © 2018 - 2017 WebFace, WebFace
 */

package serverFactory

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
)

func getAngularRouting() []string {
	return []string{
		"/home",
	}
}

func getWebAppRouting(config *Config) []Route {
	var routes []Route
	routes = append(routes)
	for _, pattern := range getAngularRouting() {
		routes = append(routes, Route{
			Pattern: pattern,
		})
	}
	return routes
}

func AddAngularRouting(router *mux.Router, config *Config) {
	home := &Route{
		Pattern: "/",
		HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
			log.Print("Serving Angular routing...")
			http.ServeFile(w, r, filepath.Join(config.WebDir, "index.html"))
		},
	}
	routes := getWebAppRouting(config)
	for _, route := range routes {
		router.PathPrefix(route.Pattern).Handler(http.StripPrefix(route.Pattern, home.HandlerFunc))
	}
	router.PathPrefix(home.Pattern).Handler(http.FileServer(http.Dir(config.WebDir)))
}
