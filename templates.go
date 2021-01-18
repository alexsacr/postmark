package postmark

import (
	"fmt"
	"net/http"
)

// Template represents a Postmark template.
type Template struct {
	Name               string      `json:"Name"`
	TemplateID         int         `json:"TemplateId,omitempty"`
	Alias              string      `json:"Alias"`
	Subject            string      `json:"Subject"`
	HtmlBody           string      `json:"HtmlBody"`
	TextBody           string      `json:"TextBody"`
	AssociatedServerID int         `json:"AssociatedServerId"`
	Active             bool        `json:"Active"`
	TemplateType       string      `json:"TemplateType"`
	LayoutTemplate     interface{} `json:"LayoutTemplate"`
}

// GetTemplate uses the `/templates/<id>` endpoint.  It requires a valid server
// token.
func (c *Client) GetTemplate(templateID string) (*Template, error) {
	var ret Template

	req := parameters{
		method:             http.MethodGet,
		endpoint:           fmt.Sprintf("/templates/%s", templateID),
		expectedStatusCode: http.StatusOK,
		respTarget:         &ret,
		token:              c.ServerToken,
	}

	err := c.doRequest(req)
	return &ret, err
}

// CreateTemplate takes a populated Template struct and returns a complete
// version of the template (including ID) by POSTing to the /templates
// endpoint.
func (c *Client) CreateTemplate(newTmpl Template) (*Template, error) {
	var ret Template

	req := parameters{
		method:             http.MethodPost,
		endpoint:           "/templates",
		expectedStatusCode: http.StatusOK,
		respTarget:         &ret,
		token:              c.ServerToken,
		payload:            newTmpl,
	}

	err := c.doRequest(req)
	return &ret, err
}

// DeleteTemplate will delete a template based on its ID.  Alias is also
// acceptable as an identifier.
func (c *Client) DeleteTemplate(templateID string) error {
	req := parameters{
		method:             http.MethodDelete,
		endpoint:           fmt.Sprintf("/templates/%s", templateID),
		expectedStatusCode: http.StatusOK,
		token:              c.ServerToken,
	}

	return c.doRequest(req)
}
