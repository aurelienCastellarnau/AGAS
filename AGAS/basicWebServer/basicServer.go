/**
 * File: basicServer.go
 * Project: LiveCoding/Etna/basic web server
 * File Created: Saturday, 25th November 2017 3:01:07 pm
 * Author: Aurélien Castellarnau (castellarnau.a@gmail.com)
 * -----
 * Last Modified: Saturday, 25th November 2017 6:02:17 pm
 * Modified By: Aurélien Castellarnau (castellarnau.a@gmail.com>)
 * -----
 * Copyright © 2018 - 2017 WebFace, WebFace
 */

package basicWebServer

import (
	"AGAS/AGAS/helpers"
	"strconv"
)

// Debug allow developper to use all debugging tools in helpers package
var Debug = &helpers.DebugBox{}

// BasicServerOpts define the server configuration
type BasicServerOpts struct {
	address string
	Port    int
	IP      string
}

// GetAddress concat ip and port and affect to address if needed
// else default address is define to 127.0.0.1:80
func (o *BasicServerOpts) GetAddress() string {
	if o.address == "" {
		o.address = "127.0.0.1"
		if o.IP != "" {
			o.address = o.IP
		}
		port := ":80"
		if o.Port != 0 {
			port = ":" + strconv.Itoa(o.Port)
		}
		o.address += port
	}
	return o.address
}
