package types_test

import (
	"testing"

	"github.com/Arhius/gotwi/fields"
	"github.com/Arhius/gotwi/trend/types"
	"github.com/stretchr/testify/assert"
)

func Test_ListTrend_SetAccessToken(t *testing.T) {
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

func Test_ListTrend_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id"

	cases := []struct {
		name   string
		params *types.ListInput
		expect string
	}{
		{
			name: "only required parameter",
			params: &types.ListInput{
				ID: "lid",
			},
			expect: endpointRoot + "lid",
		},
		{
			name: "with fields",
			params: &types.ListInput{
				ID:     "lid",
				Fields: fields.TrendFieldList{"ex1", "ex2"},
			},
			expect: endpointRoot + "lid" + "?trend.fields=ex1%2Cex2",
		},
		{
			name: "with max_trends",
			params: &types.ListInput{
				ID:        "lid",
				MaxTrends: 10,
			},
			expect: endpointRoot + "lid" + "?max_trends=10",
		},
		{
			name: "all query parameters",
			params: &types.ListInput{
				ID:        "lid",
				MaxTrends: 10,
				Fields:    fields.TrendFieldList{"ex1"},
			},
			expect: endpointRoot + "lid" + "?max_trends=10&trend.fields=ex1",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			ep := c.params.ResolveEndpoint(endpointBase)
			assert.Equal(tt, c.expect, ep)
		})
	}
}

func Test_ListTrend_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.ListInput
	}{
		{
			name:   "empty params",
			params: &types.ListInput{},
		},
		{
			name:   "some params",
			params: &types.ListInput{ID: "sid"},
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
