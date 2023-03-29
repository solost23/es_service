package es

import (
	"es_service/configs"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"log"
	"os"
)

func NewESClient(config *configs.ServerConfig) (*elastic.Client, error) {
	esConfig := config.ESConfig

	options := make([]elastic.ClientOptionFunc, 0)
	options = append(options, elastic.SetURL(fmt.Sprintf("http://%s:%d", esConfig.Host, esConfig.Port)))
	options = append(options, elastic.SetSniff(false))

	if config.Mode == gin.DebugMode {
		options = append(options, elastic.SetTraceLog(log.New(os.Stdout, "[ES]: ", log.LstdFlags)))
	}

	return elastic.NewClient(options...)
}
