//go:build wireinject
// +build wireinject

package main

import (
	"github.com/cocoide/golang-design-pattern/dependecy_injection/entity"
	"github.com/cocoide/golang-design-pattern/dependecy_injection/handler"
	"github.com/cocoide/golang-design-pattern/dependecy_injection/repository/gorm"
	"github.com/cocoide/golang-design-pattern/dependecy_injection/usecase"
	"github.com/google/wire"
)

type handlers struct {
	BankHandler *handler.BankHandler
}

func initializeHandlers(config *entity.Config) (*handlers, error) {
	wire.Build(
		wire.NewSet(
			wire.FieldsOf(new(*entity.Config), "DB"),
			gorm.NewDBClient,
		),
		wire.NewSet(
			gorm.NewBankImpl,
			wire.Struct(new(gorm.GormParams), "*"),
		),
		usecase.NewBankUsecase,
		handler.NewBankHandler,
		wire.Struct(new(handlers), "*"),
	)
	return nil, nil
}
