package search

import (
	"context"
	"es_service/internal/es"
	"es_service/internal/service/base"
	"github.com/olivere/elastic/v7"
	es_service "github.com/solost23/protopb/gen/go/protos/es"
	"math"
)

type Action struct {
	base.Action
}

func NewActionWithCtx(ctx context.Context) *Action {
	a := &Action{}
	a.SetContext(ctx)
	return a
}

func (a *Action) Deal(ctx context.Context, request *es_service.SearchRequest) (reply *es_service.SearchResponse, err error) {
	page := int(request.GetPage())
	size := int(request.GetSize())
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}

	mustQueries := make([]elastic.Query, 0)
	mustNotQueries := make([]elastic.Query, 0)
	shouldQueries := make([]elastic.Query, 0)
	if request.GetMustQuery() != nil {
		mustQueries = append(mustQueries, a.BuildInquires(request.GetMustQuery())...)
	}
	if request.GetMustNotQuery() != nil {
		mustNotQueries = append(mustNotQueries, a.BuildInquires(request.GetMustNotQuery())...)
	}
	if request.GetShouldQuery() != nil {
		shouldQueries = append(shouldQueries, a.BuildInquires(request.GetShouldQuery())...)
	}

	boolQuery := elastic.NewBoolQuery()
	if len(mustQueries) > 0 {
		boolQuery.Must(mustQueries...)
	}
	if len(mustNotQueries) > 0 {
		boolQuery.MustNot(mustNotQueries...)
	}
	if len(shouldQueries) > 0 {
		boolQuery.Should(shouldQueries...)
	}

	sorts := make([]elastic.Sorter, 0)
	if len(request.GetSort()) > 0 {
		for _, sort := range request.GetSort() {
			sorts = append(sorts, elastic.NewFieldSort(sort.GetField()).Order(sort.GetAscending()))
		}
	}

	searchResult, err := es.Search(ctx, a.GetES(),
		request.GetIndices(),
		boolQuery,
		(page-1)*10,
		size,
		request.GetPretty(),
		sorts,
	)
	if err != nil {
		return nil, err
	}

	records := make([]string, 0, len(searchResult.Hits.Hits))
	for _, hit := range searchResult.Hits.Hits {
		record, _ := hit.Source.MarshalJSON()
		records = append(records, string(record))
	}

	reply = &es_service.SearchResponse{
		PageList: &es_service.PageList{
			Size:    int32(size),
			Pages:   int64(math.Ceil(float64(searchResult.Hits.TotalHits.Value) / float64(size))),
			Total:   searchResult.Hits.TotalHits.Value,
			Current: int32(page),
		},
		Records: records,
	}
	return reply, nil
}

func (a *Action) BuildInquires(seeks *es_service.Query) []elastic.Query {
	inquires := make([]elastic.Query, 0)

	if len(seeks.GetTermQueries()) > 0 {
		for _, seek := range seeks.GetTermQueries() {
			inquires = append(inquires, elastic.NewTermQuery(seek.GetField(), seek.GetValue()))
		}
	}
	if len(seeks.GetMatchQueries()) > 0 {
		for _, seek := range seeks.GetMatchQueries() {
			inquires = append(inquires, elastic.NewMatchQuery(seek.GetField(), seek.GetValue()))
		}
	}
	if len(seeks.GetMultiMatchQueries()) > 0 {
		for _, seek := range seeks.GetMultiMatchQueries() {
			inquires = append(inquires, elastic.NewMultiMatchQuery(seek.GetValue(), seek.GetField()...))
		}
	}
	if len(seeks.GetRangeQueries()) > 0 {
		for _, seek := range seeks.GetRangeQueries() {
			rangeQuery := elastic.NewRangeQuery(seek.GetField())
			if seek.GetGt() != "" {
				rangeQuery.Gt(seek.GetGt())
			}
			if seek.GetLte() != "" {
				rangeQuery.Lte(seek.GetLte())
			}
			if seek.GetGte() != "" {
				rangeQuery.Gte(seek.GetGte())
			}
			if seek.GetLte() != "" {
				rangeQuery.Lte(seek.GetLte())
			}
			inquires = append(inquires, rangeQuery)
		}
	}
	return inquires
}
