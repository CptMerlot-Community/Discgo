package tw

import (
	"github.com/nicklaw5/helix/v2"
)

type Client struct {
	client *helix.Client
}

// TODO: Figure twitch verification
func NewClient(opts *helix.Options) (*Client, error) {
	client, err := helix.NewClient(opts)
	if err != nil {
		return nil, err
	}

	return &Client{
		client: client,
	}, err

}
