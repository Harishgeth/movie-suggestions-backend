package dtos

import (
	"encoding/json"
	"movie-suggestions-api/config"
	"time"
)

type ScrollDataCaptured struct {
	UserID           string    `json:"user_id"`
	Timestamp        time.Time `json:"timestamp"`
	DurationOfScroll int       `json:"duration_of_scroll"`
	PostID           string    `json:"post_id"`
}

type ScrollDataCapturedElasticFormat struct {
	UserID           string `json:"user_id"`
	PostID           string `json:"post_id"`
	DurationOfScroll int    `json:"duration_of_scroll"`
	Timestamp        string `json:"timestamp"`
}

// MarshalJSON customizes JSON marshaling for the ScrollDataCaptured struct.
func (s *ScrollDataCaptured) MarshalJSON() ([]byte, error) {
	type Alias ScrollDataCaptured // Create an alias to avoid infinite recursion.
	return json.Marshal(&struct {
		*Alias
		Timestamp string `json:"timestamp"`
	}{
		Alias:     (*Alias)(s),
		Timestamp: s.Timestamp.Format(time.RFC3339),
	})
}

// UnmarshalJSON customizes JSON unmarshaling for the ScrollDataCaptured struct.
func (s *ScrollDataCaptured) UnmarshalJSON(data []byte) error {
	type Alias ScrollDataCaptured // Create an alias to avoid infinite recursion.
	aux := &struct {
		*Alias
		Timestamp string `json:"timestamp"`
	}{
		Alias: (*Alias)(s),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	timestamp, err := time.Parse(time.RFC3339, aux.Timestamp)
	if err != nil {
		return err
	}
	s.Timestamp = timestamp
	return nil
}

func (s *ScrollDataCaptured) CopyForElasticSearch() *ScrollDataCapturedElasticFormat {
	return &ScrollDataCapturedElasticFormat{
		UserID:           s.UserID,
		PostID:           s.PostID,
		DurationOfScroll: s.DurationOfScroll,
		Timestamp:        s.Timestamp.Format(config.ES_TIME_FORMAT),
	}

}
