package main

import (
	"log"
	"net/http"
	"os"
	"testing"
)

var app App

func TestMain(m *testing.M) {
	log.Print("initialising...")

	app = App{
		config:     Config{},
		httpClient: &http.Client{},
	}

	configLocation := os.Getenv("CONFIG_FILE")
	if configLocation == "" {
		configLocation = "config_example.json"
	}

	dbLocation := os.Getenv("DB_FILE")
	if dbLocation == "" {
		dbLocation = "users_test.db"
	}

	app.LoadConfig(configLocation)
	app.ConnectDB(dbLocation)

	log.Print("initialised.")

	os.Exit(m.Run())

	app.db.Close()
}