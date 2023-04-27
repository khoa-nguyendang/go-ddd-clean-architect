package opensearch

import (
	ex "app/core/exceptions"
	c "app/core/interfaces/caching"
	m "app/core/models"
	"app/pkg/logger"
	ops "app/pkg/opensearch"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/jinzhu/copier"
	osapi "github.com/opensearch-project/opensearch-go/v2/opensearchapi"
	"github.com/opensearch-project/opensearch-go/v2/opensearchutil"
)

var (
	jobSearchFields []string = []string{
		"pk",
		"status",
		"content",
		"title",
		"short_description",
		"full_description",
		"tags",
		"company_name",
		"company_name",
	}
)

type JobStorage struct {
	logger     logger.Logger
	opensearch ops.OpenSearch
	timeout    time.Duration
}

func New(opensearch ops.OpenSearch, logger *logger.ApiLogger) (c.JobCache, error) {
	return JobStorage{
		opensearch: opensearch,
		timeout:    time.Second * 300,
		logger:     logger,
	}, nil
}

// CreateIndex implements caching.JobCache
func (js JobStorage) CreateIndex(ctx context.Context, index string) error {
	return js.opensearch.CreateIndex(index)
}

func (js JobStorage) Insert(ctx context.Context, job m.Job) error {
	bdy, err := json.Marshal(job)
	if err != nil {
		return err
	}

	// res, err := js.ops.Client.Create()
	req := osapi.CreateRequest{
		Index:      js.opensearch.Alias,
		DocumentID: job.PK,
		Body:       bytes.NewReader(bdy),
	}

	ctx, cancel := context.WithTimeout(ctx, js.timeout)
	defer cancel()

	res, err := req.Do(ctx, js.opensearch.Client)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode == 409 {
		return ex.ErrConflict
	}

	if res.IsError() {
		return err
	}

	return nil
}

func (js JobStorage) Update(ctx context.Context, job m.Job) error {
	bdy, err := json.Marshal(job)
	if err != nil {
		return err
	}

	// res, err := js.ops.Client.Update()
	req := osapi.UpdateRequest{
		Index:      js.opensearch.Alias,
		DocumentID: job.PK,
		Body:       bytes.NewReader([]byte(fmt.Sprintf(`{"doc":%s}`, bdy))),
	}

	ctx, cancel := context.WithTimeout(ctx, js.timeout)
	defer cancel()

	res, err := req.Do(ctx, js.opensearch.Client)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		return ex.ErrNotFound
	}

	if res.IsError() {
		return err
	}

	return nil
}

func (js JobStorage) Delete(ctx context.Context, id string) error {
	// res, err := js.ops.Client.Delete()
	req := osapi.DeleteRequest{
		Index:      js.opensearch.Alias,
		DocumentID: id,
	}

	ctx, cancel := context.WithTimeout(ctx, js.timeout)
	defer cancel()

	res, err := req.Do(ctx, js.opensearch.Client)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		return ex.ErrNotFound
	}

	if res.IsError() {
		return err
	}

	return nil
}

func (js JobStorage) FindOne(ctx context.Context, id string) (m.Job, error) {
	// res, err := js.ops.Client.Get()
	req := osapi.GetRequest{
		Index:      js.opensearch.Alias,
		DocumentID: id,
	}

	ctx, cancel := context.WithTimeout(ctx, js.timeout)
	defer cancel()

	res, err := req.Do(ctx, js.opensearch.Client)
	if err != nil {
		return m.Job{}, err
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		return m.Job{}, ex.ErrNotFound
	}

	if res.IsError() {
		return m.Job{}, err
	}

	var (
		job  m.Job
		body ops.Document[m.Job]
	)
	body.Source = job

	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		return m.Job{}, err
	}

	return job, nil
}

func (js JobStorage) FindSimilar(ctx context.Context, index, term string, pageIndex, pageSize int) ([]m.Job, error) {
	searchQueryObj := ops.SearchQuery{
		From: pageIndex,
		Size: pageSize,
		Query: ops.Query{
			MultiMatch: ops.MultiMatch{
				Query:  term,
				Fields: jobSearchFields,
			},
		},
	}
	searchQuery, err := json.Marshal(searchQueryObj)
	if err != nil {
		return make([]m.Job, 0), err
	}
	content := strings.NewReader(string(searchQuery))

	search := osapi.SearchRequest{
		Index: []string{index},
		Body:  content,
	}

	searchResponse, err := search.Do(ctx, js.opensearch.Client)
	if err != nil {
		js.logger.Errorf("FindSimilar.searchResponse.err %v \n", err)
		return make([]m.Job, 0), err
	}

	r := ops.SearchResponse[m.Job]{}
	if err = json.NewDecoder(searchResponse.Body).Decode(&r); err != nil {
		js.logger.Errorf("FindSimilar.Decode.err %v \n", err)
		return make([]m.Job, 0), err
	}

	jobs := make([]m.Job, 0)
	documents := make([]ops.Document[m.Job], 0)
	if err = copier.Copy(&documents, &r.Hits.Hits); err != nil {
		return make([]m.Job, 0), err
	}
	for _, v := range documents {
		jobs = append(jobs, v.Source)
	}
	js.logger.Infof("FindSimilar.Decode.sucess %v \n", r)
	return jobs, nil
}

func (js JobStorage) AppendItemToBulkIndex(indexer opensearchutil.BulkIndexer, job m.Job) error {
	jsonData, marshalErr := json.Marshal(job)
	if marshalErr != nil {
		return marshalErr
	}
	err := indexer.Add(
		context.Background(),
		opensearchutil.BulkIndexerItem{
			Action:     "index",
			DocumentID: job.PK,
			Body:       strings.NewReader(string(jsonData)),
			OnSuccess: func(ctx context.Context, item opensearchutil.BulkIndexerItem, res opensearchutil.BulkIndexerResponseItem) {
				//Perhaps write logs about success here?
			},
			OnFailure: func(ctx context.Context, item opensearchutil.BulkIndexerItem, res opensearchutil.BulkIndexerResponseItem, err error) {
				if err != nil {
					js.logger.Errorf("ERROR: %s", err)
					return
				}
				js.logger.Infof("ERROR: %s: %s", res.Error.Type, res.Error.Reason)
			},
		},
	)
	if err != nil {
		js.logger.Errorf("Unexpected error: %s", err)
	}
	return err
}

func (js JobStorage) PushDocuments(ctx context.Context, index string, documents []m.Job) (success, fail int, err error) {
	// Create the indexer
	indexer, err := opensearchutil.NewBulkIndexer(opensearchutil.BulkIndexerConfig{
		Client:        js.opensearch.Client, // The OpenSearch client
		Index:         index,                // The default index name
		NumWorkers:    8,                    // The number of worker goroutines (default: number of CPUs)
		FlushBytes:    5e+8,                 // The flush threshold in bytes (default: 500M),
		FlushInterval: 300 * time.Second,    // The periodic flush interval
	})
	if err != nil {
		return 0, 0, err
	}
	js.logger.Infof("PushDocuments: index: %v, length: %v \n", index, len(documents))
	// Add an item to the indexer
	for _, job := range documents {
		err = js.AppendItemToBulkIndex(indexer, job)
		if err != nil {
			js.logger.Errorf("Unexpected error when create index for document id:%s - %v", job.PK, err)
		}
	}
	// Close the indexer channel and flush remaining items
	if err := indexer.Close(context.Background()); err != nil {
		js.logger.Errorf("Unexpected error: %s", err)
	}

	// Report the indexer statistics
	stats := indexer.Stats()
	if stats.NumFailed > 0 {
		js.logger.Errorf("Indexed [%d] documents with [%d] errors", stats.NumFlushed, stats.NumFailed)
	} else {
		js.logger.Infof("Successfully indexed [%d] documents", stats.NumFlushed)
	}
	return int(stats.NumFlushed), int(stats.NumFailed), nil
}

func (js JobStorage) DeleteAllDocuments(ctx context.Context, index string) error {
	var err error
	// js.opensearch.Client.DeleteByQuery()
	return err
}
