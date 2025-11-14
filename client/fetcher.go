package client

import "github.com/gravwell/gravwell/v4/client/types"

func (c *Client) GetFetchers() (r []types.Fetcher, err error) {
	err = c.getStaticURL(fetchersUrl(), &r)
	return
}
