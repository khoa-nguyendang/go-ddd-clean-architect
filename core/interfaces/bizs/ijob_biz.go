package bizs

import (
	mds "app/core/models"
	"context"
)

type JobBiz interface {
	SearchJob(ctx context.Context, index, term string, pageIndex, pageAmount int) (mds.BaseReponse[mds.Job], error)
	SearchJobDatabase(ctx context.Context, term string, pageIndex, pageAmount int) (mds.BaseReponse[mds.Job], error)
	GetJob(ctx context.Context, jobId string) (mds.BaseReponse[mds.Job], error)
	AddJob(ctx context.Context, job interface{}) (mds.BaseReponse[mds.Job], error)
	UpdateJob(ctx context.Context, job interface{}) (mds.BaseReponse[mds.Job], error)
	PatchJob(ctx context.Context, job interface{}) (mds.BaseReponse[mds.Job], error)
	DeleteJob(ctx context.Context, jobId string) (mds.BaseReponse[mds.Job], error)
	CreateIndex(ctx context.Context, index string) error
	PushDocuments(ctx context.Context, index string, takes, skips int) (success, fail int, err error)
}
