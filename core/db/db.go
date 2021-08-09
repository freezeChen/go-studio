package db

import (
	"time"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

func InitDb(c Config) *xorm.Engine {
	c = getConfig(c)
	engine, err := xorm.NewEngine(c.DriverName, c.Source)
	if err != nil {
		panic(engine)
	}
	engine.ShowSQL(c.Show)
	engine.SetTZLocation(time.Local)
	engine.SetMaxOpenConns(c.Max)
	engine.SetMaxIdleConns(c.Idle)
	engine.SetConnMaxLifetime(4 * time.Hour)
	engine.SetMapper(names.SnakeMapper{})

	if err := engine.Ping(); err != nil {
		panic(err)
	}

	return engine
}

func getConfig(c Config) Config {
	if c.Max == 0 {
		c.Max = 10
	}
	if c.Idle == 0 {
		c.Idle = 5
	}

	return c
}
