package storage

import (
	"github.com/dgraph-io/badger"
	log "github.com/sirupsen/logrus"
	"os"
)

var db *badger.DB

func GetStorage() *badger.DB {
	// return the db if already opened
	if db != nil {
		return db
	}

	// open the db
	opts := badger.DefaultOptions
	opts.Dir = os.Getenv("STORAGE_PATH")
	opts.ValueDir = os.Getenv("STORAGE_PATH")
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatalf("can't open the storage: %v", err)
	}
	return db
}
