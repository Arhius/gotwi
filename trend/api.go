package trend

import (
	"context"
	"github.com/Arhius/gotwi"
	"github.com/Arhius/gotwi/trend/types"
)

const listEndpoint = "https://api.x.com/2/trends/by/woeid/:id"

// Returns trends list.
// https://docs.x.com/x-api/trends/trends
func List(ctx context.Context, c *gotwi.Client, p *types.ListInput) (*types.ListOutput, error) {
	res := &types.ListOutput{}
	if err := c.CallAPI(ctx, listEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
