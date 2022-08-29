package controllers

import (
	"github.com/janrockdev/eth-wallet/common"
	"github.com/janrockdev/eth-wallet/models"
	"github.com/gin-gonic/gin"
)

common.Logr.Errorf("%v", err)
	}
	if c.Param("key") == "" {
		common.Logr.Debugf("res: error: missing key!")
}