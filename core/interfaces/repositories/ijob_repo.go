package repositories

import (
	e "app/core/entities"
	mds "app/core/models"
	"context"
)

type JobRepo interface {
	SearchJob(ctx context.Context, term string, pageIndex, pageAmount int) ([]e.Job, error)
	GetJob(ctx context.Context, jobId string) (e.Job, error)
	AddJob(ctx context.Context, job mds.Job) (e.Job, error)
	AddTestJob(ctx context.Context) error
	UpdateJob(ctx context.Context, job mds.Job) (e.Job, error)
	PatchJob(ctx context.Context, job mds.Job) (e.Job, error)
	DeleteJob(ctx context.Context, jobId string) (int, error)
}
