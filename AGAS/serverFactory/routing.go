/**
 * File: routing.go
 * Project: Angular Golang Api Server/QuickStart
 * File Created: Sunday, 26th November 2017 2:52:43 pm
 * Author: Aurélien Castellarnau (castellarnau.a@gmail.com)
 * -----
 * Last Modified: Monday, 27th November 2017 12:26:03 am
 * Modified By: Aurélien Castellarnau (castellarnau.a@gmail.com>)
 * -----
 * Copyright © 2018 - 2017 WebFace, WebFace
 */

package serverFactory

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route type define what a route must contain...
type Route struct {
	Name    string
	Method  string
	Pattern string
	//TokenProtected     bool
	//HTTPBasicProtected bool
	//Reserved           bool
	HandlerFunc http.HandlerFunc
}

// GetRouting return the defined routing system
// configPath: path to the json configuration file
// set to /root/config.json by default
// cfg.RootDir: for specific root directory
// cfg.WebDir: for specific web directory
func GetRouting(cfg *Config) (*mux.Router, error) {
	//handlers := InitHandlers(cfg)
	router := mux.NewRouter().StrictSlash(true)
	AddAngularRouting(router, cfg)
	return router, nil
}
