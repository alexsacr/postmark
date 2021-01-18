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
		HtmlBody:     fmt.Sprintf("lib-test-htmlbody-%d", timestamp),
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
	assert.Equal(t, newTmpl.HtmlBody, readTmpl.HtmlBody, "wrong html on readback")
	assert.Equal(t, newTmpl.TextBody, readTmpl.TextBody, "wrong text on readback")
	assert.Equal(t, newTmpl.Active, readTmpl.Active, "wrong active on readback")
	assert.Equal(t, newTmpl.TemplateType, readTmpl.TemplateType, "wrong type on readback")
	assert.Equal(t, newTmpl.LayoutTemplate, readTmpl.LayoutTemplate, "wrong layout on readback")

	// Delete
	err = c.DeleteTemplate(newTmpl.Alias)
	require.NoError(t, err, "got error deleting template")
}
