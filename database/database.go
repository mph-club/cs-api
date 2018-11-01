package database

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/go-pg/pg"
)

func handleConnectError(connectError error) {
	connectRefuse := "5432: connect: connection refused"
	noSuch := "no such host"
	containsRefused := strings.Contains(connectError.Error(), connectRefuse)
	containsNoSuch := strings.Contains(connectError.Error(), noSuch)

	if containsRefused || containsNoSuch {
		log.Println("db not ready yet!")
		threeSeconds := time.Duration(3) * time.Second
		time.Sleep(threeSeconds)
	}
}

func connectToDB() *pg.DB {
	hostPortString := fmt.Sprintf("%s:%s", os.Getenv("PGHOST"), os.Getenv("PGPORT"))

	options := &pg.Options{
		User:     os.Getenv("PGUSER"),
		Password: os.Getenv("PGPASSWORD"),
		Database: os.Getenv("PGDATABASE"),
		Addr:     hostPortString,
	}

	db := pg.Connect(options)

	var n int
	_, err := db.QueryOne(pg.Scan(&n), "SELECT 1")
	if err != nil {
		handleConnectError(err)
		return connectToDB()
	}

	return db
}
