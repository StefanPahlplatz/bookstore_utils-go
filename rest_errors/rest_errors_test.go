package rest_errors

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("message", errors.New("database error"))
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, err.Status)
	assert.Equal(t, "message", err.Message)
	assert.Equal(t, "internal_server_error", err.Error)

	assert.NotNil(t, err.Causes)
	assert.Equal(t, 1, len(err.Causes))
	assert.Equal(t, "database error", err.Causes[0])
}

func TestNewBadRequestError(t *testing.T) {

}

func TestNewNotFoundError(t *testing.T) {

}
