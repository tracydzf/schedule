package common

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type Inject struct {
	Etcd       *EtcdConnector
	Mongo      *MongoConnector
	HttpServer *HttpConnector
}

type EtcdConnector struct {
	Cfg clientv3.Config
	Cli *clientv3.Client
}

// mongodb
type MongoConnector struct {
	Cli        *mongo.Client
	DB         *mongo.Database
	Collection *mongo.Collection
}

type HttpConnector struct {
	HttpServer http.Server
}
