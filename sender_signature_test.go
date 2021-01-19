package postmark

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCRUDSenderSignatureOK(t *testing.T) {
	if os.Getenv("INTEGRATION") != "true" {
		t.Skip()
	}

	accountToken := os.Getenv("PM_ACCOUNT_TOKEN")
	require.NotEmpty(t, accountToken, "PM_ACCOUNT_TOKEN not set for integration test")

	c := NewClient(accountToken, "")

	//timestamp := time.Now().Unix()

	// Create

	// Read
	readSig, err := c.GetSenderSignature("1540299")
	require.NoError(t, err, "failed to read signature")
	require.NotEmpty(t, readSig, "got empty signature on read")

	// Update

	// Read

	// Delete
}
