package service

import (
	"context"
	"es_service/internal/service/create"
	"es_service/internal/service/delete"
	"es_service/internal/service/search"
	"github.com/Shopify/sarama"
	"github.com/gookit/slog"
	"github.com/olivere/elastic/v7"
	"github.com/solost23/protopb/gen/go/protos/es"
	"gorm.io/gorm"
)

type ESService struct {
	sl            *slog.SugaredLogger
	mdb           *gorm.DB
	es            *elastic.Client
	kafkaProducer sarama.SyncProducer
	es.UnimplementedSearchServer
}

func NewESService(sl *slog.SugaredLogger, mdb *gorm.DB, es *elastic.Client, kafkaProducer sarama.SyncProducer) *ESService {
	return &ESService{
		sl:            sl,
		mdb:           mdb,
		es:            es,
		kafkaProducer: kafkaProducer,
	}
}

// search
func (h *ESService) Search(ctx context.Context, request *es.SearchRequest) (reply *es.SearchResponse, err error) {
	action := search.NewActionWithCtx(ctx)
	action.SetHeader(request.Header)
	action.SetSl(h.sl)
	action.SetMysql(h.mdb)
	action.SetES(h.es)
	return action.Deal(ctx, request)
}

// create
func (h *ESService) Create(ctx context.Context, request *es.CreateRequest) (reply *es.CreateResponse, err error) {
	action := create.NewActionWithCtx(ctx)
	action.SetHeader(request.Header)
	action.SetSl(h.sl)
	action.SetMysql(h.mdb)
	action.SetES(h.es)
	return action.Deal(ctx, request)
}

// delete
func (h *ESService) Delete(ctx context.Context, request *es.DeleteRequest) (reply *es.DeleteResponse, err error) {
	action := delete.NewActionWithCtx(ctx)
	action.SetHeader(request.Header)
	action.SetSl(h.sl)
	action.SetMysql(h.mdb)
	action.SetES(h.es)
	return action.Deal(ctx, request)
}
