package postmark

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClientOK(t *testing.T) {
	c := NewClient("foo", "bar")
	assert.Equal(t, "foo", c.AccountToken, "account token set wrong")
	assert.Equal(t, "bar", c.ServerToken, "server token set wrong")
	assert.Equal(t, PostmarkAPIAddr, c.APIAddr, "api addr set wrong")
}

func TestAPIErrorIntegration(t *testing.T) {
	// Make sure api errors are handled as expected.
	if os.Getenv("INTEGRATION") != "true" {
		t.Skip()
	}

	c := NewClient("", "")

	res, err := c.GetTemplate("bar")
	assert.Empty(t, res, "response object was not empty on error")
	require.NotNil(t, err, "error was not returned on intentional api failure")

	expError := "api error code 10: No Account or Server API tokens were supplied in the HTTP headers. Please add a header for either X-Postmark-Server-Token or X-Postmark-Account-Token."
	assert.Equal(t, expError, err.Error(), "incorrect error returned")
}
