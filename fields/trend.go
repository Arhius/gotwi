package fields

type TrendField string

const (
	TrendFieldCategory      TrendField = "category"
	TrendFieldPostCount     TrendField = "post_count"
	TrendFieldTrendName     TrendField = "trend_name"
	TrendFieldTrendingSince TrendField = "trending_since"
)

func (f TrendField) String() string {
	return string(f)
}

type TrendFieldList []TrendField

func (fl TrendFieldList) FieldsName() string {
	return "personalized_trend.fields"
}

func (fl TrendFieldList) Values() []string {
	if fl == nil {
		return []string{}
	}
	s := make([]string, 0)
	for _, f := range fl {
		s = append(s, f.String())
	}
	return s
}
