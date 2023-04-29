package driver

import (
	"time"

	config "github.com/UnipePos/ms-contacts/config"
	"github.com/UnipePos/ms-contacts/internal/app/contact"

	"github.com/boltdb/bolt"
)

type DB struct {
	SQL *bolt.DB
}

func BoltConnect(cfg *config.Config) (*DB, error) {
	bConn, err := bolt.Open(cfg.DataBasePath, cfg.DataBaseMode, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, err
	}

	err = bConn.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(contact.DbName))

		return err
	})

	if err != nil {
		return nil, err
	}

	return &DB{SQL: bConn}, err
}
