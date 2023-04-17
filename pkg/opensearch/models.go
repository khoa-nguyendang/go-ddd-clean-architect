package opensearch

// Document represents a single Document in Get API response body.
type Document[T any] struct {
	Source T       `json:"_source"`
	Id     string  `json:"_id"`
	Score  float64 `json:"_score"`
}

type MultiMatch struct {
	Query  string   `json:"query"`
	Fields []string `json:"fields"`
}
type Query struct {
	MultiMatch MultiMatch `json:"multi_match"`
}

type SearchQuery struct {
	From  int   `json:"from"`
	Size  int   `json:"size"`
	Query Query `json:"query"`
}

type SearchResponse[T any] struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Hits     struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore float64       `json:"max_score"`
		Hits     []Document[T] `json:"hits"`
	} `json:"hits"`
}
