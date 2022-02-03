package repository

import (
	"context"
	"go-ent-mysql/ent"
)

type UserRepository interface {
	CreateUser(ctx context.Context, name string, age int) (*ent.User, error)
}

func NewUserRepository(client *ent.Client) UserRepository {
	return &UserRepositoryImpl{client}
}

type UserRepositoryImpl struct {
	client *ent.Client
}

func (ur UserRepositoryImpl) CreateUser(ctx context.Context, name string, age int) (*ent.User, error) {
	return ur.client.User.Create().SetName(name).SetAge(age).Save(ctx)
}
