package mysql

import (
	"app/core/entities"
	ir "app/core/interfaces/repositories"
	mds "app/core/models"
	"app/pkg/logger"
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Repo struct {
	db     *sqlx.DB
	logger logger.Logger
}

// AddTestJob implements repositories.JobRepo.
func (r *Repo) AddTestJob(ctx context.Context) error {
	result, err := r.db.ExecContext(
		ctx,
		JOBS_ADD_RAW,
	)
	r.logger.Infof("AddTestJob execute database: %v", result)
	if err != nil {
		r.logger.Errorf("AddTestJob execute database.err: %v", err)
		return err
	}

	return err
}

func New(db *sqlx.DB, logger *logger.ApiLogger) ir.JobRepo {
	return &Repo{
		db:     db,
		logger: logger,
	}
}

// AddJob implements repositories.JobRepo
func (r *Repo) AddJob(ctx context.Context, job mds.Job) (entities.Job, error) {
	newId := uuid.New().String()
	result, err := r.db.ExecContext(
		ctx,
		JOBS_ADD_RAW,
		newId,
		job.Status,
		job.CreatedDate,
		job.ActivatedDate,
		job.OrderId,
		job.Comments,
		job.AuthorId,
		job.Title,
		job.Tags,
		job.CompanyId,
		job.Rating,
		job.ShortDescription,
		job.FullDescription,
		job.TopicID,
		job.Code,
		job.ImageThumbnailPath,
		job.VideoThumbnailPath,
		job.IsApproved,
	)
	r.logger.Infof("AddJob execute database: %v", result)
	if err != nil {
		r.logger.Errorf("AddJob execute database.err: %v", err)
		return entities.Job{}, err
	}

	jobEntity, err := r.GetJob(ctx, newId)
	return jobEntity, err
}

// DeleteJob implements repositories.JobRepo
func (r *Repo) DeleteJob(ctx context.Context, jobId string) (int, error) {
	panic("unimplemented")
}

// GetJob implements repositories.JobRepo
func (r *Repo) GetJob(ctx context.Context, jobId string) (entities.Job, error) {
	job := entities.Job{}
	if err := r.db.QueryRowContext(ctx, JOB_GET_RAW, jobId).Scan(&job); err != nil {
		r.logger.Infof("GetJob error for id: %v, err: %v", jobId, err)
		if errors.Cause(err) == sql.ErrNoRows {
			return entities.Job{}, nil
		}
		return entities.Job{}, err
	}

	return job, nil
}

// PatchJob implements repositories.JobRepo
func (r *Repo) PatchJob(ctx context.Context, job mds.Job) (entities.Job, error) {
	panic("unimplemented")
}

// SearchJob implements repositories.JobRepo
func (r *Repo) SearchJob(ctx context.Context, term string, pageIndex int, pageAmount int) ([]entities.Job, error) {
	var jobs []entities.Job = make([]entities.Job, 0)
	skips := (pageIndex - 1) * pageAmount
	if skips < 0 {
		skips = 0
	}
	r.logger.Infof("SearchJob: %v - %v - %v ", term, skips, pageAmount)
	rows, err := r.db.Query(JOBS_GET_RAW, term, pageAmount, skips)
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
func (r *Repo) UpdateJob(ctx context.Context, job mds.Job) (entities.Job, error) {
	result, err := r.db.ExecContext(
		ctx,
		JOBS_UPDATE_RAW,
		job.Status,
		job.CreatedDate,
		job.ActivatedDate,
		job.OrderId,
		job.Comments,
		job.AuthorId,
		job.Title,
		job.Tags,
		job.CompanyId,
		job.Rating,
		job.ShortDescription,
		job.FullDescription,
		job.TopicID,
		job.Code,
		job.ImageThumbnailPath,
		job.VideoThumbnailPath,
		job.IsApproved,
		job.PK,
	)
	r.logger.Infof("UpdateJob execute database: %v", result)
	if err != nil {
		r.logger.Errorf("UpdateJob execute database.err: %v", err)
		return entities.Job{}, err
	}

	jobEntity, err := r.GetJob(ctx, job.PK)
	return jobEntity, err
}
