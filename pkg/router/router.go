package router

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"../logger"
	"../middlewares"
	"../models"
	"github.com/gorilla/mux"
)

//InitializeServer starts server
func InitializeServer(routes []models.Route) {
	mainRouter := mux.NewRouter().StrictSlash(true)
	mainRouter.Use(middlewares.LoggingMiddleware)
	// mwAuthorize and mwAuthenticate basically work the same
	mw := []func(http.Handler) http.Handler{middlewares.AuthenticationMiddleware}

	for _, router := range routes {

		if router.IsPublic == true {
			mainRouter.HandleFunc(router.Path, router.HandlerFunction).Methods(router.MethodName)
			continue
		}

		mainRouter.Handle(router.Path, middleware(http.HandlerFunc(router.HandlerFunction), mw...)).Methods(router.MethodName)

	}

	portNumber, err := strconv.Atoi(os.Getenv("PORT_NUMBER"))
	if err != nil {
		logger.Log(err.Error(), logger.WARNING)
		portNumber = 10000
	}

	var listenAddress = fmt.Sprintf("%s%d", ":", portNumber)
	logger.Log("Service Started", logger.INFO)
	error := http.ListenAndServe(listenAddress, mainRouter)
	if error != nil {
		logger.Log(error.Error(), logger.ERROR)
		os.Exit(1)
	}

}

func middleware(h http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
	for _, mw := range middleware {
		h = mw(h)
	}
	return h
}
