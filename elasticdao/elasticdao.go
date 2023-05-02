package elasticDao

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"movie-suggestions-api/utils/log"

	"movie-suggestions-api/config"
	"movie-suggestions-api/dtos"
	"net/http"
)

type MovieElasticDao interface {
	CreateMovieIndexIfNotExists() error
	PutDataInElasticSearch(scrollData *dtos.ScrollDataCaptured) error
	GetSuggestionsPageSortOrder(userID string, page, size int) ([]string, error)
	GetTrendingPageSortOrder(page, size int) ([]string, error)
}

type moviesElasticOps struct {
	l *log.Logger
}

func GetMovieDao(l *log.Logger) MovieElasticDao {
	return &moviesElasticOps{
		l: l,
	}
}

func (m *moviesElasticOps) CreateMovieIndexIfNotExists() error {
	// Elasticsearch URL and index name
	esURL := config.ELASTIC_URI
	indexName := config.MOVIE_INDEX

	// Check if index exists
	resp, err := http.Head(fmt.Sprintf("%s/%s", esURL, indexName))
	if err != nil {
		m.l.Error("Error converting index mapping to JSON:", err)
		return err
	}

	// If index does not exist, create it
	if resp.StatusCode == http.StatusNotFound {
		// Define index mapping
		indexMapping := map[string]interface{}{
			"mappings": map[string]interface{}{
				"properties": map[string]interface{}{
					"user_id": map[string]interface{}{
						"type": "keyword",
					},
					"timestamp": map[string]interface{}{
						"type":   "date",
						"format": "yyyy-MM-dd HH:mm:ss",
					},
					"duration_of_scroll": map[string]interface{}{
						"type": "integer",
					},
					"post_id": map[string]interface{}{
						"type": "keyword",
					},
				},
			},
		}

		// Convert mapping to JSON bytes
		indexMappingBytes, err := json.Marshal(indexMapping)
		if err != nil {
			m.l.Info("Error converting index mapping to JSON:", err)
			return err
		}
		m.l.Info("Elasticsearch specifics: ", esURL, indexName, fmt.Sprintf("%s/%s", esURL, indexName))

		// Send PUT request to create index with mapping
		req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/%s", esURL, indexName), bytes.NewBuffer(indexMappingBytes))
		if err != nil {
			m.l.Error("Error creating request:", err)
			return err
		}
		req.Header.Set("Content-Type", "application/json")
		_, err = http.DefaultClient.Do(req)
		if err != nil {
			m.l.Error("Error creating index:", err)
			return err
		}
		m.l.Info("Index created successfully")
	}
	return nil
}

func (m *moviesElasticOps) PutDataInElasticSearch(scrollData *dtos.ScrollDataCaptured) error {
	// Create an instance of the ScrollDataCaptured struct
	esURL := config.ELASTIC_URI
	indexName := config.MOVIE_INDEX

	// Convert struct to JSON bytes
	scrollDataBytes, err := json.Marshal(scrollData.CopyForElasticSearch())
	if err != nil {
		m.l.Error("Error converting struct to JSON:", err)
		return err
	}
	m.l.Info("Elasticsearch specifics: ", esURL, indexName, fmt.Sprintf("%s/%s", esURL, indexName), string(scrollDataBytes))

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/%s/_doc", esURL, indexName), bytes.NewBuffer(scrollDataBytes))
	if err != nil {
		m.l.Error("Error creating request:", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		m.l.Error("Error adding document to index:", err)
		return err
	}
	defer resp.Body.Close()

	// Check response status code
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		m.l.Info("Document added to index successfully")
	} else {
		errBody, err := io.ReadAll(resp.Body)
		if err != nil {
			m.l.Error("Unable to read the body of failure to insert reason to elasticsearch")
		}
		m.l.Error("Error adding document to index:", resp.StatusCode, string(errBody))
		return err
	}
	return nil

}

func (m *moviesElasticOps) GetSuggestionsPageSortOrder(userID string, page, size int) ([]string, error) {
	// Define query parameters
	query := map[string]interface{}{
		"from": (page - 1) * size,
		"size": size,
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"filter": map[string]interface{}{
					"term": map[string]interface{}{
						"user_id": userID,
					},
				},
				"must": map[string]interface{}{
					"range": map[string]interface{}{
						"timestamp": map[string]interface{}{
							"gte": "now-15d",
						},
					},
				},
			},
		},
		"aggs": map[string]interface{}{
			"post": map[string]interface{}{
				"terms": map[string]interface{}{
					"field": "post_id",
					"order": map[string]interface{}{
						"avg_duration": "desc",
					},
				},
				"aggs": map[string]interface{}{
					"avg_duration": map[string]interface{}{
						"avg": map[string]interface{}{
							"field": "duration_of_scroll",
						},
					},
				},
			},
		},
	}

	// Encode query
	queryBytes, err := json.Marshal(query)
	if err != nil {
		return nil, fmt.Errorf("error encoding Elasticsearch query: %v", err)
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s/_search", config.ELASTIC_URI, config.MOVIE_INDEX), bytes.NewBuffer(queryBytes))
	if err != nil {
		return nil, fmt.Errorf("error creating Elasticsearch request: %v", err)
	}

	// Set request headers
	req.Header.Set("Content-Type", "application/json")

	// Send request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error executing Elasticsearch request: %v", err)
	}
	defer res.Body.Close()

	// Decode response
	var response map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("error decoding Elasticsearch response: %v", err)
	}

	// Extract post IDs
	var postIDs []string
	for _, bucket := range response["aggregations"].(map[string]interface{})["post"].(map[string]interface{})["buckets"].([]interface{}) {
		postID := bucket.(map[string]interface{})["key"].(string)
		postIDs = append(postIDs, postID)
	}

	return postIDs, nil
}

func (m *moviesElasticOps) GetTrendingPageSortOrder(page, size int) ([]string, error) {
	// Define query parameters
	query := map[string]interface{}{
		"from": (page - 1) * size,
		"size": size,
		"query": map[string]interface{}{
			"range": map[string]interface{}{
				"timestamp": map[string]interface{}{
					"gte": "now-15d",
				},
			},
		},
		"aggs": map[string]interface{}{
			"post": map[string]interface{}{
				"terms": map[string]interface{}{
					"field": "post_id",
					"order": map[string]interface{}{
						"avg_duration": "desc",
					},
				},
				"aggs": map[string]interface{}{
					"avg_duration": map[string]interface{}{
						"avg": map[string]interface{}{
							"field": "duration_of_scroll",
						},
					},
				},
			},
		},
	}

	// Encode query
	queryBytes, err := json.Marshal(query)
	if err != nil {
		return nil, fmt.Errorf("error encoding Elasticsearch query: %v", err)
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s/_search", config.ELASTIC_URI, config.MOVIE_INDEX), bytes.NewBuffer(queryBytes))
	if err != nil {
		return nil, fmt.Errorf("error creating Elasticsearch request: %v", err)
	}

	// Set request headers
	req.Header.Set("Content-Type", "application/json")

	// Send request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error executing Elasticsearch request: %v", err)
	}
	defer res.Body.Close()

	// Decode response
	var response map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("error decoding Elasticsearch response: %v", err)
	}

	// Extract post IDs
	var postIDs []string
	for _, bucket := range response["aggregations"].(map[string]interface{})["post"].(map[string]interface{})["buckets"].([]interface{}) {
		postID := bucket.(map[string]interface{})["key"].(string)
		postIDs = append(postIDs, postID)
	}

	return postIDs, nil
}
