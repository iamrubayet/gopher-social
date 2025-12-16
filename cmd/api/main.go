package main

import (
	"log"

	"github.com/iamrubayet/gopher-social/internal/db"
	"github.com/iamrubayet/gopher-social/internal/env"
	"github.com/iamrubayet/gopher-social/internal/store"
)

func main() {

	cfg := config{
		addr: env.GetString("ADDR", ":8085"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://root:password@localhost:5433/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 25),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 25),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Println("Database connection pool established")

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}
	mux := app.mount()
	log.Fatal(app.run(mux))

}
