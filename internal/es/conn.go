package es

import (
	"es_service/configs"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"log"
	"os"
	"time"
)

func NewESClient(config *configs.ServerConfig) (*elastic.Client, error) {
	esConfig := config.ESConfig

	options := make([]elastic.ClientOptionFunc, 0)
	options = append(options, elastic.SetURL(fmt.Sprintf("http://%s:%d", esConfig.Host, esConfig.Port)))
	// 是否允许定期检查集群
	options = append(options, elastic.SetSniff(false))
	// 启动Gzip压缩, 会将请求数据压缩
	options = append(options, elastic.SetGzip(true))
	// 监控检查时间间隔
	options = append(options, elastic.SetHealthcheckInterval(10*time.Second))

	options = append(options, elastic.SetErrorLog(log.New(os.Stdout, "[ES]: ", log.LstdFlags)))
	if config.Mode == gin.DebugMode {
		options = append(options, elastic.SetTraceLog(log.New(os.Stdout, "[ES]: ", log.LstdFlags)))
	}

	return elastic.NewClient(options...)
}
