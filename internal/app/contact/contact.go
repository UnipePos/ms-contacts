package contact

import "time"

const DbName string = "contacts"

type Contact struct {
	ID           string    `json:"id,omitempty"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	PhoneNumber  string    `json:"phone_number"`
	CreationDate time.Time `json:"creation_date"`
}
