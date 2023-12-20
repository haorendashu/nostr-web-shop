package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nostr-web-shop/modules/consts"
	"nostr-web-shop/modules/models"
	"strconv"
)

func BaseProductList(c *gin.Context) {
	args := &models.ProductQueryArgs{}
	// TODO handle args
	list := models.ProductList(args)
	result := Result(consts.API_CODE_OK, "OK")
	result["list"] = list
	c.JSON(http.StatusOK, result)
}

func BaseProductGet(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Result(consts.API_CODE_ERROR, "id parse error"))
		return
	}

	productDto := productDtoGet(id)

	result := Result(consts.API_CODE_OK, "OK")
	result["data"] = productDto

	c.JSON(http.StatusOK, result)
}
