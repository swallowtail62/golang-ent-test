package main

import (
	"context"
	"fmt"
	"go-ent-mysql/ent"
	"go-ent-mysql/env"
	"go-ent-mysql/repository"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		env.Conf.DBUser,
		env.Conf.DBPassword,
		env.Conf.DBHost,
		env.Conf.DBPort,
		env.Conf.DBDatabase,
	)
	client, err := ent.Open("mysql", dataSource)
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

	userRepository := repository.NewUserRepository(client.Debug())
	user, err := userRepository.CreateUser(context.Background(), "test user", 18)
	if err != nil {
		log.Fatalf("failed to create user: %v", err)
	}
	log.Printf("Succeeded creating user: %v", user)
}
