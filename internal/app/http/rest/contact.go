package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/UnipePos/ms-contacts/internal/app/contact"
	"github.com/UnipePos/ms-contacts/internal/pkg/log"

	"github.com/go-chi/chi/v5"
)

type ContactHandler interface {
	GetContacts(rw http.ResponseWriter, r *http.Request)
	GetContact(rw http.ResponseWriter, r *http.Request)
	PostContact(rw http.ResponseWriter, r *http.Request)
	DeleteContact(rw http.ResponseWriter, r *http.Request)
}

type contactHandler struct {
	contactService contact.ContactService
}

func NewContactHandler(cs contact.ContactService) ContactHandler {
	return &contactHandler{
		contactService: cs,
	}
}

func (ch *contactHandler) GetContacts(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)

	lg := log.WithContext(r.Context())

	contacts, err := ch.contactService.FindAll(r.Context())
	if err != nil {
		lg.Error(err)
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}

	if err := json.NewEncoder(rw).Encode(contacts); err != nil {
		lg.Errorf("cannot encode json: %s", err)
		http.Error(rw, "cannot encode json", http.StatusBadRequest)
	}
}

func (ch *contactHandler) GetContact(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)

	lg := log.WithContext(r.Context())

	id := chi.URLParam(r, "id")

	ct, err := ch.contactService.FindByID(r.Context(), id)
	if err != nil {
		lg.Error(err)
		http.Error(rw, err.Error(), http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(rw).Encode(ct); err != nil {
		lg.Errorf("cannot encode json: %s", err)
		http.Error(rw, "cannot encode json", http.StatusBadRequest)
		return
	}

	http.Error(rw, "contact not found", http.StatusNotFound)
}

func (ch *contactHandler) PostContact(rw http.ResponseWriter, r *http.Request) {
	lg := log.WithContext(r.Context())

	var ctc contact.Contact
	if err := json.NewDecoder(r.Body).Decode(&ctc); err != nil {
		lg.Errorf("cannot read json: %s", err)
		http.Error(rw, "cannot read json", http.StatusBadRequest)
		return
	}

	id, err := ch.contactService.Create(r.Context(), &ctc)
	if err != nil {
		lg.Errorf("cannot insert data: %s", err)
		http.Error(rw, "cannot insert data", http.StatusBadRequest)
		return
	}

	rw.Header().Add("Location", fmt.Sprintf("/contacts/%s", id))
	rw.WriteHeader(http.StatusCreated)
}

func (ch *contactHandler) DeleteContact(rw http.ResponseWriter, r *http.Request) {
	lg := log.WithContext(r.Context())
	rw.WriteHeader(http.StatusNoContent)

	id := chi.URLParam(r, "id")

	_, err := ch.contactService.Delete(r.Context(), id)

	if err != nil {
		lg.Error("contact not found. error: %s", id)
		http.Error(rw, err.Error(), http.StatusNotFound)
		return
	}
}
