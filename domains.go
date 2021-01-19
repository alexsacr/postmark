package postmark

import (
	"fmt"
	"net/http"
)

// Domain represents a Postmark domain.
type Domain struct {
	ID                            int    `json:"ID,omitempty"`
	Name                          string `json:"Name,omitempty"`
	ReturnPathDomain              string `json:"ReturnPathDomain,omitempty"`
	ReturnPathDomainVerified      bool   `json:"ReturnPathDomainVerified,omitempty"`
	ReturnPathDomainCNAMEValue    string `json:"ReturnPathDomainCNAMEValue,omitempty"`
	SPFHost                       string `json:"SPFHost,omitempty"`
	SPFTextValue                  string `json:"SPFTextValue,omitempty"`
	DKIMVerified                  bool   `json:"DKIMVerified,omitempty"`
	WeakDKIM                      bool   `json:"WeakDKIM,omitempty"`
	DKIMHost                      string `json:"DKIMHost,omitempty"`
	DKIMTextValue                 string `json:"DKIMTextValue,omitempty"`
	DKIMPendingHost               string `json:"DKIMPendingHost,omitempty"`
	DKIMPendingTextValue          string `json:"DKIMPendingTextValue,omitempty"`
	DKIMRevokedHost               string `json:"DKIMRevokedHost,omitempty"`
	DKIMRevokedTextValue          string `json:"DKIMRevokedTextValue,omitempty"`
	SafeToRemoveRevokedKeyFromDNS bool   `json:"SafeToRemoveRevokedKeyFromDNS,omitempty"`
	DKIMUpdateStatus              string `json:"DKIMUpdateStatus,omitempty"`
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

// UpdateDomain updates a domain based on the passed struct.
func (c *Client) UpdateDomain(domainID string, upDomain Domain) error {
	req := parameters{
		method:             http.MethodPut,
		endpoint:           fmt.Sprintf("/domains/%s", domainID),
		expectedStatusCode: http.StatusOK,
		token:              c.AccountToken,
		payload:            upDomain,
	}

	return c.doRequest(req)
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
