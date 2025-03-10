package types_test

import (
	"testing"

	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/trend/personalizedtrend/types"
	"github.com/stretchr/testify/assert"
)

func Test_PersonalizedTrend_SetAccessToken(t *testing.T) {
	cases := []struct {
		name   string
		token  string
		expect string
	}{
		{
			name:   "normal",
			token:  "test-token",
			expect: "test-token",
		},
		{
			name:   "empty",
			token:  "",
			expect: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			p := &types.ListInput{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_PersonalizedTrend_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"

	cases := []struct {
		name   string
		params *types.ListInput
		expect string
	}{
		{
			name:   "empty params",
			params: &types.ListInput{},
			expect: endpointRoot,
		},
		{
			name: "with fields",
			params: &types.ListInput{
				Fields: fields.TrendFieldList{"category", "post_count"},
			},
			expect: endpointRoot + "?personalized_trend.fields=category%2Cpost_count",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			ep := c.params.ResolveEndpoint(endpointRoot)
			assert.Equal(tt, c.expect, ep)
		})
	}
}

func Test_PersonalizedTrend_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.ListInput
	}{
		{
			name:   "empty params",
			params: &types.ListInput{},
		},
		{
			name: "some params",
			params: &types.ListInput{
				Fields: fields.TrendFieldList{"category", "post_count"},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			r, err := c.params.Body()
			assert.NoError(tt, err)
			assert.Nil(tt, r)
		})
	}
}
