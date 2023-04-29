package boltdb

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/UnipePos/ms-contacts/internal/app/contact"

	"github.com/boltdb/bolt"
)

type contactRepository struct {
	conn *bolt.DB
}

func NewContactRepository(db *bolt.DB) contact.ContactRepository {
	return &contactRepository{
		conn: db,
	}
}

func (cr *contactRepository) FindAll(ctx context.Context) ([]*contact.Contact, error) {
	cts := make([]*contact.Contact, 0)
	err := cr.conn.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(contact.DbName))
		return b.ForEach(func(k, v []byte) error {
			var temp *contact.Contact
			err := json.Unmarshal(v, &temp)

			if err != nil {
				return err
			}

			cts = append(cts, temp)

			return nil
		})
	})

	return cts, err
}

func (cr *contactRepository) FindByID(ctx context.Context, uuid string) (*contact.Contact, error) {
	var ct *contact.Contact
	err := cr.conn.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(contact.DbName))
		ctc := b.Get([]byte(uuid))

		err := json.Unmarshal(ctc, &ct)
		if err != nil {
			return err
		}

		return nil
	})

	if ct == nil {
		return nil, errors.New("contact not found.")
	}

	return ct, err
}

func (cr *contactRepository) Create(ctx context.Context, c *contact.Contact) (string, error) {
	err := cr.conn.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(contact.DbName))

		contact, err := json.Marshal(c)
		if err != nil {
			return err
		}

		return b.Put([]byte(c.ID), contact)
	})

	return c.ID, err
}

func (cr *contactRepository) Delete(ctx context.Context, uuid string) (bool, error) {
	err := cr.conn.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(contact.DbName))

		if cont := bucket.Get([]byte(uuid)); cont == nil {
			return errors.New("contact not found.")
		}

		return bucket.Delete([]byte(uuid))
	})

	if err != nil {
		return false, err
	}

	return true, nil
}
