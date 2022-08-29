package common

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"github.com/janrockdev/eth-wallet/dbapi/models"
	"io/ioutil"
	"log"
)

func EncodeToBytes(p interface{}) []byte {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(p)
	if err != nil {
		log.Fatal(err) //TODO REPLACE
	}
	return buf.Bytes()
}

func Compress(s []byte) []byte {
	zipbuf := bytes.Buffer{}
	zipped := gzip.NewWriter(&zipbuf)
	_, err := zipped.Write(s)
	if err != nil {
		return nil
	}
	err = zipped.Close()
	if err != nil {
		return nil
	}
	return zipbuf.Bytes()
}

func Decompress(s []byte) []byte {
	rdr, _ := gzip.NewReader(bytes.NewReader(s))
	data, err := ioutil.ReadAll(rdr)
	if err != nil {
		log.Fatal(err)
	}
	err = rdr.Close()
	if err != nil {
		return nil
	}
	return data
}

func DecodeToRule(s []byte) models.Rule {
	r := models.Rule{}
	dec := gob.NewDecoder(bytes.NewReader(s))
	err := dec.Decode(&r)
	if err != nil {
		log.Fatal(err)
	}
	return r
}

func DecodeToRules(s [][]byte) []models.Rule {
	var r models.Rule
	var res []models.Rule
	for _, val := range s {
		dec := gob.NewDecoder(bytes.NewReader(val))
		err := dec.Decode(&r)
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, r)
	}
	return res
}
