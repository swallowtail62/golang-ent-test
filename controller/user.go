package controller

import (
	"go-ent-mysql/ent"
	"go-ent-mysql/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type userPathParams struct {
	UserID int `uri:"userID" binding:"required"`
}

func RegisterUserRoutes(router gin.IRouter, dbClient *ent.Client, translator ut.Translator) {
	r := router.Group("/user")
	userRepository := repository.NewUserRepository(dbClient)

	r.GET("", func(c *gin.Context) {
		users, err := userRepository.FindAll(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"users": users})
	})
	r.GET("/:userID", func(c *gin.Context) {
		var params userPathParams
		if err := c.ShouldBindUri(&params); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user, err := userRepository.FindByID(c, params.UserID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"user": user})
	})
	r.PATCH("/:userID", func(c *gin.Context) {
		var params userPathParams
		if err := c.ShouldBindUri(&params); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		payload := struct {
			Name string `json:"name" binding:"required" ja:"名前"`
			Age  int    `json:"age" binding:"required_without=Name" ja:"年齢"`
		}{}
		if err := c.ShouldBindJSON(&payload); err != nil {
			if e, ok := err.(validator.ValidationErrors); ok {
				c.JSON(http.StatusBadRequest, gin.H{"error": e.Translate(translator)})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			return
		}
		user, err := userRepository.Update(c, &repository.UserUpdatePayload{
			ID:   params.UserID,
			Name: payload.Name,
			Age:  payload.Age,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"user": user})
	})
	r.POST("", func(c *gin.Context) {
		type UserCreatePayload struct {
			Name string `json:"name" binding:"required"`
			Age  int    `json:"age" binding:"required"`
		}
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
}
