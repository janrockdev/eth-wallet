package main

import (
	"github.com/janrockdev/eth-wallet/dbapi/common"
	"github.com/janrockdev/eth-wallet/dbapi/models"
)

type line struct {
	Key   string
	Value string
}

func UpdateToken(db common.DB, key []byte, newval line) {
	res, err := db.Has([]byte("registry"), key)
	if err != nil {
		return
	}
	if res {
		err := db.Delete([]byte("registry"), key)
		if err != nil {
			return
		}
		err = db.Set([]byte("registry"), key, common.EncodeToBytes(models.CFG{Key: newval.Key, Value: newval.Value}))
		if err != nil {
			return
		}
	}
}

func main() {

	db, err := common.NewBadgerDB("dbapi/db")
	if err != nil {
		common.Logr.Fatal(err)
	}
	defer func(db common.DB) {
		err := db.Close()
		if err != nil {
			common.Logr.Fatal(err)
		}
	}(db)

	UpdateToken(db, []byte("totwallets"), line(models.CFG{Key: "totwallets", Value: "7"}))
}
