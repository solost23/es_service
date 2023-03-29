package delete

import (
	"context"
	"es_service/configs"
	"es_service/internal/es"
	"es_service/internal/models"
	es_service "github.com/solost23/protopb/gen/go/protos/es"
	"testing"
)

type arg struct {
	ctx     context.Context
	request *es_service.DeleteRequest
}

type test struct {
	arg arg
}

func TestAction_Deal(T *testing.T) {
	// 初始化链接
	MySQLClient, _ := models.InitMysql(&configs.MySQLConf{
		DataSourceName:  "root:123@tcp(localhost:3306)/es_service?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai",
		MaxOpenConn:     20,
		MaxIdleConn:     10,
		MaxConnLifeTime: 100,
	})

	EsClient, _ := es.NewESClient(&configs.ESConf{
		Host: "127.0.0.1",
		Port: 9200,
	})

	tests := []test{
		{
			arg: arg{
				ctx: context.Background(),
				request: &es_service.DeleteRequest{
					Index:      "twitter",
					DocumentId: "1",
				},
			},
		},
		{
			arg: arg{
				ctx: context.Background(),
				request: &es_service.DeleteRequest{
					Index:      "twitter",
					DocumentId: "2",
				},
			},
		},
	}

	action := NewActionWithCtx(context.Background())
	action.SetMysql(MySQLClient)
	action.SetES(EsClient)
	for _, t := range tests {
		if _, err := action.Deal(t.arg.ctx, t.arg.request); err != nil {
			T.Error(err)
		}
	}
}
