package gh

import (
	"context"
	"net/http"

	"github.com/google/go-github/v42/github"
)

type Client struct {
	client *github.Client
}

func CreateClient(h *http.Client) *Client {
	c := &Client{}
	c.client = github.NewClient(h)
	return c
}

func (c Client) GetUser(username string) (*github.User, error) {
	user, _, err := c.client.Users.Get(context.TODO(), username)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (c Client) CheckValidUser(username string) bool {
	_, err := c.GetUser(username)
	if err != nil {
		// check error if its a error user doesn't exist its a invalid user provided
		return false
	}
	return true
}

// c.client.Repositories.List(context.TODO(), username, nil)
