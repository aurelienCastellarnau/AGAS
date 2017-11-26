/**
 * File: handlers.go
 * Project: Angular Golang Api Server/QuickStart
 * File Created: Sunday, 26th November 2017 2:52:38 pm
 * Author: Aurélien Castellarnau (castellarnau.a@gmail.com)
 * -----
 * Last Modified: Monday, 27th November 2017 12:25:53 am
 * Modified By: Aurélien Castellarnau (castellarnau.a@gmail.com>)
 * -----
 * Copyright © 2018 - 2017 WebFace, WebFace
 */

package serverFactory

import (
	"fmt"
	"liveCoding/helpers"
	"log"
	"net/http"
	"path/filepath"
)

type Handler struct {
	cfg *Config
}

func InitHandlers(cfg *Config) *Handler {
	return &Handler{
		cfg: cfg,
	}
}

// HelloWorld write a hello sentence to the http stream
func (h *Handler) HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World with a simple Golang web server.")
}

// ServeFiles do what is named for using the define config.WebDir
func (h *Handler) ServeFiles(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: defined webDir: %s", helpers.Debug().GetStack(h.ServeFiles), h.cfg.WebDir)
	http.ServeFile(w, r, filepath.Join(h.cfg.WebDir, "index.html"))
}
