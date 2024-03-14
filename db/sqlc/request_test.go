package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateRequest(t *testing.T) {
	arg := CreateRequestParams{
		Email:  "email2@gmail.com",
		Status: "ready",
	}
	req, err := testQueries.CreateRequest(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, req)

}
