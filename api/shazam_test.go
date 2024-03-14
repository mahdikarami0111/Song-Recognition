package api

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShazamSearch(t *testing.T) {
	res := shazamApi("Emotionally Scarred.mp3")
	require.NotEmpty(t, res)
}
