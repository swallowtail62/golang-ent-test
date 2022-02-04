package repository

import (
	"context"
	"go-ent-mysql/ent"
)

type UserUpdatePayload struct {
	ID   int
	Name string
	Age  int
}

type UserRepository interface {
	CreateUser(ctx context.Context, name string, age int) (*ent.User, error)
	FindAll(ctx context.Context) (ent.Users, error)
	FindByID(ctx context.Context, userID int) (*ent.User, error)
	Update(ctx context.Context, payload *UserUpdatePayload) (*ent.User, error)
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

func (ur UserRepositoryImpl) FindAll(ctx context.Context) (ent.Users, error) {
	return ur.client.User.Query().All(ctx)
}

func (ur UserRepositoryImpl) FindByID(ctx context.Context, userID int) (*ent.User, error) {
	return ur.client.User.Get(ctx, userID)
}

func (ur UserRepositoryImpl) Update(ctx context.Context, payload *UserUpdatePayload) (*ent.User, error) {
	builder := ur.client.User.UpdateOneID(payload.ID)
	if len(payload.Name) > 0 {
		builder.SetName(payload.Name)
	}
	if payload.Age > 0 {
		builder.SetAge(payload.Age)
	}
	return builder.Save(ctx)
}
