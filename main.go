package main

import (
	"context"
	"go-ent-mysql/ent"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client, err := ent.Open("mysql", "enter:password@tcp(localhost:3336)/enttest")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	if err := client.Schema.WriteTo(context.Background(), os.Stdout); err != nil {
		log.Fatalf("failed printing schema changes: %v", err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
