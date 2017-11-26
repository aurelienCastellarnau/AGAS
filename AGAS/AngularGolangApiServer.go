/**
 * File: AngularGolangApiServer.go
 * Project: Angular Golang Api Server/QuickStart
 * File Created: Sunday, 26th November 2017 2:49:46 pm
 * Author: Aurélien Castellarnau (castellarnau.a@gmail.com)
 * -----
 * Last Modified: Monday, 27th November 2017 12:26:50 am
 * Modified By: Aurélien Castellarnau (castellarnau.a@gmail.com>)
 * -----
 * Copyright © 2018 - 2017 WebFace, WebFace
 */

package main

import (
	bws "AGAS/AGAS/basicWebServer"
	hlp "AGAS/AGAS/helpers"
	sf "AGAS/AGAS/serverFactory"
	"flag"
	"log"
	"net/http"
)

func main() {
	debug := hlp.Debug()
	opts := &bws.BasicServerOpts{}
	flag.IntVar(&opts.Port, "port", 80, "define the server's port")
	flag.StringVar(&opts.IP, "ip", "127.0.0.1", "define server's ip")
	root := flag.String("root", "", "allow to define a specific root point for app file system")
	web := flag.String("web", "", "allow to define a specific web directory")
	configPath := flag.String("config", "", "allow to define a specific path for json server configuration")
	flag.Parse()

	cfg, err := sf.MakeConfig(*configPath, *root, *web)
	if err != nil {
		log.Fatal(err)
	}
	routing, err := sf.GetRouting(cfg)
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", routing)
	err = http.ListenAndServe(opts.GetAddress(), nil)
	if err != nil {
		log.Printf("%s: %s", debug.GetStack(main), err)
	}
}
