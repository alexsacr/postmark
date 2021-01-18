package postmark

import (
	"fmt"
	"net/http"
)

// Server represents a Postmark server.
type Server struct {
	ID                         int      `json:"ID"`
	Name                       string   `json:"Name"`
	APITokens                  []string `json:"ApiTokens"`
	Color                      string   `json:"Color"`
	SMTPAPIActivated           bool     `json:"SmtpApiActivated"`
	RawEmailEnabled            bool     `json:"RawEmailEnabled"`
	ServerLink                 string   `json:"ServerLink"`
	InboundAddress             string   `json:"InboundAddress"`
	InboundHookURL             string   `json:"InboundHookUrl"`
	BounceHookURL              string   `json:"BounceHookUrl"`
	OpenHookURL                string   `json:"OpenHookUrl"`
	DeliveryHookURL            string   `json:"DeliveryHookUrl"`
	PostFirstOpenOnly          bool     `json:"PostFirstOpenOnly"`
	InboundDomain              string   `json:"InboundDomain"`
	InboundHash                string   `json:"InboundHash"`
	InboundSpamThreshold       int      `json:"InboundSpamThreshold"`
	TrackOpens                 bool     `json:"TrackOpens"`
	TrackLinks                 string   `json:"TrackLinks"`
	IncludeBounceContentInHook bool     `json:"IncludeBounceContentInHook"`
	ClickHookURL               string   `json:"ClickHookUrl"`
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
