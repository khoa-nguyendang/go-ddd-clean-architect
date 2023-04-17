package mysql

import (
	"app/core/entities"
	e "app/core/entities"
	ir "app/core/interfaces/repositories"
	"app/pkg/logger"
	"context"

	"github.com/jmoiron/sqlx"
)

type Repo struct {
	db     *sqlx.DB
	logger logger.Logger
}

func New(db *sqlx.DB, logger *logger.ApiLogger) ir.JobRepo {
	return &Repo{
		db:     db,
		logger: logger,
	}
}

// AddJob implements repositories.JobRepo
func (r *Repo) AddJob(ctx context.Context, job interface{}) (e.Job, error) {
	panic("unimplemented")
}

// DeleteJob implements repositories.JobRepo
func (r *Repo) DeleteJob(ctx context.Context, jobId string) (int, error) {
	panic("unimplemented")
}

// GetJob implements repositories.JobRepo
func (r *Repo) GetJob(ctx context.Context, jobId string) (e.Job, error) {
	panic("unimplemented")
}

// PatchJob implements repositories.JobRepo
func (r *Repo) PatchJob(ctx context.Context, job interface{}) (e.Job, error) {
	panic("unimplemented")
}

// SearchJob implements repositories.JobRepo
func (r *Repo) SearchJob(ctx context.Context, term string, pageIndex int, pageAmount int) ([]e.Job, error) {
	var jobs []entities.Job = make([]e.Job, 0)
	skips := (pageIndex - 1) * pageAmount
	if skips < 0 {
		skips = 0
	}
	r.logger.Infof("SearchJob: %v - %v - %v ", term, skips, pageAmount)
	rows, err := r.db.Query(JOBS_GET, term, pageAmount, skips)
	if err != nil {
		r.logger.Infof("r.db.Query err: %v", err)
		return jobs, err
	}
	defer rows.Close()
	for rows.Next() {
		j := entities.Job{}
		err = rows.Scan(&j.PK,
			&j.Status,
			&j.CreatedDate,
			&j.ActivatedDate,
			&j.OrderId,
			&j.Content,
			&j.AuthorId,
			&j.Title,
			&j.Tags,
			&j.CompanyId,
			&j.Rating,
			&j.ShortDescription,
			&j.FullDescription,
			&j.TopicID,
			&j.Code,
			&j.ImageThumbnailPath,
			&j.VideoThumbnailPath,
			&j.IsApproved,
			&j.Comments,
			&j.Author,
			&j.Company,
		)
		if err != nil {
			r.logger.Error(err)
			continue
		}
		jobs = append(jobs, j)
	}
	return jobs, nil
}

// UpdateJob implements repositories.JobRepo
func (r *Repo) UpdateJob(ctx context.Context, job interface{}) (e.Job, error) {
	panic("unimplemented")
}
