package resources

type Trend struct {
	Category      *string `json:"category"`
	PostCount     *int    `json:"post_count"`
	TrendName     *string `json:"trend_name"`
	TrendingSince *string `json:"trending_since"`
}
