package f3sjira

import (
	"errors"
	"net/http"

	"github.com/andygrunwald/go-jira"
)

type SearchOptions = jira.SearchOptions

func NewDefaultSearchOptions() *SearchOptions {
	return &SearchOptions{
		StartAt:    0,
		MaxResults: 50,
		Expand:     "",
	}
}

type JiraClient struct {
	config JiraConfig
	client *jira.Client
}

func NewJiraClient(config JiraConfig) JiraClient {
	return JiraClient{
		config,
		nil,
	}
}

func (c *JiraClient) IsConnected() bool {
	return c.client != nil
}

func (c *JiraClient) Connect() error {

	if c.IsConnected() {
		return errors.New("already connected")
	}

	if client, err := jira.NewClient(getHttpClient(&c.config), c.config.Url); err == nil {
		c.client = client
		return nil
	} else {
		return err
	}
}

func (c *JiraClient) Disconnect() {
	c.client = nil
}

func (c *JiraClient) Search(jql string, options *SearchOptions, f func(jira.Issue) error) error {
	return c.client.Issue.SearchPages(jql, options, f)
}

func getHttpClient(config *JiraConfig) *http.Client {
	if config == nil {
		return http.DefaultClient
	} else {
		if config.Auth == nil {
			return http.DefaultClient
		} else {
			tp := jira.BasicAuthTransport{
				Username: config.Auth.Usr,
				Password: string(config.Auth.Pwd),
			}
			return tp.Client()
		}
	}
}
