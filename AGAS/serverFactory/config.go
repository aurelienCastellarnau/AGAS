/**
 * File: config.go
 * Project: Angular Golang Api Server/QuickStart
 * File Created: Sunday, 26th November 2017 2:54:46 pm
 * Author: Aurélien Castellarnau (castellarnau.a@gmail.com)
 * -----
 * Last Modified: Monday, 27th November 2017 12:25:43 am
 * Modified By: Aurélien Castellarnau (castellarnau.a@gmail.com>)
 * -----
 * Copyright © 2018 - 2017 WebFace, WebFace
 */

package serverFactory

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"liveCoding/helpers"
	"os"
	"path/filepath"
)

type Config struct {
	// globale configuration
	RootDir string `json:"root_dir"`
	WebDir  string `json:"web_dir"`
	// private file management
	file     *os.File
	fileInfo os.FileInfo
	filePath string
}

var debug = helpers.Debug()

// MakeConfig
// try to parse the defined configuration point by configPath
// default root is ../, considering server is executed from
// root/bin/server.go
// web is optionnal and relative to root
// return a pointer to Config structure
func MakeConfig(configPath, root, web string) (*Config, error) {
	var err error

	if configPath == "" {
		configPath = "../config/config.json"
	}
	configPath, err = filepath.Abs(configPath)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", debug.GetStack(MakeConfig), err)
	}
	c := &Config{
		filePath: configPath,
	}
	err = c.parseConfig(root, web)
	return c, err
}

// initConfig,
// if root is empty Dir(./exec.go) become the default root dir
// if web is empty Dir(root + "public") become the default web dir
func (c *Config) initConfig(root, web string) error {
	var err error

	if root == "" {
		root = filepath.Dir(filepath.Dir(os.Args[0]))
		if err != nil {
			return fmt.Errorf("%s: %s", debug.GetStack(c.initConfig), err)
		}
	}
	if web == "" {
		web = "public/lcPartOne/dist"
	}
	web = filepath.Join(root, web)
	c.RootDir = root
	c.WebDir = web
	return nil
}

// createFile return an open stream to path
// creating all needed folders on the way.
func (c *Config) createFile(path string) (*os.File, error) {
	if path == "" {
		return nil, fmt.Errorf("%s : the provided path is empty", debug.GetStack(c.createFile))
	}
	err := os.MkdirAll(filepath.Dir(path), 0666)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", debug.GetStack(c.createFile), err)
	}
	newFile, err := os.Create(path)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", debug.GetStack(c.createFile), err)
	}
	return newFile, nil
}

func (c *Config) parseConfig(root, web string) error {
	var err error

	c.fileInfo, err = os.Stat(c.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			err = c.initConfig(root, web)
			if err != nil {
				return err
			}
			buffer, err := json.Marshal(c)
			if err != nil {
				return fmt.Errorf("%s: marshalling %s", debug.GetStack(c.parseConfig), err)
			}
			err = os.MkdirAll(filepath.Dir(c.filePath), 0666)
			if err != nil {
				return fmt.Errorf("%s: creating config folder %s", debug.GetStack(c.parseConfig), err)
			}
			err = ioutil.WriteFile(c.filePath, buffer, 0666)
			if err != nil {
				return fmt.Errorf("%s: writing to file %s", debug.GetStack(c.parseConfig), err)
			}
			return nil
		}
		return fmt.Errorf("%s: nil? %s", debug.GetStack(c.parseConfig), err)
	}
	c.file, err = os.OpenFile(c.filePath, os.O_RDONLY, 0444)
	if err != nil {
		return fmt.Errorf("%s: %s", debug.GetStack(c.parseConfig), err)
	}
	defer c.file.Close()
	buffer := make([]byte, 1024)
	bufferLen, err := c.file.Read(buffer)
	if err != nil {
		return fmt.Errorf("%s: reading config file %s", debug.GetStack(c.parseConfig), err)
	}
	rightBuffer := buffer[:bufferLen]
	err = json.Unmarshal(rightBuffer, c)
	if err != nil {
		return fmt.Errorf("%s: unmarshalling file content %s", debug.GetStack(c.parseConfig), err)
	}
	return nil
}
