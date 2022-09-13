package main

import (
	"github.com/janrockdev/eth-wallet/dbapi/common"
)

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

	//del := db.Delete([]byte("registry"), []byte("totwallets"))
	//common.Logr.Info(del)

	ex, _ := db.Get([]byte("registry"), []byte("totwallets"))
	common.Logr.Info(string(ex))

	res, err := common.DB.All(db, []byte("registry"))
	for _, v := range res {
		common.Logr.Info(common.DecodeToStruct([]byte(v)).Key + " | " + common.DecodeToStruct([]byte(v)).Value)
	}

}
