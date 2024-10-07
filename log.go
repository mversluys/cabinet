package main

import (
	"log"
	"os"
	"path/filepath"
)

func init() {
	logDir := filepath.Join(os.Getenv("HOME"), "Library", "Logs", "Cabinet")
	os.MkdirAll(logDir, 0755)
	logFile := filepath.Join(logDir, "app.log")

	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	os.Stdout = file
	os.Stderr = file
	log.SetOutput(file)
}
