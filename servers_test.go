package postmark

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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
		Name:                    fmt.Sprintf("lib-test-name-%d", timestamp),
		Color:                   "purple",
		SmtpApiActivated:        true,
		RawEmailEnabled:         true,
		InboundHookUrl:          fmt.Sprintf("https://lib-test-domain-%d.com/inboundhook", timestamp),
		TrackOpens:              true,
		TrackLinks:              "HtmlAndText",
		InboundSpamThreshold:    5,
		EnableSMTPAPIErrorHooks: true,
	}

	retSrv, err := c.CreateServer(newServer)
	require.NoError(t, err, "errored creating server")
	require.NotEmpty(t, retSrv, "got empty server returned")
	require.NotEmpty(t, retSrv.ID, "missing server ID")

	// Read
	readSrv, err := c.GetServer(strconv.Itoa(retSrv.ID))
	require.NoError(t, err, "error reading server")
	require.NotEmpty(t, readSrv, "got empty server response on read")

	assert.Equal(t, newServer.Name, readSrv.Name, "name wrong")
	assert.Equal(t, newServer.Color, readSrv.Color, "color wrong")
	assert.Equal(t, newServer.SmtpApiActivated, readSrv.SmtpApiActivated, "smtpapi wrong")
	assert.Equal(t, newServer.RawEmailEnabled, readSrv.RawEmailEnabled, "rawemail wrong")
	assert.Equal(t, newServer.InboundHookUrl, readSrv.InboundHookUrl, "hook url wrong")
	assert.Equal(t, newServer.PostFirstOpenOnly, readSrv.PostFirstOpenOnly, "firstopen wrong")
	assert.Equal(t, newServer.TrackOpens, readSrv.TrackOpens, "trackopens wrong")
	assert.Equal(t, newServer.TrackLinks, readSrv.TrackLinks, "tracklinks wrong")
	assert.Equal(t, newServer.InboundSpamThreshold, readSrv.InboundSpamThreshold, "spam threshold wrong")
	assert.Equal(t, newServer.EnableSMTPAPIErrorHooks, readSrv.EnableSMTPAPIErrorHooks, "errorhooks wrong")

	// Update
	updatedServer := Server{
		Name:                    newServer.Name + "-updated",
		Color:                   "blue",
		SmtpApiActivated:        false,
		RawEmailEnabled:         false,
		InboundHookUrl:          newServer.InboundHookUrl + "-updated",
		TrackOpens:              false,
		TrackLinks:              "TextOnly",
		InboundSpamThreshold:    4,
		EnableSMTPAPIErrorHooks: false,
	}

	err = c.UpdateServer(strconv.Itoa(readSrv.ID), updatedServer)
	require.NoError(t, err, "error updating server")

	// Read
	readSrv, err = c.GetServer(strconv.Itoa(retSrv.ID))
	require.NoError(t, err, "error reading server")
	require.NotEmpty(t, readSrv, "got empty server response on read")

	assert.Equal(t, updatedServer.Name, readSrv.Name, "name wrong")
	assert.Equal(t, updatedServer.Color, readSrv.Color, "color wrong")
	assert.Equal(t, updatedServer.SmtpApiActivated, readSrv.SmtpApiActivated, "smtpapi wrong")
	assert.Equal(t, updatedServer.RawEmailEnabled, readSrv.RawEmailEnabled, "rawemail wrong")
	assert.Equal(t, updatedServer.InboundHookUrl, readSrv.InboundHookUrl, "hook url wrong")
	assert.Equal(t, updatedServer.TrackOpens, readSrv.TrackOpens, "trackopens wrong")
	assert.Equal(t, updatedServer.TrackLinks, readSrv.TrackLinks, "tracklinks wrong")
	assert.Equal(t, updatedServer.InboundSpamThreshold, readSrv.InboundSpamThreshold, "spam threshold wrong")
	assert.Equal(t, updatedServer.EnableSMTPAPIErrorHooks, readSrv.EnableSMTPAPIErrorHooks, "errorhooks wrong")

	// Delete

}
