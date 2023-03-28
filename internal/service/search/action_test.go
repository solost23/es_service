package search

import (
	"context"
	"es_service/configs"
	"es_service/internal/models"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/solost23/protopb/gen/go/protos/common"
	es_service "github.com/solost23/protopb/gen/go/protos/es"
	"log"
	"os"
	"testing"
)

func TestAction_Deal(T *testing.T) {
	mdb, _ := models.InitMysql(&configs.MySQLConf{
		DataSourceName:  "root:123@tcp(localhost:3306)/es_service?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai",
		MaxOpenConn:     20,
		MaxIdleConn:     10,
		MaxConnLifeTime: 100,
	})
	esClient, _ := elastic.NewClient(
		// default: http://127.0.0.1:9200
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetSniff(false),
		elastic.SetTraceLog(log.New(os.Stdout, "[es]: ", log.LstdFlags)),
	)

	type arg struct {
		ctx     context.Context
		request *es_service.SearchRequest
	}
	type want struct {
		err error
	}
	type test struct {
		arg  arg
		want want
	}
	tests := []test{
		{
			arg: arg{
				ctx: context.Background(),
				request: &es_service.SearchRequest{
					Header: &common.RequestHeader{
						OperatorUid: 100,
						TraceId:     101,
					},
					MustQuery: &es_service.Query{
						//TermQueries: []*es_service.TermQuery{
						//	{Field: "Username", Value: "alex1"},
						//},
						//RangeQueries: []*es_service.RangeQuery{
						//	{Field: "retweets", Gte: "0"},
						//	{Field: "retweets", Lte: "3"},
						//},
						MultiMatchQueries: []*es_service.MultiMatchQuery{
							{Field: []string{"Username", "Introduce"}, Value: "alex2"},
						},
					},
					//MustNotQuery: &es.Query{
					//	TermQuery: []*es.TermQuery{
					//		{Filed: "user", Value: "alex"},
					//	},
					//},
					//ShouldQuery: &es.Query{
					//TermQuery: []*es.TermQuery{
					//	{Filed: "user", Value: "alex1"},
					//	{Filed: "message", Value: "Take alex1"},
					//},
					//MatchQuery: []*es.MatchQuery{
					//	{Field: "user", Value: "alex1"},
					//	{Field: "message", Value: "Take alex1"},
					//},
					//RangeQuery: []*es.RangeQuery{
					//	{Field: "user", Gt: "w"},
					//	{Field: "user", Lt: "z"},
					//},
					//},

					//Indices: []string{"twitter", "twitter1"},
					//Sort: []*es.Sort{
					//	{Field: "user", Ascending: true},
					//	{Field: "created", Ascending: false},
					//},
					//Page:   1,
					//Size:   10,
					Pretty: true,
				},
			},
			want: want{},
		},
	}
	action := NewActionWithCtx(context.Background())
	action.SetMysql(mdb)
	action.SetES(esClient)
	for _, test := range tests {
		got, err := action.Deal(test.arg.ctx, test.arg.request)
		if err != test.want.err {
			T.Errorf(err.Error(), got)
		}
		fmt.Printf("got: %v \n", got)
	}
}

func BenchmarkAction_Deal(B *testing.B) {
	mdb, _ := models.InitMysql(&configs.MySQLConf{
		DataSourceName:  "root:123@tcp(localhost:3306)/es_service?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai",
		MaxOpenConn:     20,
		MaxIdleConn:     10,
		MaxConnLifeTime: 100,
	})
	esClient, _ := elastic.NewClient(
		// default: http://127.0.0.1:9200
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetSniff(false),
		//elastic.SetTraceLog(log.New(os.Stdout, "[es]: ", log.LstdFlags)),
	)

	type arg struct {
		ctx     context.Context
		request *es_service.SearchRequest
	}
	type want struct {
		err error
	}
	type test struct {
		arg  arg
		want want
	}
	tests := []test{
		{
			arg: arg{
				ctx: context.Background(),
				request: &es_service.SearchRequest{
					Header: &common.RequestHeader{
						OperatorUid: 100,
						TraceId:     101,
					},
					MustQuery: &es_service.Query{
						TermQueries: []*es_service.TermQuery{
							{Field: "user", Value: "alex1"},
						},
						RangeQueries: []*es_service.RangeQuery{
							{Field: "retweets", Gte: "0"},
							{Field: "retweets", Lte: "3"},
						},
					},
					//MustNotQuery: &es.Query{
					//	TermQuery: []*es.TermQuery{
					//		{Filed: "user", Value: "alex"},
					//	},
					//},
					//ShouldQuery: &es.Query{
					//TermQuery: []*es.TermQuery{
					//	{Filed: "user", Value: "alex1"},
					//	{Filed: "message", Value: "Take alex1"},
					//},
					//MatchQuery: []*es.MatchQuery{
					//	{Field: "user", Value: "alex1"},
					//	{Field: "message", Value: "Take alex1"},
					//},
					//RangeQuery: []*es.RangeQuery{
					//	{Field: "user", Gt: "w"},
					//	{Field: "user", Lt: "z"},
					//},
					//},

					//Indices: []string{"twitter", "twitter1"},
					//Sort: []*es.Sort{
					//	{Field: "user", Ascending: true},
					//	{Field: "created", Ascending: false},
					//},
					//Page:   1,
					//Size:   10,
					Pretty: true,
				},
			},
			want: want{},
		},
	}
	action := NewActionWithCtx(context.Background())
	action.SetMysql(mdb)
	action.SetES(esClient)

	for i := 0; i < B.N; i++ {
		_, _ = action.Deal(tests[0].arg.ctx, tests[0].arg.request)
	}
}
