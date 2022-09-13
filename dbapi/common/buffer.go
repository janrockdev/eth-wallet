package common

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"github.com/janrockdev/eth-wallet/dbapi/models"
	"io/ioutil"
)

func EncodeToBytes(p interface{}) []byte {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(p)
	if err != nil {
		Logr.Fatal(err)
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
		Logr.Fatal(err)
	}
	err = rdr.Close()
	if err != nil {
		return nil
	}
	return data
}

func DecodeToStruct(s []byte) models.CFG {
	r := models.CFG{}
	dec := gob.NewDecoder(bytes.NewReader(s))
	err := dec.Decode(&r)
	if err != nil {
		Logr.Fatal(err)
	}
	return r
}

func DecodeToStructs(s [][]byte) []models.CFG {
	var r models.CFG
	var res []models.CFG
	for _, val := range s {
		dec := gob.NewDecoder(bytes.NewReader(val))
		err := dec.Decode(&r)
		if err != nil {
			Logr.Fatal(err)
		}
		res = append(res, r)
	}
	return res
}
