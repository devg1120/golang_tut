package main

import (
	"log"

	"github.com/cocoide/golang-design-pattern/dependecy_injection/entity"
	"github.com/kelseyhightower/envconfig"
)

func newConfig() *entity.Config {
	var c entity.Config
	if err := envconfig.Process("", &c); err != nil {
		log.Fatal(err)
	}
	return &c
}
