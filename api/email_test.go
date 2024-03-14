package api

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEmail(t *testing.T) {
	apiKey := "45f3ba214ce75a269d6c9ad10be25618-b02bcf9f-781f5621"
	domain := "sandbox7de32744eb5e46c9921f51799d0aca30.mailgun.org"
	resp, err := SendEmail(domain, apiKey, "mahdikarami0115@gmail.com", "this is a teast message nigger")
	require.NoError(t, err)
	log.Println(resp)
}
