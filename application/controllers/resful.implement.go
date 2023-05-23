package controllers

import (
	mds "app/core/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping implements AppServer
func (s *RestfulServer) Ping(ctx context.Context, jobId string) (mds.BaseReponse[mds.Job], error) {
	return s.jb.GetJob(ctx, jobId)
}

// AddJob implements AppServer
func (s *RestfulServer) AddJob(ctx context.Context, job mds.Job) (mds.BaseReponse[mds.Job], error) {
	return s.jb.AddJob(ctx, job)
}

// AddJob implements AppServer
func (s *RestfulServer) AddTestJob(ctx context.Context) error {
	return s.jb.AddTestJob(ctx)
}

// DeleteJob implements AppServer
func (s *RestfulServer) DeleteJob(ctx context.Context, jobId string) (int, error) {
	return s.jb.DeleteJob(ctx, jobId)
}

// SearchJob implements AppServer
func (s *RestfulServer) SearchJob(ctx context.Context, index, term string, pageIndex int, pageAmount int) (mds.BaseReponse[mds.Job], error) {
	return s.jb.SearchJob(ctx, index, term, pageIndex, pageAmount)
}

// SearchJob implements AppServer
func (s *RestfulServer) SearchJobDatabase(ctx context.Context, term string, pageIndex int, pageAmount int) (mds.BaseReponse[mds.Job], error) {
	return s.jb.SearchJobDatabase(ctx, term, pageIndex, pageAmount)
}

// UpdateJob implements AppServer
func (s *RestfulServer) UpdateJob(ctx context.Context, job mds.Job) (mds.BaseReponse[mds.Job], error) {
	return s.jb.UpdateJob(ctx, job)
}

// GetJob implements AppServer
func (s *RestfulServer) GetJob(ctx context.Context, jobId string) (mds.BaseReponse[mds.Job], error) {
	return s.jb.GetJob(ctx, jobId)
}

// PatchJob implements AppServer
func (s *RestfulServer) PatchJob(ctx context.Context, job mds.Job) (mds.BaseReponse[mds.Job], error) {
	return s.jb.PatchJob(ctx, job)
}

func (s *RestfulServer) CreateIndex(ctx context.Context, index string) error {
	return s.jb.CreateIndex(ctx, index)
}

func (s *RestfulServer) PushDocuments(ctx context.Context, index string, takes, skips int) (success, fail int, err error) {
	return s.jb.PushDocuments(ctx, index, takes, skips)
}

func (s *RestfulServer) ping(c *gin.Context) {
	c.String(http.StatusOK, "%s", "pong")
}
