package noop

import (
	"github.com/flipkart-incubator/dkv/internal/storage"
	"github.com/flipkart-incubator/dkv/pkg/serverpb"
)

type DB interface {
	storage.KVStore
}

type noopDB struct {
}

func OpenDB() (DB, error) {
	return &noopDB{}, nil
}

func (_ *noopDB) Put(key, val []byte) error {
	return nil
}

func (_ *noopDB) Get(keys ...[]byte) ([]*serverpb.KVPair, error) {
	return nil, nil
}

func (_ *noopDB) GetSnapshot() ([]byte, error) {
	return nil, nil
}

func (_ *noopDB) PutSnapshot(_ []byte) error {
	return nil
}

func (_ *noopDB) Iterate(_ storage.IterationOptions) storage.Iterator {
	return nil
}

func (_ *noopDB) Close() error {
	return nil
}
