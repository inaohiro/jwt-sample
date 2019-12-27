package config

import (
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
)

type KeyEnv struct {
	Private string `required:"true"`
	Public  string
}

func KeyConfig() (*KeyEnv, error) {
	var e KeyEnv
	if err := envconfig.Process("key", &e); err != nil {
		log.Fatal(err.Error())
	}
	return &e, nil
}

// Database configuration
type DBEnv struct {
	User            string `required:"true"`
	Pass            string `required:"true"`
	Addr            string `default:"127.0.0.1:3306"`
	Name            string
	Driver          string `default:"mysql"`
	MaxOpenConns    int    `default:"10" split_words:"true"`
	MaxIdleConns    int    `default:"10" split_words:"true"`
	ConnMaxLifetime int    `default:"10" split_words:"true"`
}

func NewDBConfig() (*DBEnv, error) {
	var e DBEnv
	err := envconfig.Process("db", &e)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &e, nil
}

func (db *DBEnv) DataSource() string {
	return fmt.Sprintf("%s:%s@%s(%s)/%s", db.User, db.Pass, "tcp", db.Addr, db.Name)
}

// HTTP Server configuration
type ServerEnv struct {
	Addr string `default:"localhost"`
	Port int    `default:"8080"`
}

func ServerConfig() (*ServerEnv, error) {
	var e ServerEnv
	err := envconfig.Process("http", &e)
	if err != nil {
		log.Fatal(err.Error())
	}

	return &e, nil
}

// gRPC Server configuration
type GrpcEnv struct {
	Addr string `default:"localhost"`
	Port int    `default:"50051"`
}

func GrpcServerConfig() (*GrpcEnv, error) {
	var e GrpcEnv
	err := envconfig.Process("grpc", &e)
	if err != nil {
		log.Fatal(err.Error())
	}

	return &e, nil
}
