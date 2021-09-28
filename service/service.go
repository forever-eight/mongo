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

func (s *Service) Change(input *ds.ProjectString) error {
	id, err := primitive.ObjectIDFromHex(input.ID)
	if err != nil {
		return err
	}

	exist, err := s.r.FindProjectByID(s.ctx, id)
	if err != nil {
		return err
	}
	if input.Title != "" {
		exist.Title = input.Title
	}
	// todo исправить живо, потому что не хочет webhook и apiurl вставляться
	if input.Channels.Jivo != nil {
		for i := 0; i < len(input.Channels.Jivo); i++ {
			exist.Channels.Jivo = append(exist.Channels.Jivo, input.Channels.Jivo[i])
			exist.Channels.Jivo[len(exist.Channels.Jivo)-1].WebhookPath = input.Channels.Jivo[i].WebhookPath
			exist.Channels.Jivo[len(exist.Channels.Jivo)-1].ApiUrl = input.Channels.Jivo[i].ApiUrl
		}
	}
	if input.Channels.Tg != nil {
		for i := 0; i < len(input.Channels.Tg); i++ {
			exist.Channels.Tg = append(exist.Channels.Tg, input.Channels.Tg[i])
		}
	}
	if input.Channels.Vk != nil {
		for i := 0; i < len(input.Channels.Vk); i++ {
			exist.Channels.Vk = append(exist.Channels.Vk, input.Channels.Vk[i])
		}
	}
	return s.r.ChangeProject(s.ctx, id, exist)
}
