package fields

type TrendField string

const (
	TrendFieldTrendName  TrendField = "trend_name"
	TrendFieldTweetCount TrendField = "tweet_count"
)

func (f TrendField) String() string {
	return string(f)
}

type TrendFieldList []TrendField

func (fl TrendFieldList) FieldsName() string {
	return "trend.fields"
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

type PersonalizedTrendField string

const (
	PersonalizedTrendFieldCategory      PersonalizedTrendField = "category"
	PersonalizedTrendFieldPostCount     PersonalizedTrendField = "post_count"
	PersonalizedTrendFieldTrendName     PersonalizedTrendField = "trend_name"
	PersonalizedTrendFieldTrendingSince PersonalizedTrendField = "trending_since"
)

func (f PersonalizedTrendField) String() string {
	return string(f)
}

type PersonalizedTrendFieldList []PersonalizedTrendField

func (fl PersonalizedTrendFieldList) FieldsName() string {
	return "personalized_trend.fields"
}

func (fl PersonalizedTrendFieldList) Values() []string {
	if fl == nil {
		return []string{}
	}
	s := make([]string, 0)
	for _, f := range fl {
		s = append(s, f.String())
	}
	return s
}
