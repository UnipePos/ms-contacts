package contact

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type ContactService interface {
	FindAll(ctx context.Context) ([]*Contact, error)
	FindByID(ctx context.Context, uuid string) (*Contact, error)
	Create(ctx context.Context, c *Contact) (string, error)
	Delete(ctx context.Context, uuid string) (bool, error)
}

type contactService struct {
	repo ContactRepository
}

func NewContactService(repo ContactRepository) ContactService {
	return &contactService{
		repo: repo,
	}
}

func (cs *contactService) FindAll(ctx context.Context) ([]*Contact, error) {
	return cs.repo.FindAll(ctx)
}

func (cs *contactService) FindByID(ctx context.Context, uuid string) (*Contact, error) {
	return cs.repo.FindByID(ctx, uuid)
}

func (cs *contactService) Create(ctx context.Context, c *Contact) (string, error) {
	c.ID = uuid.New().String()
	c.CreationDate = time.Now()

	return cs.repo.Create(ctx, c)
}

func (cs *contactService) Delete(ctx context.Context, uuid string) (bool, error) {
	return cs.repo.Delete(ctx, uuid)
}
