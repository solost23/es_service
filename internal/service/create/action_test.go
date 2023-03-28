package create

import (
	"context"
	"encoding/json"
	"es_service/configs"
	"es_service/internal/models"
	"github.com/olivere/elastic/v7"
	es_service "github.com/solost23/protopb/gen/go/protos/es"
	"log"
	"os"
	"testing"
	"time"
)

const (
	defaultTimeFormat = "2006/01/02 15:04:05"
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
		request *es_service.CreateRequest
	}
	type want struct {
		err error
	}
	type test struct {
		arg  arg
		want want
	}

	type Tweet struct {
		Username  string
		Msg       string
		CreatedAt string
	}
	tweet1Json, _ := json.Marshal(Tweet{Username: "alex1", Msg: "hello alex1", CreatedAt: time.Now().Format(defaultTimeFormat)})
	tweet2Json, _ := json.Marshal(Tweet{Username: "alex2", Msg: "hello alex2", CreatedAt: time.Now().AddDate(0, 0, 1).Format(defaultTimeFormat)})
	tweet3Json, _ := json.Marshal(Tweet{Username: "alex3", Msg: "hello alex3", CreatedAt: time.Now().AddDate(0, 0, 2).Format(defaultTimeFormat)})
	tests := []test{
		{
			arg: arg{
				ctx: context.Background(),
				request: &es_service.CreateRequest{
					Index:      "tweet",
					DocumentId: "1",
					Document:   string(tweet1Json),
				},
			},
			want: want{},
		},
		{
			arg: arg{
				ctx: context.Background(),
				request: &es_service.CreateRequest{
					Index:      "tweet",
					DocumentId: "2",
					Document:   string(tweet2Json),
				},
			},
			want: want{},
		},
		{
			arg: arg{
				ctx: context.Background(),
				request: &es_service.CreateRequest{
					Index:      "tweet",
					DocumentId: "3",
					Document:   string(tweet3Json),
				},
			},
			want: want{},
		},
	}
	action := NewActionWithCtx(context.Background())
	action.SetMysql(mdb)
	action.SetES(esClient)
	for _, test := range tests {
		_, err := action.Deal(test.arg.ctx, test.arg.request)
		if err != test.want.err {
			T.Errorf(err.Error())
		}
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
		request *es_service.CreateRequest
	}
	type want struct {
		err error
	}
	type test struct {
		arg  arg
		want want
	}

	type Tweet struct {
		Username  string
		Msg       string
		CreatedAt string
	}
	tweet1Json, _ := json.Marshal(Tweet{Username: "alex1", Msg: "hello alex1", CreatedAt: time.Now().Format(defaultTimeFormat)})
	tests := []test{
		{
			arg: arg{
				ctx: context.Background(),
				request: &es_service.CreateRequest{
					Index:      "tweet",
					DocumentId: "1",
					Document:   string(tweet1Json),
				},
			},
		},
	}
	action := NewActionWithCtx(context.Background())
	action.SetMysql(mdb)
	action.SetES(esClient)

	for i := 0; i < B.N; i++ {
		_, _ = action.Deal(tests[0].arg.ctx, tests[0].arg.request)
	}
}
