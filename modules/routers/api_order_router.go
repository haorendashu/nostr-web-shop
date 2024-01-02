package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nostr-web-shop/modules/consts"
	"nostr-web-shop/modules/models"
	"strconv"
)

func ApiOrderGet(c *gin.Context) {
	seller, _ := c.Get(consts.API_SELLER)
	code, _ := c.Get(consts.API_CODE)

	orderProductId := c.Param("orderProductId")
	if orderProductId == "" {
		c.JSON(http.StatusOK, Result(consts.RESULT_CODE_ERROR, "orderProductId can't be null"))
		return
	}

	orderApiDto := models.OrderApiGet(orderProductId)
	if orderApiDto.Seller != seller || orderApiDto.Code != code {
		c.JSON(http.StatusOK, Result(consts.RESULT_CODE_ERROR, "orderProductId and code error"))
		return
	}

	result := Result(consts.RESULT_CODE_OK, "OK")
	result["data"] = orderApiDto
	c.JSON(http.StatusOK, result)
}

func ApiOrderList(c *gin.Context) {
	sellerItf, _ := c.Get(consts.API_SELLER)
	codeItf, _ := c.Get(consts.API_CODE)
	sinceStr := c.Param(consts.API_SINCE)
	since, err := strconv.ParseInt(sinceStr, 64, 10)
	if err != nil {
		since = 0
	}

	l := models.OrderApiList(sellerItf.(string), codeItf.(string), since)
	result := Result(consts.RESULT_CODE_OK, "OK")
	result["list"] = l
	c.JSON(http.StatusOK, result)
}
