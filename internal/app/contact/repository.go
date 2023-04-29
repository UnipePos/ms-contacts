package contact

import (
	"context"
)

type ContactRepository interface {
	FindAll(ctx context.Context) ([]*Contact, error)
	FindByID(ctx context.Context, uuid string) (*Contact, error)
	Create(ctx context.Context, c *Contact) (string, error)
	Delete(ctx context.Context, uuid string) (bool, error)
}
