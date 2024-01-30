package app

import (
	"fmt"
	"github.com/gorilla-go/pig"
	"github.com/gorilla-go/pig/di"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"pig_framework/config"
)

type Provider struct {
}

func (p Provider) Handle(context *pig.Context, f func(*pig.Context)) {
	di.ProvideLazy[*gorm.DB](context.Container(), func(c *di.Container) (any, error) {
		db, err := gorm.Open(mysql.New(mysql.Config{
			DSN: fmt.Sprintf(
				"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
				config.Config[string]("MYSQL_USERNAME"),
				config.Config[string]("MYSQL_PASSWORD"),
				config.Config[string]("MYSQL_HOST"),
				config.Config[int]("MYSQL_PORT"),
				config.Config[string]("MYSQL_DATABASE"),
			),
		}), &gorm.Config{})
		if err != nil {
			return nil, err
		}

		return db, nil
	})

	f(context)
}
