package caching

import (
	"context"
)

type OpenSearchCache[T any] interface {
	Insert(ctx context.Context, t T) error
	Update(ctx context.Context, t T) error
	Delete(ctx context.Context, id string) error
	FindOne(ctx context.Context, id string) (T, error)
	FindSimilar(ctx context.Context, index, term string, pageIndex, pageSize int) ([]T, error)
	CreateIndex(ctx context.Context, index string) error
	DeleteAllDocuments(ctx context.Context, index string) error
	PushDocuments(ctx context.Context, index string, documents []T) (success, fail int, err error)
}
