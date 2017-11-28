package main

import (
	"flag"

	log "github.com/Sirupsen/logrus"
	_ "github.com/cloudflare/promsaint/backends" // i want it to be in global namespace?
	_ "github.com/cloudflare/promsaint/logging"
	"github.com/cloudflare/promsaint/server"
)

var (
	Version   string
	BuildTime string
)

func main() {
	flag.Parse()

	log.WithFields(log.Fields{
		"version": Version,
		"build":   BuildTime,
	}).Warn("Starting Promsaint!")
	s := server.NewPromsaint()
	s.Start()
}
