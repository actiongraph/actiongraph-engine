package main

import (
	"github.com/actiongraph/actiongraph-engine/app"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	// init the logger
	initLogger()

	// init the env variables
	if _, err := os.Stat("./.env"); !os.IsNotExist(err) {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		log.Info(".env file loaded successfully")
	}
	log.Info("starting service ", os.Getenv("SERVICE_NAME"), " version ", os.Getenv("SERVICE_VERSION_NUMBER"))

	app := app.NewApp()
	app.Start()
}

// initLogger .. initialize the logger
func initLogger() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}
