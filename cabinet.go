package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type cabinetConfig struct {
	Mame   string   `json:"mame"`
	Romset []string `json:"romset"`
}

func getExecutableDir() string {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}

func getAppBundleDir() string {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	exPath := filepath.Dir(ex)
	for exPath != "/" {
		if filepath.Ext(exPath) == ".app" {
			exPath = filepath.Dir(exPath)
			return exPath
		}
		exPath = filepath.Dir(exPath)
	}
	return ""
}

func checkFile(dir string, file string) bool {

	path := filepath.Join(dir, file)
	_, err := os.Stat(path)
	if err == nil {
		log.Println("checked for configuration, found it!", path)
		return true
	} else {
		log.Println("checked for configuration, not found", path)
		return false
	}
}

func cabinetConfiguration() *cabinetConfig {
	const CONFIG_FILE = "cabinet.json"
	var directory = "."

	if !checkFile(directory, CONFIG_FILE) {
		directory = getExecutableDir()
		if !checkFile(directory, CONFIG_FILE) {
			directory = getAppBundleDir()
			if !checkFile(directory, CONFIG_FILE) {
				log.Fatal("can't find configuration", CONFIG_FILE)
			}
		}
	}

	f, err := os.Open(filepath.Join(directory, CONFIG_FILE))
	if err != nil {
		log.Fatal("error opening configuration", err)
	}
	defer f.Close()

	c := &cabinetConfig{}

	err = json.NewDecoder(f).Decode(&c)
	if err != nil {
		log.Fatal("error decoding cabinet.json", err)
	}

	return c
}
