package handler

import (
	"github.com/Shopify/sarama"
	"github.com/gookit/slog"
	"github.com/olivere/elastic/v7"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Config struct {
	Server        *grpc.Server
	Sl            *slog.SugaredLogger
	MySQLClient   *gorm.DB
	KafkaProducer sarama.SyncProducer
	ESClient      *elastic.Client
}
