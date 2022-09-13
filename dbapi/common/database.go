package common

import (
	"context"
	"fmt"
	"github.com/janrockdev/eth-wallet/dbapi/models"
	"os"
	"time"

	badger "github.com/dgraph-io/badger/v3"
)

const (
	badgerDiscardRatio = 0.5
	badgerGCInterval   = 10 * time.Second
)

var (
	BadgerRegistryNamespace = []byte("registry")
)

type (
	DB interface {
		Get(namespace, key []byte) (value []byte, err error)
		Set(namespace, key []byte, value []byte) error
		All(namespace []byte) (values []string, err error)
		Has(namespace, key []byte) (bool, error)
		Delete(namespace, key []byte) error
		Update(namespace, key []byte, newval []byte) error
		Close() error
	}
	BadgerDB struct {
		db         *badger.DB
		ctx        context.Context
		cancelFunc context.CancelFunc
		logger     badger.Logger
	}
)

func NewBadgerDB(dataDir string) (DB, error) {
	if err := os.MkdirAll(dataDir, 0774); err != nil {
		return nil, err
	}
	opts := badger.DefaultOptions("")
	opts.SyncWrites = true
	opts.Dir, opts.ValueDir = dataDir, dataDir

	opts.NumVersionsToKeep = 1
	opts.CompactL0OnClose = true
	opts.NumLevelZeroTables = 1
	opts.NumLevelZeroTablesStall = 2
	opts.ValueLogFileSize = 1024 * 1024 * 10

	badgerDB, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}

	bdb := &BadgerDB{
		db: badgerDB,
	}
	bdb.ctx, bdb.cancelFunc = context.WithCancel(context.Background())

	go bdb.runGC()
	return bdb, nil
}

func ConnectBadgerDB(dataDir string) (DB, error) {
	opts := badger.DefaultOptions("")
	opts.SyncWrites = true
	opts.Dir, opts.ValueDir = dataDir, dataDir

	badgerBD, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}

	bdb := &BadgerDB{
		db: badgerBD,
	}
	bdb.ctx, bdb.cancelFunc = context.WithCancel(context.Background())

	go bdb.runGC()
	return bdb, nil
}

func (bdb *BadgerDB) Get(namespace []byte, key []byte) (value []byte, err error) {
	err = bdb.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(badgerNamespaceKey(namespace, key))
		if err != nil {
			return err
		}
		value, err = item.ValueCopy(value)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (bdb *BadgerDB) Set(namespace, key []byte, value []byte) error {
	err := bdb.db.Update(func(txn *badger.Txn) error {
		return txn.Set(badgerNamespaceKey(namespace, key), value)
	})

	if err != nil {
		bdb.logger.Debugf("failed to set key %s for namespace %s: %v", key, namespace, err)
		return err
	}

	return nil
}

func (bdb *BadgerDB) All(namespace []byte) (res []string, err error) {
	err = bdb.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			_ = item.Value(func(v []byte) error {
				res = append(res, string(v))
				return nil
			})
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (bdb *BadgerDB) Has(namespace, key []byte) (ok bool, err error) {
	_, err = bdb.Get(namespace, key)
	switch err {
	case badger.ErrKeyNotFound:
		ok, err = false, nil
	case nil:
		ok, err = true, nil
	}
	return ok, err
}

func (bdb BadgerDB) Delete(namespace []byte, key []byte) (err error) {
	err = bdb.db.Update(func(txn *badger.Txn) error {
		return txn.Delete(badgerNamespaceKey(namespace, key))
	})
	return err
}

func (bdb BadgerDB) Update(namespace []byte, key []byte, newval []byte) (err error) {
	_, err = bdb.Has(namespace, key)
	if err != nil {
		return err
	} else {
		Logr.Debug("Record Found!")
		err := bdb.Delete(namespace, key)
		if err != nil {
			return err
		} else {
			Logr.Debug("Record Deleted!")
		}
		err = bdb.Set(namespace, key, EncodeToBytes(models.CFG{Key: string(key), Value: string(newval)}))
		if err != nil {
			return err
		}
	}
	return err
}

func (bdb *BadgerDB) Close() error {
	bdb.cancelFunc()
	return bdb.db.Close()
}

func (bdb *BadgerDB) runGC() {
	ticker := time.NewTicker(badgerGCInterval)
	for {
		select {
		case <-ticker.C:
			err := bdb.db.RunValueLogGC(badgerDiscardRatio)
			if err != nil {
				if err == badger.ErrNoRewrite {
					Logr.Debugf("no BadgerDB GC occured: %v", err)
				} else {
					Logr.Errorf("failed to GC BadgerDB: %v", err)
				}
			}
		case <-bdb.ctx.Done():
			return
		}
	}
}

func badgerNamespaceKey(namespace, key []byte) []byte {
	prefix := []byte(fmt.Sprintf("%s/", namespace))
	return append(prefix, key...)
}
