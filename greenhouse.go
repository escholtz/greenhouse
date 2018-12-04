// Package greenhouse provides a wrapper around the greenhouse.io Job Board API.
// The API is defined at: https://developers.greenhouse.io/job-board.html
package greenhouse

import (
	"encoding/json"
	"net/http"
)

// A Client manages communication with the Greenhouse API.
type Client struct {
	baseURL string
	client  *http.Client
}

// NewClient returns a new Greenhouse API client.
func NewClient() *Client {
	return &Client{
		baseURL: "https://boards-api.greenhouse.io",
		client:  http.DefaultClient,
	}
}

// WithHTTPClient sets a custom http client to be used for all future requests.
func (c *Client) WithHTTPClient(client *http.Client) *Client {
	if client != nil {
		c.client = client
	}
	return c
}

// Board describes an organization's job board. Usually it includes a company
// name and sometimes a description of the organization as html.
type Board struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type apiError struct {
	Status  int    `json:"status"`
	Message string `json:"error"`
}

func (m *apiError) Error() string {
	return m.Message
}

func (c *Client) getJSON(url string, out interface{}) error {
	resp, err := c.client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		m := &apiError{}
		err := json.NewDecoder(resp.Body).Decode(m)
		if err != nil {
			return err
		}
		return m
	}

	// Also check content type?

	return json.NewDecoder(resp.Body).Decode(out)
}

// Board returns the organization's name and optionally an html description of
// the organization.
func (c *Client) Board(boardToken string) (*Board, error) {
	board := &Board{}
	url := c.baseURL + "/v1/boards/" + boardToken
	err := c.getJSON(url, board)
	return board, err
}

// Jobs describes an organization's job listings.
type Jobs struct {
	Jobs []struct {
		ID            int    `json:"id"`
		InternalJobID int    `json:"internal_job_id"`
		Title         string `json:"title"`
		Education     string `json:"education,omitempty"`
		UpdatedAt     string `json:"updated_at"`
		Location      struct {
			Name string `json:"name"`
		} `json:"location"`
		AbsoluteURL string `json:"absolute_url"`
		// TODO: varies by company, some use []string for Value, some
		// use string
		// Metadata    []struct {
		// 	ID        int      `json:"id"`
		// 	Name      string   `json:"name"`
		// 	Value     []string `json:"value"`
		// 	ValueType string   `json:"value_type"`
		// } `json:"metadata"`
		Content     string `json:"content"`
		Departments []struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			ChildIds []int  `json:"child_ids"`
			ParentID int    `json:"parent_id"`
		} `json:"departments"`
		Offices []struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			Location string `json:"location"`
			ChildIds []int  `json:"child_ids"`
			ParentID int    `json:"parent_id"`
		} `json:"offices"`
	} `json:"jobs"`
	Meta struct {
		Total int `json:"total"`
	} `json:"meta"`
}

// Jobs returns the job openings, including content, for the specified
// organization.
func (c *Client) Jobs(boardToken string) (*Jobs, error) {
	out := &Jobs{}
	url := c.baseURL + "/v1/boards/" + boardToken + "/jobs?content=true"
	err := c.getJSON(url, out)
	return out, err
}
