package searchspace

import (
	"context"

	"github.com/Arhius/gotwi"
	"github.com/Arhius/gotwi/space/searchspace/types"
)

const (
	listEndpoint = "https://api.twitter.com/2/spaces/search"
)

// Return live or scheduled Spaces matching your specified search terms.
// This endpoint performs a keyword search, meaning that it will return Spaces
// that are an exact case-insensitive match of the specified search term. The search term will match the original title of the Space.
// https://developer.twitter.com/en/docs/twitter-api/spaces/search/api-reference/get-spaces-search
func List(ctx context.Context, c *gotwi.Client, p *types.ListInput) (*types.ListOutput, error) {
	res := &types.ListOutput{}
	if err := c.CallAPI(ctx, listEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
