/**
 * File: debuggingTools.go
 * Project: LiveCoding/Etna/basic web server
 * File Created: Saturday, 25th November 2017 6:32:19 pm
 * Author: Aurélien Castellarnau (castellarnau.a@gmail.com)
 * -----
 * Last Modified: Saturday, 25th November 2017 6:32:22 pm
 * Modified By: Aurélien Castellarnau (castellarnau.a@gmail.com>)
 * -----
 * Copyright © 2018 - 2017 WebFace, WebFace
 */

package helpers

import (
	"reflect"
	"runtime"
)

// DebugBox regroup all development utilitaries
type DebugBox struct{}

// GetFunctionName return the package.name of the passed function
func (b *DebugBox) GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func (b *DebugBox) GetStack(i interface{}) string {
	return "[" + b.GetFunctionName(i) + "] "
}

var d = &DebugBox{}

func Debug() *DebugBox {
	return d
}
