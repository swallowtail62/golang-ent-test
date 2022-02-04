package main

import (
	"context"
	"fmt"
	"go-ent-mysql/ent"
	"go-ent-mysql/env"
	"go-ent-mysql/repository"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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
		log.Fatalf("failed opening connection to DB: %v", err)
	}
	defer client.Close()
	if err := client.Schema.WriteTo(context.Background(), os.Stdout); err != nil {
		log.Fatalf("failed printing schema changes: %v", err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	userRepository := repository.NewUserRepository(client.Debug())

	r := gin.Default()
	r.GET("/health-check", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	type UserCreatePayload struct {
		Name string `json:"name" binding:"required"`
		Age  int    `json:"age" binding:"required"`
	}

	userRoutes := r.Group("/user")
	userRoutes.POST("", func(c *gin.Context) {
		var payload UserCreatePayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user, err := userRepository.CreateUser(c, payload.Name, payload.Age)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"user": user})
	})

	if err = r.Run(fmt.Sprintf(":%d", env.Conf.PORT)); err != nil {
		log.Fatalf("The port %d is in use.", env.Conf.PORT)
	}
}
