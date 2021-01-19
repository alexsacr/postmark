package postmark

import (
	"fmt"
	"net/http"
)

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

func (c *Client) CreateSenderSignature(sendsig SenderSignature) (*SenderSignature, error) {
	return nil, nil
}

func (c *Client) UpdateSenderSignature(signatureID string, upsendsig SenderSignature) error {
	return nil
}

func (c *Client) DeleteSenderSignature(signatureID string) error {
	return nil
}
