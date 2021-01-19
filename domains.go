package postmark

import (
	"fmt"
	"net/http"
)

// Domain represents a Postmark domain.
type Domain struct {
	ID               int    `json:"ID,omitempty"`
	Name             string `json:"Name,omitempty"`
	ReturnPathDomain string `json:"ReturnPathDomain,omitempty"`
}

// GetDomain uses the `/domains/<id>` endpoint.  It requires a valid account token.
func (c *Client) GetDomain(domainID string) (*Domain, error) {
	var ret Domain

	req := parameters{
		method:             http.MethodGet,
		endpoint:           fmt.Sprintf("/domains/%s", domainID),
		expectedStatusCode: http.StatusOK,
		respTarget:         &ret,
		token:              c.AccountToken,
	}

	err := c.doRequest(req)
	return &ret, err
}

// CreateDomain creates a domain based on the passed struct and returns a
// domain struct with the ID populated.
func (c *Client) CreateDomain(newDomain Domain) (*Domain, error) {
	var ret Domain

	req := parameters{
		method:             http.MethodPost,
		endpoint:           "/domains",
		expectedStatusCode: http.StatusOK,
		respTarget:         &ret,
		token:              c.AccountToken,
		payload:            newDomain,
	}

	err := c.doRequest(req)
	return &ret, err
}

// DeleteDomain will delete a domain based on the domainID.
func (c *Client) DeleteDomain(domainID string) error {
	req := parameters{
		method:             http.MethodDelete,
		endpoint:           fmt.Sprintf("/domains/%s", domainID),
		expectedStatusCode: http.StatusOK,
		token:              c.AccountToken,
	}

	return c.doRequest(req)
}
