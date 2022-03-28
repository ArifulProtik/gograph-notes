package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/ArifulProtik/gograph-notes/config"
	"github.com/ArifulProtik/gograph-notes/ent"
	"github.com/ArifulProtik/gograph-notes/ent/migrate"
	"github.com/ArifulProtik/gograph-notes/log"
	_ "github.com/lib/pq"
)

func DbClient(cfg *config.Config, logger log.Logger) *ent.Client {
	// "host=<host> port=<port> user=<user> dbname=<database> password=<pass>"
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		cfg.DB_HOST, cfg.DB_PORT, cfg.DB_USER, cfg.DB_NAME, cfg.DB_PASS)
	client, err := ent.Open("postgres", psqlInfo)
	if err != nil {
		logger.Fatal(err)
	}
	// defer client.Close()
	logger.Info("Database Connected")
	if err := client.Schema.Create(context.Background(), migrate.WithDropIndex(true), migrate.WithDropColumn(true)); !errors.Is(err, nil) {
		logger.Fatalf("Error: failed creating schema resources %v\n", err)
	}

	return client
}
