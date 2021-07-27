package main

import (
	"github.com/yaasin-raki2/banking/app"
	"github.com/yaasin-raki2/banking/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}
