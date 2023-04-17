package caching

import "app/core/models"

type JobCache interface {
	OpenSearchCache[models.Job]
	// Insert(ctx context.Context, job m.Job) error
	// Update(ctx context.Context, job m.Job) error
	// Delete(ctx context.Context, id string) error
	// FindOne(ctx context.Context, id string) (m.Job, error)
	// FindSimilar(ctx context.Context, term string, pageIndex, pageSize int) ([]m.Job, error)
	// CreateIndex(ctx context.Context, index string) error
	// PushDocuments(ctx context.Context, index string, documents []m.Job) error
}
