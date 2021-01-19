package postmark

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCRUDTemplateIntegrationOK(t *testing.T) {
	if os.Getenv("INTEGRATION") != "true" {
		t.Skip()
	}

	// Create
	serverToken := os.Getenv("PM_SERVER_TOKEN")
	require.NotEmpty(t, serverToken, "PM_SERVER_TOKEN not set for integration test")

	c := NewClient("", serverToken)

	timestamp := time.Now().Unix()

	newTmpl := Template{
		Name:         fmt.Sprintf("lib-test-name-%d", timestamp),
		Alias:        fmt.Sprintf("lib-test-alias-%d", timestamp),
		Subject:      fmt.Sprintf("lib-test-subject-%d", timestamp),
		HTMLBody:     fmt.Sprintf("lib-test-htmlbody-%d", timestamp),
		TextBody:     fmt.Sprintf("lib-test-textbody-%d", timestamp),
		Active:       true,
		TemplateType: "Standard",
	}

	retT, err := c.CreateTemplate(newTmpl)
	require.NoError(t, err, "error creating template")
	require.NotEmpty(t, retT, "got empty return template")
	require.NotEmpty(t, retT.TemplateID, "got empty template id")

	// Read
	readTmpl, err := c.GetTemplate(retT.Alias)
	require.NoError(t, err, "error getting template")
	require.NotEmpty(t, readTmpl, "got empty template from read")

	assert.Equal(t, newTmpl.Name, readTmpl.Name, "wrong name on readback")
	assert.Equal(t, newTmpl.Alias, readTmpl.Alias, "wrong alias on readback")
	assert.Equal(t, newTmpl.Subject, readTmpl.Subject, "wrong subject on readback")
	assert.Equal(t, newTmpl.HTMLBody, readTmpl.HTMLBody, "wrong html on readback")
	assert.Equal(t, newTmpl.TextBody, readTmpl.TextBody, "wrong text on readback")
	assert.Equal(t, newTmpl.Active, readTmpl.Active, "wrong active on readback")
	assert.Equal(t, newTmpl.TemplateType, readTmpl.TemplateType, "wrong type on readback")
	assert.Equal(t, newTmpl.LayoutTemplate, readTmpl.LayoutTemplate, "wrong layout on readback")

	// Update
	upTmpl := Template{
		Name:     readTmpl.Name + "-update",
		Alias:    readTmpl.Alias + "-update",
		Subject:  readTmpl.Subject + "-update",
		HTMLBody: readTmpl.HTMLBody + "-update",
		TextBody: readTmpl.TextBody + "-update",
	}
	err = c.UpdateTemplate(retT.Alias, upTmpl)
	require.NoError(t, err, "got error updating template")

	// Read
	readTmpl, err = c.GetTemplate(upTmpl.Alias)
	require.NoError(t, err, "error getting template")
	require.NotEmpty(t, readTmpl, "got empty template from read")

	assert.Equal(t, upTmpl.Name, readTmpl.Name, "wrong name on update readback")
	assert.Equal(t, upTmpl.Alias, readTmpl.Alias, "wrong alias on update readback")
	assert.Equal(t, upTmpl.Subject, readTmpl.Subject, "wrong subject on update readback")
	assert.Equal(t, upTmpl.HTMLBody, readTmpl.HTMLBody, "wrong html on update readback")
	assert.Equal(t, upTmpl.TextBody, readTmpl.TextBody, "wrong text on update readback")

	// Delete
	err = c.DeleteTemplate(upTmpl.Alias)
	require.NoError(t, err, "got error deleting template")
}
