package resources

type Trend struct {
	TrendName  *string `json:"trend_name"`
	TweetCount *int    `json:"tweet_count"`
}

type PersonalizedTrend struct {
	Category      *string `json:"category"`
	PostCount     *string `json:"post_count"`
	TrendName     *string `json:"trend_name"`
	TrendingSince *string `json:"trending_since"`
}
