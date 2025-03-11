package trend

import (
	"context"
	"github.com/Arhius/gotwi"
	"github.com/Arhius/gotwi/user/trend/types"
)

const listEndpoint = "https://api.twitter.com/2/users/personalized_trends"

// Returns Personalized trends for the authenticated user.
// https://docs.x.com/x-api/users/personalized-trends
func List(ctx context.Context, c *gotwi.Client, p *types.ListInput) (*types.ListOutput, error) {
	res := &types.ListOutput{}
	if err := c.CallAPI(ctx, listEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
