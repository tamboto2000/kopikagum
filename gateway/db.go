package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db = dbConn()

func dbConn() *gorm.DB {
	dsn := "host=localhost user=postgres password=kepler22b dbname=kopi_kagum_gateway port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), new(gorm.Config))
	if err != nil {
		panic(err.Error())
	}

	return db
}
