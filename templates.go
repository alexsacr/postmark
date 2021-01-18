package postmark

import (
	"fmt"
	"net/http"
)

// Template represents a Postmark template.
type Template struct {
	Name               string      `json:"Name,omitempty"`
	TemplateID         int         `json:"TemplateId,omitempty"`
	Alias              string      `json:"Alias,omitempty"`
	Subject            string      `json:"Subject,omitempty"`
	HtmlBody           string      `json:"HtmlBody,omitempty"`
	TextBody           string      `json:"TextBody,omitempty"`
	AssociatedServerID int         `json:"AssociatedServerId,omitempty"`
	Active             bool        `json:"Active,omitempty"`
	TemplateType       string      `json:"TemplateType,omitempty"`
	LayoutTemplate     interface{} `json:"LayoutTemplate,omitempty"`
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

// CreateTemplate takes a populated Template struct and returns a template
// struct that includes the ID.
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

// UpdateTemplate accepts a templateID (or alias) and a template object
// with the values set you wish to change.
func (c *Client) UpdateTemplate(templateID string, upTmpl Template) error {
	req := parameters{
		method:             http.MethodPut,
		endpoint:           fmt.Sprintf("/templates/%s", templateID),
		expectedStatusCode: http.StatusOK,
		token:              c.ServerToken,
		payload:            upTmpl,
	}

	return c.doRequest(req)
}
