package bootstrap

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"schedule/common"
	"time"
)

var Provides = wire.NewSet(
	UseETCD,
	UseMongo,
	UseHttp,
)

// UseETCD 初始化etcd
func UseETCD(Config *common.Config) (EtcdConn *common.EtcdConnector, err error) {
	cfg := clientv3.Config{
		Endpoints:   Config.Endpoints,
		DialTimeout: time.Duration(Config.DialTimeOut) * time.Millisecond,
	}
	cli, err := clientv3.New(cfg)
	EtcdConn.Cfg = cfg
	EtcdConn.Cli = cli
	return
}

func UseMongo(Config *common.Config) (MongoConn *common.MongoConnector, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(Config.ConnectTimeOut)*time.Millisecond)
	defer cancel()
	MongoConn.Cli, err = mongo.Connect(ctx, options.Client().ApplyURI(Config.ApplyUri))
	MongoConn.DB = MongoConn.Cli.Database(Config.DBName)
	MongoConn.Collection = MongoConn.DB.Collection(Config.CollectionName)
	return
}

func UseHttp(Config *common.Config) (HttpConnector *common.HttpConnector, err error) {
	HttpConnector.HttpServer = http.Server{
		Addr:         fmt.Sprintf(":%d", Config.Port),
		Handler:      NewRouter(),
		ReadTimeout:  time.Duration(Config.ReadTimeOut) * time.Millisecond,
		WriteTimeout: time.Duration(Config.WriteTimeOut) * time.Millisecond,
	}
	return HttpConnector, nil
}

func NewRouter() *gin.Engine {

	r := gin.Default()
	return r
}
