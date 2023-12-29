package main

import (
	"playground/grpc/config"
	"playground/grpc/infrastructures"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	db := infrastructures.NewMongoDB(cfg.DBConfig.URI)

	_ = db
}
