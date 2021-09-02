package repository

import "cloud.google.com/go/firestore"

type JobRepository interface {
	Get(id string) (Job, error)
	Create(job Job) (Job, error)
	Update(job Job) (Job, error)
}

type jobRepository struct {
	db *firestore.Client
}

func NewJobRepository(db *firestore.Client) JobRepository {
	return &jobRepository{db: db}
}

func (ref *jobRepository)Get(id string) (Job, error) {
	return Job{}, nil
}

func (ref *jobRepository)Create(job Job) (Job, error) {
	return Job{}, nil
}

func (ref *jobRepository)Update(job Job) (Job, error) {
	return Job{}, nil
}
