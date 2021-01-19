package postmark

import (
	"fmt"
	"net/http"
)

// Server represents a Postmark server.
type Server struct {
	ID                         int      `json:"ID"`
	Name                       string   `json:"Name"`
	ApiTokens                  []string `json:"ApiTokens"`
	Color                      string   `json:"Color"`
	SmtpApiActivated           bool     `json:"SmtpApiActivated"`
	RawEmailEnabled            bool     `json:"RawEmailEnabled"`
	ServerLink                 string   `json:"ServerLink"`
	InboundAddress             string   `json:"InboundAddress"`
	InboundHookUrl             string   `json:"InboundHookUrl"`
	BounceHookUrl              string   `json:"BounceHookUrl"`
	OpenHookUrl                string   `json:"OpenHookUrl"`
	DeliveryHookUrl            string   `json:"DeliveryHookUrl"`
	PostFirstOpenOnly          bool     `json:"PostFirstOpenOnly"`
	InboundDomain              string   `json:"InboundDomain"`
	InboundHash                string   `json:"InboundHash"`
	InboundSpamThreshold       int      `json:"InboundSpamThreshold"`
	TrackOpens                 bool     `json:"TrackOpens"`
	TrackLinks                 string   `json:"TrackLinks"`
	IncludeBounceContentInHook bool     `json:"IncludeBounceContentInHook"`
	ClickHookUrl               string   `json:"ClickHookUrl"`
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
