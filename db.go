package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/ArifulProtik/gograph-notes/config"
	"github.com/ArifulProtik/gograph-notes/ent"
	"github.com/ArifulProtik/gograph-notes/log"
	_ "github.com/lib/pq"
)

func DbClient(cfg *config.Config, logger log.Logger) *ent.Client {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s",
		cfg.DB_HOST, cfg.DB_PORT, cfg.DB_USER, cfg.DB_PASS, cfg.DB_NAME)
	client, err := ent.Open("postgres", psqlInfo)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("Database Connected")
	if err := client.Schema.Create(context.Background()); !errors.Is(err, nil) {
		logger.Fatalf("Error: failed creating schema resources %v\n", err)
	}

	defer client.Close()

	return client
}
