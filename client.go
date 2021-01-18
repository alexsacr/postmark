package postmark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	// PostmarkAPIAddr is the default, production target for API requests.
	PostmarkAPIAddr = "https://api.postmarkapp.com"
)

// Client provides access to Postmark API endpoints in this module.
type Client struct {
	APIAddr      string
	AccountToken string
	ServerToken  string
}

type apiError struct {
	ErrorCode int    `json:"ErrorCode"`
	Message   string `json:"Message"`
}

type parameters struct {
	method             string
	endpoint           string
	expectedStatusCode int
	respTarget         interface{}
	token              string
	payload            interface{}
}

// NewClient returns a new client configured for use.
func NewClient(apiToken string, serverToken string) *Client {
	return &Client{
		APIAddr:      PostmarkAPIAddr,
		AccountToken: apiToken,
		ServerToken:  serverToken,
	}
}

func (c *Client) doRequest(opts parameters) error {
	url := fmt.Sprintf("%s%s", c.APIAddr, opts.endpoint)

	req, err := http.NewRequest(opts.method, url, nil)
	if err != nil {
		return err
	}

	if opts.payload != nil {
		reqBody, err := json.Marshal(opts.payload)
		if err != nil {
			return err
		}
		req.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	if opts.token == c.AccountToken {
		req.Header.Add("X-Postmark-Account-Token", c.AccountToken)
	}
	if opts.token == c.ServerToken {
		req.Header.Add("X-Postmark-Server-Token", c.ServerToken)
	}

	httpC := &http.Client{}

	resp, err := httpC.Do(req)
	if err != nil {
		return err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Handle API errors
	if resp.StatusCode != opts.expectedStatusCode {
		var apiErr apiError
		err = json.Unmarshal(body, &apiErr)
		if err != nil {
			return err
		}

		return fmt.Errorf("api error code %d: %s", apiErr.ErrorCode, apiErr.Message)
	}

	if opts.respTarget == nil {
		// Short-circuit if we're not actually looking for the response
		return nil
	}

	err = json.Unmarshal(body, opts.respTarget)
	if err != nil {
		return err
	}

	return nil
}
