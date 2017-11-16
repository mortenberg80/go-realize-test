package main

import (
	"github.com/mortenberg80/go-realize-test/pkg/server"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Infof("Starting go-realize-test")

	server.Start()
}
