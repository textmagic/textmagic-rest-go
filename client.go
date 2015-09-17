package textmagic

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

const baseURL = "https://rest.textmagic.com/api/v2"

var (
	httpClient = &http.Client{}        // Re-usable HTTP client
	emptyData  = url.Values{}.Encode() // Cache empty data request
)

// Client represents a API client.
type Client struct {
	username string
	token    string
	baseURL  string
}

// NewClient creates returns a client for the given
// username / token pair.
func NewClient(username, token string) *Client {
	return &Client{username, token, baseURL}
}

// SetBaseURL sets the API base URL.
func (c *Client) SetBaseURL(u string) {
	c.baseURL = u
}

// Request makes an API request, automatically decoding
// the JSON payload for responses returning objects.
func (c *Client) Request(method, uri string, p, d url.Values, dst interface{}) error {
	var payload *strings.Reader

	if d != nil {
		payload = strings.NewReader(d.Encode())
	} else {
		payload = strings.NewReader(emptyData)
	}

	if p != nil {
		uri += "?" + p.Encode()
	}

	req, err := http.NewRequest(method, c.baseURL+"/"+uri, payload)

	if err != nil {
		return err
	}

	if method != "GET" && method != "HEAD" {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	req.Header.Add("Accept-Charset", "utf-8")
	req.Header.Add("Accept-Language", "en-us")
	// To avoid Header.Add key capitalization.
	req.Header["X-TM-Username"] = []string{c.username}
	req.Header["X-TM-Key"] = []string{c.token}

	resp, err := httpClient.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 && resp.StatusCode != 201 && resp.StatusCode != 204 {
		var e *Error

		if err = json.NewDecoder(resp.Body).Decode(&e); err != nil {
			return err
		}

		return e
	} else if method == "DELETE" {
		if resp.StatusCode != 204 {
			return NewError(resp.StatusCode, "unexpected status code for DELETE")
		}

		return nil
	} else if resp.StatusCode == 204 {
		return nil
	}

	return json.NewDecoder(resp.Body).Decode(dst)
}

func (c *Client) get(uri string, params, data url.Values, dst interface{}) error {
	return c.Request("GET", uri, params, data, dst)
}

func (c *Client) post(uri string, params, data url.Values, dst interface{}) error {
	return c.Request("POST", uri, params, data, dst)
}

func (c *Client) put(uri string, params, data url.Values, dst interface{}) error {
	return c.Request("PUT", uri, params, data, dst)
}

func (c *Client) delete(uri string, params, data url.Values, dst interface{}) error {
	return c.Request("DELETE", uri, params, data, nil)
}

// Ping sends a ping request to the API to test credentials.
func (c *Client) Ping() error {
	var p *struct {
		Ping string `json:"ping"`
	}

	if err := c.get("ping", nil, nil, &p); err != nil {
		return err
	} else if p == nil || p.Ping != "pong" {
		return ErrPing
	}

	return nil
}
