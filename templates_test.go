package postmark

import (
	"fmt"
	"os"
	"testing"
	"time"

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

func TestCreateTemplateIntegrationOK(t *testing.T) {
	if os.Getenv("INTEGRATION") != "true" {
		t.Skip()
	}

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
}
