package postmark

import (
	"fmt"
	"net/http"
)

// Server represents a Postmark server.
type Server struct {
	ID                         int      `json:"ID,omitempty"`
	Name                       string   `json:"Name,omitempty"`
	ApiTokens                  []string `json:"ApiTokens,omitempty"`
	Color                      string   `json:"Color,omitempty"`
	SmtpApiActivated           bool     `json:"SmtpApiActivated"`
	RawEmailEnabled            bool     `json:"RawEmailEnabled"`
	ServerLink                 string   `json:"ServerLink,omitempty"`
	InboundAddress             string   `json:"InboundAddress,omitempty"`
	InboundHookUrl             string   `json:"InboundHookUrl,omitempty"`
	BounceHookUrl              string   `json:"BounceHookUrl,omitempty"`
	OpenHookUrl                string   `json:"OpenHookUrl,omitempty"`
	DeliveryHookUrl            string   `json:"DeliveryHookUrl,omitempty"`
	PostFirstOpenOnly          bool     `json:"PostFirstOpenOnly"`
	InboundDomain              string   `json:"InboundDomain,omitempty"`
	InboundHash                string   `json:"InboundHash,omitempty"`
	InboundSpamThreshold       int      `json:"InboundSpamThreshold,omitempty"`
	TrackOpens                 bool     `json:"TrackOpens"`
	TrackLinks                 string   `json:"TrackLinks,omitempty"`
	IncludeBounceContentInHook bool     `json:"IncludeBounceContentInHook"`
	ClickHookUrl               string   `json:"ClickHookUrl,omitempty"`
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
