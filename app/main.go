package main

import (
	"os"

	"github.com/NasSilverBullet/twitter-clone-api/app/frameworks"
)

func main() {
	logger := frameworks.NewLogger(os.Stdout, os.Stderr)

	if err := frameworks.LoadEnv(logger); err != nil {
		logger.Error(err)
	}

	validator := frameworks.NewValidator()

	sqlHandler, err := frameworks.NewSQLHandler(logger)
	if err != nil {
		logger.Error(err)
	}

	if err := frameworks.Routes(logger, validator, sqlHandler); err != nil {
		logger.Error(err)
	}
}
