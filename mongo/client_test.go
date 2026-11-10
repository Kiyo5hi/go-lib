package mongo

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewClientWithoutOptionsRejectsInvalidURIWithoutPanic(t *testing.T) {
	// Invalid URI is rejected locally, without a database or DNS dependency.
	client, err := NewClient(context.Background(), &Config{Host: "invalid host"})
	require.Error(t, err)
	require.Nil(t, client)
}
