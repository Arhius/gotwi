package follow

import (
	"context"

	"github.com/Arhius/gotwi"
	"github.com/Arhius/gotwi/user/follow/types"
)

const (
	listFollowingsEndpoint  = "https://api.twitter.com/2/users/:id/following"
	listFollowersEndpoint   = "https://api.twitter.com/2/users/:id/followers"
	createFollowingEndpoint = "https://api.twitter.com/2/users/:id/following"
	deleteFollowingEndpoint = "https://api.twitter.com/2/users/:source_user_id/following/:target_user_id"
)

// Returns a list of users the specified user ID is following.
// https://developer.twitter.com/en/docs/twitter-api/users/follows/api-reference/get-users-id-following
func ListFollowings(ctx context.Context, c *gotwi.Client, p *types.ListFollowingsInput) (*types.ListFollowingsOutput, error) {
	res := &types.ListFollowingsOutput{}
	if err := c.CallAPI(ctx, listFollowingsEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Returns a list of users who are followers of the specified user ID.
// https://developer.twitter.com/en/docs/twitter-api/users/follows/api-reference/get-users-id-followers
func ListFollowers(ctx context.Context, c *gotwi.Client, p *types.ListFollowersInput) (*types.ListFollowersOutput, error) {
	res := &types.ListFollowersOutput{}
	if err := c.CallAPI(ctx, listFollowersEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Allows a user ID to follow another user.
// If the target user does not have public Tweets, this endpoint will send a follow request.
// The request succeeds with no action when the authenticated user sends a request to a user
// they're already following, or if they're sending a follower request to a user that does not have public Tweets.
// https://developer.twitter.com/en/docs/twitter-api/users/follows/api-reference/post-users-source_user_id-following
func CreateFollowing(ctx context.Context, c *gotwi.Client, p *types.CreateFollowingInput) (*types.CreateFollowingOutput, error) {
	res := &types.CreateFollowingOutput{}
	if err := c.CallAPI(ctx, createFollowingEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Allows a user ID to unfollow another user.
// The request succeeds with no action when the authenticated user sends a request to a user they're not following or have already unfollowed.
// https://developer.twitter.com/en/docs/twitter-api/users/follows/api-reference/delete-users-source_id-following
func DeleteFollowing(ctx context.Context, c *gotwi.Client, p *types.DeleteFollowingInput) (*types.DeleteFollowingOutput, error) {
	res := &types.DeleteFollowingOutput{}
	if err := c.CallAPI(ctx, deleteFollowingEndpoint, "DELETE", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
