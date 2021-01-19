package postmark

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCRUDServerIntegrationOK(t *testing.T) {
	if os.Getenv("INTEGRATION") != "true" {
		t.Skip()
	}

	accountToken := os.Getenv("PM_ACCOUNT_TOKEN")
	require.NotEmpty(t, accountToken, "PM_ACCOUNT_TOKEN not set for integration test")

	c := NewClient(accountToken, "")

	timestamp := time.Now().Unix()

	// Create
	newServer := Server{
		Name:                       fmt.Sprintf("lib-test-name-%d", timestamp),
		Color:                      "Purple",
		SmtpApiActivated:           true,
		RawEmailEnabled:            true,
		InboundHookUrl:             fmt.Sprintf("https://lib-test-domain-%d.com/inboundhook", timestamp),
		IncludeBounceContentInHook: true,
		PostFirstOpenOnly:          true,
		TrackOpens:                 true,
		TrackLinks:                 "HtmlAndText",
		InboundSpamThreshold:       5,
		EnableSMTPAPIErrorHooks:    true,
	}

	retSrv, err := c.CreateServer(newServer)
	require.NoError(t, err, "errored creating server")
	require.NotEmpty(t, retSrv, "got empty server returned")
	require.NotEmpty(t, retSrv.ID, "missing server ID")

	// Read
	readSrv, err := c.GetServer(strconv.Itoa(retSrv.ID))
	require.NoError(t, err, "error reading server")
	require.NotEmpty(t, readSrv, "got empty server response on read")

	// Update

	// Read

	// Delete
}
