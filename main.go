package main

import (
	"os"

	"github.com/yaasin-raki2/banking/app"
	"github.com/yaasin-raki2/banking/logger"
)

//set environment variables
func setEnvs() {
	os.Setenv("SERVER_ADDRESS", "localhost")
	os.Setenv("SERVER_PORT", "8080")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "4000")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "postgres")
	os.Setenv("DB_NAME", "postgres")
}

func main() {
	setEnvs()
	logger.Info("Starting the application...")
	app.Start()
}
