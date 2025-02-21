package datastore

import (
	"backend/ent/gen"
	"backend/ent/gen/migrate"
	"context"
	"database/sql"
	"fmt"
	"log"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type EntDB struct {
	gen.Client
}

func NewDatabase(host, port, user, name, password string) (*EntDB, error) {
	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		user,
		password,
		host,
		port,
		name)

	db, err := sql.Open("pgx", dataSourceName)
	if err != nil {
		log.Fatalf("Failed opening connection to postgres: %v", err)
	}

	drv := entsql.OpenDB(dialect.Postgres, db)
	client := gen.NewClient(gen.Driver(drv))

	if err := client.Schema.Create(context.TODO()); err != nil {
		defer client.Close()
		log.Fatalf("Failed creating schema resources: %v", err)
	}

	return &EntDB{*client}, nil
}

func (D EntDB) GenerateSchema() error {
	err := D.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		return err
	}

	return nil
}
