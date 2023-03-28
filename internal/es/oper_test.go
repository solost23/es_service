package es

import (
	"context"
	"es_service/configs"
	"fmt"
	"github.com/olivere/elastic/v7"
	"testing"
	"time"
)

const mapping = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
			"properties":{
				"user":{
					"type":"keyword"
				},
				"message":{
					"type":"text",
					"store": true,
					"fielddata": true
				},
				"image":{
					"type":"keyword"
				},
				"created":{
					"type":"date"
				},
				"tags":{
					"type":"keyword"
				},
				"location":{
					"type":"geo_point"
				},
				"suggest_field":{
					"type":"completion"
				}
			}
	}
}`

var (
	es *elastic.Client
)

func init() {
	es, _ = NewESClient(&configs.ESConf{Host: "127.0.0.1", Port: 9200})
}

func TestCreateIndex(t *testing.T) {
	if got := CreateIndex(context.Background(), es, mapping, "twitter10"); got != nil {
		t.Errorf(got.Error())
	}
}

func TestDeleteIndex(t *testing.T) {
	if got := DeleteIndex(context.Background(), es, "twitter10"); got != nil {
		t.Errorf(got.Error())
	}
}

type Tweet struct {
	User     string                `json:"user"`
	Message  string                `json:"message"`
	Retweets int                   `json:"retweets"`
	Image    string                `json:"image,omitempty"`
	Created  time.Time             `json:"created,omitempty"`
	Tags     []string              `json:"tags,omitempty"`
	Location string                `json:"location,omitempty"`
	Suggest  *elastic.SuggestField `json:"suggest_field,omitempty"`
}

func TestCreateDocument(t *testing.T) {
	tweet1 := Tweet{User: "alex1", Message: "Take alex1", Retweets: 0}
	if put, err := CreateDocument(context.Background(), es, "tweet3", "1", tweet1); err != nil {
		t.Errorf(err.Error(), put)
	}
}

func TestDeleteDocument(t *testing.T) {
	if result, err := DeleteDocument(context.Background(), es, "tweet3", "1"); err != nil {
		t.Errorf(err.Error(), result)
	}
}

func TestSearchDocument(t *testing.T) {
	query := elastic.NewTermQuery("user", "alex1")
	result, err := Search(
		context.Background(),
		es,
		query,
		"user",
		true,
		0,
		10,
		true,
		"tweet3",
		"twitter",
	)
	if err != nil {
		t.Errorf(err.Error())
	}
	for _, hit := range result.Hits.Hits {
		r, err := hit.Source.MarshalJSON()
		if err != nil {
			t.Errorf(err.Error())
		}
		fmt.Println(string(r))
	}
}
