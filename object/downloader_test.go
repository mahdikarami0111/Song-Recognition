package object

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDownloadObject(t *testing.T) {
	bucket := "cc1-hw1-mk"
	filename := "Emotionally Scarred.mp3"
	err := DownloadObject(bucket, filename)
	require.NoError(t, err)
}
