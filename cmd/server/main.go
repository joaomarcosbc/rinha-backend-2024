package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joaomarcosbc/rinha-backend-2024/cmd/handler"
	application "github.com/joaomarcosbc/rinha-backend-2024/internal/application/transacao"
	infra "github.com/joaomarcosbc/rinha-backend-2024/internal/infra/transacao"
)

func main() {
	var psqlconn string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", "localhost",
		"admin", "rinha", "rinhabackenddb")

	poolConfig, err := pgxpool.ParseConfig(psqlconn)
	if err != nil {
		panic(err)
	}

	db, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}

	err = db.Ping(context.Background())
	if err != nil {
		panic(err)
	}

	defer db.Close()

	fmt.Println("Database connected")

	repo := infra.NewRepository(db)
	service := application.NewService(repo)
	handler := handler.NewTransacaoController(service)

	router := chi.NewRouter()
	router.Get("/clientes/{id}/extrato", handler.ObterExtrato)

	http.ListenAndServe(":8080", router)
}
