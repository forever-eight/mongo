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

func (s *Service) Delete(ID string) error {
	id, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	return s.r.DeleteProject(s.ctx, id)
}

func (s *Service) Add(project *ds.Project) error {
	id := primitive.NewObjectID()
	project.ID = id
	return s.r.AddProject(s.ctx, project)
}

func (s *Service) Find(ID string) (*ds.Project, error) {
	id, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, err
	}
	return s.r.FindProjectByID(s.ctx, id)
}
