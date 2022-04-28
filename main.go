package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func Index() {
	addresses := []string{"http://127.0.0.1:9200"}
	config := elasticsearch.Config{
		Addresses: addresses,
	}
	// new client
	es, err := elasticsearch.NewClient(config)
	failOnError(err, "Error creating the client")
	// Index creates or updates a document in an index
	var buf bytes.Buffer
	doc := map[string]interface{}{
		"title":   "你看到外面的世界是什么样的？",
		"content": "外面的世界真的很精彩",
	}
	if err := json.NewEncoder(&buf).Encode(doc); err != nil {
		failOnError(err, "Error encoding doc")
	}
	res, err := es.Index("demo", &buf, es.Index.WithDocumentType("doc"))
	if err != nil {
		failOnError(err, "Error Index response")
	}
	defer res.Body.Close()
	fmt.Println(res.String())
}

func Search() {
	addresses := []string{"http://127.0.0.1:9200"}
	config := elasticsearch.Config{
		Addresses: addresses,
	}

	es, err := elasticsearch.NewClient(config)
	failOnError(err, "Error creating the client")

	res, err := es.Info()
	failOnError(err, "Error getting response")
	fmt.Println(res.String())

	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": "中国",
			},
		},
		"highlight": map[string]interface{}{
			"pre_tags":  []string{"<font color='red'>"},
			"post_tags": []string{"</font>"},
			"fields": map[string]interface{}{
				"title": map[string]interface{}{},
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		failOnError(err, "Error encoding query")
	}

	res, err = es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("demo"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		failOnError(err, "Error getting response")
	}
	defer res.Body.Close()
	fmt.Println(res.String())
}

func main() {
	Search()
}
