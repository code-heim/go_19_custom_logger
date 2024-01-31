package main

import "custom_logger/logger"

func main() {
	logger.Info("This is an info print!")
	logger.Warning("This is a warning print!")
	logger.Error("This is an error print!")

	logger.SetLevel(logger.WarningLevel)

	logger.Info("This is an info print!")
	logger.Warning("This is a warning print!")
	logger.Error("This is an error print!")

}
