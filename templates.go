package postmark

import (
	"fmt"
	"net/http"
)

// Template represents a Postmark template.
type Template struct {
	Name               string      `json:"Name"`
	TemplateID         int         `json:"TemplateId"`
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
