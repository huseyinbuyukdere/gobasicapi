package main

import (
	"os"
	"time"

	"../pkg/api"
	"../pkg/logger"
	"../pkg/models"
	"../pkg/router"
	"github.com/joho/godotenv"
)

//LogFile Name
var logFilePath = "Log-" + time.Now().Format("01-02-2006") + ".log"

//Routes Which Will Be Registered
var routes = []models.Route{
	models.Route{HandlerFunction: api.Login, Path: "/login", MethodName: "POST", IsPublic: false},
}

func init() {
	logger.InitializeLogger(logFilePath)
	err := godotenv.Load()
	if err != nil {
		logger.Log(err.Error(), logger.ERROR)
		os.Exit(1)
	}
	router.InitializeServer(routes)
}

func main() {
	logger.Log("Started App", logger.INFO)
}
