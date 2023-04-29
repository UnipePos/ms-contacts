package contact_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/UnipePos/ms-contacts/internal/app/contact"
	"github.com/stretchr/testify/require"
)

var contactJson = []byte(`{
	"id": "1",
	"name": "Thomas",
	"email": "thomas.matheus60@gmail.com"
}`)

var contactEmptyJson = []byte(`{}`)

func TestContact(t *testing.T) {
	var ctc contact.Contact
	body := bytes.NewReader(contactJson)

	err := json.NewDecoder(body).Decode(&ctc)

	require.NotEmpty(t, ctc.ID)
	require.NotEmpty(t, ctc.Email)
	require.NotEmpty(t, ctc.Name)

	require.NotNil(t, ctc)
	require.Nil(t, err)
}

func TestConcactEmptyFields(t *testing.T) {
	var ctc contact.Contact
	body := bytes.NewReader(contactEmptyJson)

	err := json.NewDecoder(body).Decode(&ctc)

	require.Empty(t, ctc.ID)
	require.Empty(t, ctc.Email)
	require.Empty(t, ctc.Name)

	require.NotNil(t, ctc)
	require.Nil(t, err)
}
