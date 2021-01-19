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

func TestCRUDSenderSignatureOK(t *testing.T) {
	if os.Getenv("INTEGRATION") != "true" {
		t.Skip()
	}

	accountToken := os.Getenv("PM_ACCOUNT_TOKEN")
	require.NotEmpty(t, accountToken, "PM_ACCOUNT_TOKEN not set for integration test")

	c := NewClient(accountToken, "")

	timestamp := time.Now().Unix()

	// Create
	sigOpts := SenderSignatureOpts{
		FromEmail:        fmt.Sprintf("test+%d@example123.com", timestamp),
		Name:             "foobarbaz",
		ReplyToEmail:     "test-reply@example123.com",
		ReturnPathDomain: "pm.example123.com",
	}

	retSig, err := c.CreateSenderSignature(sigOpts)
	require.NoError(t, err, "error creating signature")
	require.NotEmpty(t, retSig, "got empty signature object returned on create")

	// Read
	readSig, err := c.GetSenderSignature(strconv.Itoa(retSig.ID))
	require.NoError(t, err, "failed to read signature")
	require.NotEmpty(t, readSig, "got empty signature on read")

	assert.Equal(t, sigOpts.FromEmail, readSig.EmailAddress, "wrong email on read")
	assert.Equal(t, sigOpts.Name, readSig.Name, "wrong name on read")
	assert.Equal(t, sigOpts.ReplyToEmail, readSig.ReplyToEmailAddress, "wrong replyto on read")
	assert.Equal(t, sigOpts.ReturnPathDomain, readSig.ReturnPathDomain, "wrong return path domain on read")

	// Update
	upOpts := SenderSignatureOpts{
		Name:             "foobarbaz-updated",
		ReplyToEmail:     "test-reply-updated@example123.com",
		ReturnPathDomain: "pm-updated.example123.com",
	}

	err = c.UpdateSenderSignature(strconv.Itoa(retSig.ID), upOpts)
	require.NoError(t, err, "error on update")

	// Read
	upSig, err := c.GetSenderSignature(strconv.Itoa(retSig.ID))
	require.NoError(t, err, "failed to read signature")
	require.NotEmpty(t, upSig, "got empty signature on read")

	assert.Equal(t, upOpts.Name, upSig.Name, "wrong name on read")
	assert.Equal(t, upOpts.ReplyToEmail, upSig.ReplyToEmailAddress, "wrong replyto on read")
	assert.Equal(t, upOpts.ReturnPathDomain, upSig.ReturnPathDomain, "wrong return path domain on read")

	// Delete
	err = c.DeleteSenderSignature(strconv.Itoa(retSig.ID))
	require.NoError(t, err, "error on deletion")
}
