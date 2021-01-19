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

func TestCRUDDomain(t *testing.T) {
	if os.Getenv("INTEGRATION") != "true" {
		t.Skip()
	}

	accountToken := os.Getenv("PM_ACCOUNT_TOKEN")
	require.NotEmpty(t, accountToken, "PM_ACCOUNT_TOKEN not set for integration test")

	c := NewClient(accountToken, "")

	timestamp := time.Now().Unix()

	testDomain := fmt.Sprintf("lib-test-name-%d.com", timestamp)

	// Create
	newDomain := Domain{
		Name:             testDomain,
		ReturnPathDomain: fmt.Sprintf("bounces.%s", testDomain),
	}

	retDomain, err := c.CreateDomain(newDomain)
	require.NoError(t, err, "errored creating domain")
	require.NotEmpty(t, retDomain, "got empty domain returned")
	require.NotEmpty(t, retDomain.ID, "missing domain ID")

	// Read
	readDomain, err := c.GetDomain(strconv.Itoa(retDomain.ID))
	require.NoError(t, err, "error reading domain")
	require.NotEmpty(t, readDomain, "got empty domain response on read")

	assert.Equal(t, newDomain.Name, readDomain.Name, "name wrong")
	assert.Equal(t, newDomain.ReturnPathDomain, readDomain.ReturnPathDomain, "return path wrong")

	// Update
	updatedDomain := Domain{
		ReturnPathDomain: fmt.Sprintf("bounces-2.%s", testDomain),
	}

	err = c.UpdateDomain(strconv.Itoa(readDomain.ID), updatedDomain)
	require.NoError(t, err, "error updating domain")

	// Read
	readDomain, err = c.GetDomain(strconv.Itoa(readDomain.ID))

	assert.Equal(t, updatedDomain.ReturnPathDomain, readDomain.ReturnPathDomain, "return path not updated")

	// Delete
	err = c.DeleteDomain(strconv.Itoa(retDomain.ID))
	require.NoError(t, err, "delete failed")
}
