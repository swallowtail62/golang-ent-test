package repository

import (
	"context"
	"go-ent-mysql/ent"
)

type TeamRepository interface {
	CreateTeam(ctx context.Context, name string) (*ent.Team, error)
}

func NewTeamRepository(client *ent.Client) TeamRepository {
	return &TeamRepositoryImpl{client}
}

type TeamRepositoryImpl struct {
	client *ent.Client
}

func (tr TeamRepositoryImpl) CreateTeam(ctx context.Context, name string) (*ent.Team, error) {
	return tr.client.Team.Create().SetName(name).Save(ctx)
}
