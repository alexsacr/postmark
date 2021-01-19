package postmark

import (
	"fmt"
	"net/http"
)

// SenderSignature represents a Postmark sender signature
type SenderSignature struct {
	Domain                        string `json:"Domain"`
	EmailAddress                  string `json:"EmailAddress"`
	ReplyToEmailAddress           string `json:"ReplyToEmailAddress"`
	Name                          string `json:"Name"`
	Confirmed                     bool   `json:"Confirmed"`
	SPFVerified                   bool   `json:"SPFVerified"`
	SPFHost                       string `json:"SPFHost"`
	SPFTextValue                  string `json:"SPFTextValue"`
	DKIMVerified                  bool   `json:"DKIMVerified"`
	WeakDKIM                      bool   `json:"WeakDKIM"`
	DKIMHost                      string `json:"DKIMHost"`
	DKIMTextValue                 string `json:"DKIMTextValue"`
	DKIMPendingHost               string `json:"DKIMPendingHost"`
	DKIMPendingTextValue          string `json:"DKIMPendingTextValue"`
	DKIMRevokedHost               string `json:"DKIMRevokedHost"`
	DKIMRevokedTextValue          string `json:"DKIMRevokedTextValue"`
	SafeToRemoveRevokedKeyFromDNS bool   `json:"SafeToRemoveRevokedKeyFromDNS"`
	DKIMUpdateStatus              string `json:"DKIMUpdateStatus"`
	ReturnPathDomain              string `json:"ReturnPathDomain"`
	ReturnPathDomainVerified      bool   `json:"ReturnPathDomainVerified"`
	ReturnPathDomainCNAMEValue    string `json:"ReturnPathDomainCNAMEValue"`
	ID                            int    `json:"ID"`
}

// SenderSignatureOpts is passed to CreateSenderSignature to create a new
// sender signature.
type SenderSignatureOpts struct {
	FromEmail        string `json:"FromEmail"`
	Name             string `json:"Name"`
	ReplyToEmail     string `json:"ReplyToEmail,omitempty"`
	ReturnPathDomain string `json:"ReturnPathDomain,omitempty"`
}

// GetSenderSignature returns a SenderSignature from an ID.
func (c *Client) GetSenderSignature(signatureID string) (*SenderSignature, error) {
	var ret SenderSignature

	req := parameters{
		method:             http.MethodGet,
		endpoint:           fmt.Sprintf("/senders/%s", signatureID),
		expectedStatusCode: http.StatusOK,
		respTarget:         &ret,
		token:              c.AccountToken,
	}

	err := c.doRequest(req)
	return &ret, err
}

// CreateSenderSignature accepts a SenderSignatureOpts struct and returns a
// SenderSignature struct with the generated ID.
func (c *Client) CreateSenderSignature(newSig SenderSignatureOpts) (*SenderSignature, error) {
	var ret SenderSignature

	req := parameters{
		method:             http.MethodPost,
		endpoint:           "/senders",
		expectedStatusCode: http.StatusOK,
		respTarget:         &ret,
		token:              c.AccountToken,
		payload:            newSig,
	}

	err := c.doRequest(req)
	return &ret, err
}

// UpdateSenderSignature accepts a SenderSignatureOpts struct and signature ID
// update.  Note that FromEmail cannot be updated.
func (c *Client) UpdateSenderSignature(signatureID string, upSig SenderSignatureOpts) error {
	req := parameters{
		method:             http.MethodPut,
		endpoint:           fmt.Sprintf("/senders/%s", signatureID),
		expectedStatusCode: http.StatusOK,
		token:              c.AccountToken,
		payload:            upSig,
	}

	return c.doRequest(req)
}

// DeleteSenderSignature will delete a signature by ID.
func (c *Client) DeleteSenderSignature(signatureID string) error {
	req := parameters{
		method:             http.MethodDelete,
		endpoint:           fmt.Sprintf("/senders/%s", signatureID),
		expectedStatusCode: http.StatusOK,
		token:              c.AccountToken,
	}

	return c.doRequest(req)
}
