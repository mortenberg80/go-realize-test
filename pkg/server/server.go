package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mortenberg80/go-realize-test/pkg/config"
	"github.com/mortenberg80/go-realize-test/pkg/controller"

	log "github.com/sirupsen/logrus"
)

var router *mux.Router

func Start() {
	router = handlers()
	log.Infof("Listening on port %d...", config.App.Port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", config.App.Port), router)
	log.Fatal(err.Error())
}

func handlers() *mux.Router {
	r := mux.NewRouter()
	r.Handle("/helloworld", http.HandlerFunc(controller.HelloWorld)).Methods("GET")

	return r
}
