// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package repository

import (
	"context"
)

type Querier interface {
	CreateActivity(ctx context.Context, arg CreateActivityParams) (Activity, error)
	CreateProject(ctx context.Context, name string) (Project, error)
	DeleteActivity(ctx context.Context, id int32) error
	GetProjects(ctx context.Context) ([]Project, error)
	GetUser(ctx context.Context, id int32) (User, error)
	UpdateActivity(ctx context.Context, arg UpdateActivityParams) (Activity, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
