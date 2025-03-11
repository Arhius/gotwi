package types

import (
	"github.com/Arhius/gotwi/fields"
	"github.com/Arhius/gotwi/internal/util"
	"io"
	"net/url"
	"strconv"
	"strings"
)

type ListMaxTrends int

func (m ListMaxTrends) Valid() bool {
	return m > 1 && m <= 50
}

func (m ListMaxTrends) String() string {
	return strconv.Itoa(int(m))
}

type ListInput struct {
	accessToken string

	// Path parameter
	ID string // List ID

	// Query parameters
	MaxTrends ListMaxTrends
	Fields    fields.TrendFieldList
}

var listQueryParameters = map[string]struct{}{
	"max_trends":   {},
	"trend.fields": {},
}

func (p *ListInput) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *ListInput) AccessToken() string {
	return p.accessToken
}

func (p *ListInput) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	encoded := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", encoded, 1)

	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, listQueryParameters)
		endpoint += "?" + qs
	}

	return endpoint
}

func (p *ListInput) Body() (io.Reader, error) {
	return nil, nil
}

func (p *ListInput) ParameterMap() map[string]string {
	m := map[string]string{}

	if p.MaxTrends.Valid() {
		m["max_trends"] = p.MaxTrends.String()
	}

	m = fields.SetFieldsParams(m, p.Fields)

	return m
}
