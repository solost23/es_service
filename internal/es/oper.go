package es

import (
	"context"
	"errors"
	"fmt"
	"github.com/olivere/elastic/v7"
)

// create index
func CreateIndex(ctx context.Context, client *elastic.Client, mapping string, indices ...string) error {
	exists, err := client.IndexExists(indices...).Do(ctx)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}

	// create index
	for _, index := range indices {
		ci, err := client.CreateIndex(index).Body(mapping).Do(ctx)
		if err != nil {
			return err
		}
		if !ci.Acknowledged {
			return errors.New(fmt.Sprintf("索引:%s创建失败", index))
		}
	}
	return nil
}

// delete index
func DeleteIndex(ctx context.Context, client *elastic.Client, indices ...string) error {
	result, err := client.DeleteIndex(indices...).Do(ctx)
	if err != nil {
		return err
	}
	if !result.Acknowledged {
		return errors.New(fmt.Sprintf("索引:%s删除失败", indices))
	}
	return nil
}

// create document
func CreateDocument(ctx context.Context, client *elastic.Client, index string, id string, body any) (*elastic.IndexResponse, error) {
	put, err := client.Index().
		Index(index).
		Id(id).
		BodyJson(body).
		Do(ctx)
	if err != nil {
		return nil, err
	}
	return put, nil
}

// delete document
func DeleteDocument(ctx context.Context, client *elastic.Client, index string, id string) (*elastic.DeleteResponse, error) {
	result, err := client.Delete().
		Index(index).
		Id(id).
		Do(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// search document
func Search(ctx context.Context, client *elastic.Client, index []string, query elastic.Query, from int, size int, pretty bool, sorts []elastic.Sorter) (*elastic.SearchResult, error) {
	searchResult, err := client.Search().
		Index(index...).
		Query(query).
		SortBy(sorts...).
		From(from).Size(size).
		Pretty(pretty).
		Do(ctx)
	if err != nil {
		return nil, err
	}
	return searchResult, nil
}
