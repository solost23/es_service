package handler

import (
	"es_service/internal/service"
	"github.com/solost23/protopb/gen/go/protos/es"
)

func Init(config Config) (err error) {
	// 1.gRPC::user service
	//oss.RegisterOssServer(config.Server, service.NewOSSService(config.Sl, config.MysqlConnect, config.RedisClient, config.KafkaProducer, config.MinioClient))
	es.RegisterSearchServer(config.Server, service.NewESService(config.Sl, config.MySQLClient, config.ESClient, config.KafkaProducer))
	return
}
