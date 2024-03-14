package api

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSpotifySearch(t *testing.T) {
	res, err := spotifySearch("Emotionally Scarred")
	require.NoError(t, err)
	require.NotEmpty(t, res)
	log.Println(res)
}

func TestSpotifyRecommend(t *testing.T) {
	res, err := spotifyRecommend("7ge7BHazYbVKi8qyZUX1Bm")
	require.NoError(t, err)
	require.NotEmpty(t, res)
	log.Println(res)
}
