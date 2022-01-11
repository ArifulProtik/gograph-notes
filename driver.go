package main

import (
	"fmt"

	"github.com/ArifulProtik/gograph-notes/config"
	"github.com/ArifulProtik/gograph-notes/log"
	"github.com/ArifulProtik/gograph-notes/server"
)

func main() {
	cfg, err := config.LoadConfig("./", "app", "env")
	if err != nil {
		fmt.Println(err)
	}
	appLogger := log.NewApiLogger(cfg)
	appLogger.InitLogger()
	appLogger.Infof("App: %s Version: %s Status: %s", cfg.App, cfg.Version, cfg.Status)
	dbclient := DbClient(cfg, appLogger)

	s := server.NewServer(cfg, appLogger, dbclient)
	s.Run()

}
