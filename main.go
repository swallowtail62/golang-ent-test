package main

import (
	"context"
	"fmt"
	"go-ent-mysql/controller"
	"go-ent-mysql/ent"
	"go-ent-mysql/env"
	"go-ent-mysql/validation"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client := openDB()
	defer client.Close()

	jaTranslator := validation.RegisterJaTranslation()

	r := gin.Default()
	r.GET("/health-check", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	controller.RegisterUserRoutes(r, client, jaTranslator)

	if err := r.Run(fmt.Sprintf(":%d", env.Conf.PORT)); err != nil {
		log.Fatalf("The port %d is in use.", env.Conf.PORT)
	}
}

// Establish connection to DataBase and migrate, then return ent client
func openDB() *ent.Client {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		env.Conf.DBUser,
		env.Conf.DBPassword,
		env.Conf.DBHost,
		env.Conf.DBPort,
		env.Conf.DBDatabase,
	)
	client, err := ent.Open("mysql", dataSource)
	if err != nil {
		log.Fatalf("failed opening connection to DB: %v", err)
	}
	if err := client.Schema.WriteTo(context.Background(), os.Stdout); err != nil {
		log.Fatalf("failed printing schema changes: %v", err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client.Debug()
}
