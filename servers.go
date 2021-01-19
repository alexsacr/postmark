package postmark

import (
	"fmt"
	"net/http"
)

// Server represents a Postmark server.
type Server struct {
	ID                         int      `json:"ID,omitempty"`
	Name                       string   `json:"Name,omitempty"`
	APITokens                  []string `json:"ApiTokens,omitempty"`
	Color                      string   `json:"Color,omitempty"`
	SMTPAPIActivated           bool     `json:"SmtpApiActivated"`
	RawEmailEnabled            bool     `json:"RawEmailEnabled"`
	ServerLink                 string   `json:"ServerLink,omitempty"`
	InboundAddress             string   `json:"InboundAddress,omitempty"`
	InboundHookURL             string   `json:"InboundHookUrl,omitempty"`
	BounceHookURL              string   `json:"BounceHookUrl,omitempty"`
	OpenHookURL                string   `json:"OpenHookUrl,omitempty"`
	DeliveryHookURL            string   `json:"DeliveryHookUrl,omitempty"`
	PostFirstOpenOnly          bool     `json:"PostFirstOpenOnly"`
	InboundDomain              string   `json:"InboundDomain,omitempty"`
	InboundHash                string   `json:"InboundHash,omitempty"`
	InboundSpamThreshold       int      `json:"InboundSpamThreshold,omitempty"`
	TrackOpens                 bool     `json:"TrackOpens"`
	TrackLinks                 string   `json:"TrackLinks,omitempty"`
	IncludeBounceContentInHook bool     `json:"IncludeBounceContentInHook"`
	ClickHookURL               string   `json:"ClickHookUrl,omitempty"`
	EnableSMTPAPIErrorHooks    bool     `json:"EnableSmtpApiErrorHooks"`
}

// GetServer uses the `/servers/<id>` endpoint.  It requires a valid account
// token, not server token.
func (c *Client) GetServer(serverID string) (*Server, error) {
	var ret Server

	req := parameters{
		method:             http.MethodGet,
		endpoint:           fmt.Sprintf("/servers/%s", serverID),
		expectedStatusCode: http.StatusOK,
		respTarget:         &ret,
		token:              c.AccountToken,
	}

	err := c.doRequest(req)
	return &ret, err
}

// CreateServer creates a server based on the passed struct and returns a
// server struct with the ID populated.
func (c *Client) CreateServer(newsrvr Server) (*Server, error) {
	var ret Server

	req := parameters{
		method:             http.MethodPost,
		endpoint:           "/servers",
		expectedStatusCode: http.StatusOK,
		respTarget:         &ret,
		token:              c.AccountToken,
		payload:            newsrvr,
	}

	err := c.doRequest(req)
	return &ret, err
}

// UpdateServer updates a server based on the passed struct.  Boolean values
// will always be sent and *must* be set.
func (c *Client) UpdateServer(serverID string, upsrvr Server) error {
	req := parameters{
		method:             http.MethodPut,
		endpoint:           fmt.Sprintf("/servers/%s", serverID),
		expectedStatusCode: http.StatusOK,
		token:              c.AccountToken,
		payload:            upsrvr,
	}

	return c.doRequest(req)
}

// DeleteServer will delete a server based on the serverID.  Note that server
// deletion via API requires special access that requires a support ticket to
// Postmark.
func (c *Client) DeleteServer(serverID string) error {
	req := parameters{
		method:             http.MethodDelete,
		endpoint:           fmt.Sprintf("/servers/%s", serverID),
		expectedStatusCode: http.StatusOK,
		token:              c.AccountToken,
	}

	return c.doRequest(req)
}
