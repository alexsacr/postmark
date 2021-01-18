package postmark

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetTemplateIntegrationOK(t *testing.T) {
	if os.Getenv("INTEGRATION") != "true" {
		t.Skip()
	}

	serverToken := os.Getenv("PM_SERVER_TOKEN")
	require.NotEmpty(t, serverToken, "PM_SERVER_TOKEN not set for integration test")

	templateID := os.Getenv("PM_TEST_TOKEN_ID")
	require.NotEmpty(t, templateID, "PM_TEST_TOKEN_ID not set for integration test")

	c := NewClient("", serverToken)

	tmpl, err := c.GetTemplate(templateID)
	require.NoError(t, err, "errored getting template by id")
	require.NotEmpty(t, tmpl, "empty template response")

	assert.NotEmpty(t, tmpl.TemplateID, "template ID empty")
}
