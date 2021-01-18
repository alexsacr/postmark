package postmark

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetServerIntegrationOK(t *testing.T) {
	if os.Getenv("INTEGRATION") != "true" {
		t.Skip()
	}

	accountToken := os.Getenv("PM_ACCOUNT_TOKEN")
	require.NotEmpty(t, accountToken, "PM_ACCOUNT_TOKEN not set for integration test")

	serverID := os.Getenv("PM_TEST_SERVER_ID")
	require.NotEmpty(t, serverID, "PM_TEST_SERVER_ID not set for integration test")

	c := NewClient(accountToken, "")

	s, err := c.GetServer(serverID)
	require.NoError(t, err, "errored getting server by ID")
	require.NotEmpty(t, s, "empty server response")

	assert.NotEmpty(t, s.Name, "empty server name")
}
