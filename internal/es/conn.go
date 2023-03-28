package es

import (
	"es_service/configs"
	"github.com/olivere/elastic/v7"
	"log"
	"os"
	"strconv"
)

func NewESClient(ESConfig *configs.ESConf) (*elastic.Client, error) {
	return elastic.NewClient(
		// default: http://127.0.0.1:9200
		elastic.SetURL("http://"+ESConfig.Host+":"+strconv.Itoa(ESConfig.Port)),
		elastic.SetSniff(false),
		elastic.SetTraceLog(log.New(os.Stdout, "[es]: ", log.LstdFlags)),
	)
}
