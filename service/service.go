package service

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/forever-eight/mongo.git/ds"
	"github.com/forever-eight/mongo.git/repository"
)

type Service struct {
	r   *repository.Repository
	ctx context.Context
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		r: repos}

}

func (s *Service) Delete(ID primitive.ObjectID) error {
	return s.r.DeleteProject(s.ctx, ID)
}

func (s *Service) Add(project *ds.Project) error {
	return s.r.AddProject(s.ctx, project)
}

func (s *Service) Find(ID primitive.ObjectID) (*ds.Project, error) {
	return s.r.FindProjectByID(s.ctx, ID)
}
