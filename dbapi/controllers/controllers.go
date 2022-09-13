package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/janrockdev/eth-wallet/dbapi/common"
	"github.com/janrockdev/eth-wallet/dbapi/models"
	"net/http"
	"strconv"
)

func FindStat(c *gin.Context) {
	db, err := common.ConnectBadgerDB("db")
	if err != nil {
		common.Logr.Errorf("%v", err)
	}
	if c.Param("key") == "" {
		common.Logr.Debugf("res: error: missing key!")
		err := db.Close()
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": "error: missing key!"})
	} else {
		res, err := db.Get([]byte("registry"), []byte(c.Param("key")))
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"data": "error: key not found!"})
			common.Logr.Warnf("res: %v", err)
		} else {
			dec := common.DecodeToStruct(res)
			c.JSON(http.StatusOK, gin.H{"data": dec.Value})
			common.Logr.Debugf("res: %x", dec.Value)
		}
		err = db.Close()
		if err != nil {
			return
		}
	}
}

// ShowStats godoc
// @Summary get stats content
// @Description get stats content
// @Tags stats
// @Accept json
// @Produce json
// @Success 200 {array} models.Stats
// @Failure 400 {object} models.HTTPError
// @Failure 404 {object} models.HTTPError
// @Failure 500 {object} models.HTTPError
// @Router	/stats [get]
func ShowStats(c *gin.Context) {
	db, err := common.ConnectBadgerDB("db")
	if err != nil {
		common.Logr.Errorf("%v", err)
	}

	res, err := db.All([]byte("registry"))
	var final []models.CFG

	for _, v := range res {
		dec := common.DecodeToStruct([]byte(v))
		final = append(final, dec)
	}
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"data": "error: no key!"})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": final})
	}
	err = db.Close()
	if err != nil {
		return
	}
}

// CreateWallet WalletCreate godoc
// @Summary post wallet/create content
// @Description post wallet/create content
// @Tags wallet/create
// @Accept json
// @Produce json
// @Success 200 {array} models.Stats
// @Failure 400 {object} models.HTTPError
// @Failure 404 {object} models.HTTPError
// @Failure 500 {object} models.HTTPError
// @Router	/wallet/create [post]
func CreateWallet(c *gin.Context) {
	common.Logr.Info("wallet/create")
	var wallets int
	db, err := common.ConnectBadgerDB("db")
	if err != nil {
		common.Logr.Errorf("%v", err)
	}
	res, err := db.Get([]byte("registry"), []byte("totwallets"))
	if err != nil {
		common.Logr.Errorf("%v", err)
	}
	wallets, _ = strconv.Atoi(common.DecodeToStruct(res).Value)
	common.Logr.Info(wallets)
	wallets = wallets + 1
	common.Logr.Info(strconv.Itoa(wallets))
	err = db.Update([]byte("registry"), []byte("totwallets"), []byte(strconv.Itoa(wallets)))
	if err != nil {
		common.Logr.Errorf("%v", err)
	} else {
		common.Logr.Info("New wallet creted!")
	}
	err = db.Close()
	if err != nil {
		return
	}
}
