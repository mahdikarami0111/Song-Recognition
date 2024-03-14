package object

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUploadObject(t *testing.T) {
	bucket := "cc1-hw1-mk"
	filename := "Emotionally Scarred.mp3"
	err := UploadObject(bucket, filename)
	require.NoError(t, err)
}
