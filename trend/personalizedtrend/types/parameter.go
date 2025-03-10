package types

import (
	"github.com/Arhius/gotwi/fields"
	"github.com/Arhius/gotwi/internal/util"
	"io"
)

type ListInput struct {
	accessToken string
	Fields      fields.TrendFieldList
}

var listQueryParameters = map[string]struct{}{
	"personalized_trend.fields": {},
}

func (p *ListInput) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *ListInput) AccessToken() string {
	return p.accessToken
}

func (p *ListInput) ResolveEndpoint(endpointBase string) string {
	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, listQueryParameters)
		endpointBase += "?" + qs
	}

	return endpointBase
}

func (p *ListInput) Body() (io.Reader, error) {
	return nil, nil
}

func (p *ListInput) ParameterMap() map[string]string {
	m := map[string]string{}
	m = fields.SetFieldsParams(m, p.Fields)
	return m
}
