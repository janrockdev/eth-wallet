package main

import (
	"bufio"
	"github.com/janrockdev/eth-wallet/dbapi/common"
	"github.com/janrockdev/eth-wallet/dbapi/models"
	"os"
	"strings"
)

var dat line

type line struct {
	Key   string
	Value string
}

func StoreToken(db common.DB, r line) {
	err := db.Set([]byte("registry"), []byte(r.Key), []byte(common.EncodeToBytes(models.CFG{Key: r.Key, Value: r.Value})))
	if err != nil {
		return
	}
}

func main() {
	file, err := os.Open("init/init.csv")
	if err != nil {
		common.Logr.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			common.Logr.Fatal(err)
		}
	}(file)

	db, err := common.NewBadgerDB("db")
	if err != nil {
		common.Logr.Fatal(err)
	}
	defer func(db common.DB) {
		err := db.Close()
		if err != nil {
			common.Logr.Fatal(err)
		}
	}(db)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		ss := strings.Split(scanner.Text(), ",")
		dat.Key = ss[0]
		dat.Value = ss[1]
		StoreToken(db, dat)
	}
}
