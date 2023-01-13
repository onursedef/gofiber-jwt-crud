package config

import (
	"fmt"
	"path/filepath"

	"github.com/joho/godotenv"
)

var APP_ENV_SETTING = "development"

type AppEnv int

const (
	Development = 0
	Production  = 1
)

func GetAppEnv() string {
	return APP_ENV_SETTING
}

func SetAppEnv(appEnv AppEnv) {
	if appEnv == Development {
		APP_ENV_SETTING = "development"
	} else {
		APP_ENV_SETTING = "production"
	}
}

func LoadEnv() {
	// get this folder path
	// filepath.Dir("")

	path, err := filepath.Abs("")

	if err != nil {
		//
	}

	if GetAppEnv() == "development" {
		godotenv.Load(fmt.Sprintf("%s/.env.development", path))
	} else {
		godotenv.Load(fmt.Sprintf("%s/.env.production", path))
	}
}
