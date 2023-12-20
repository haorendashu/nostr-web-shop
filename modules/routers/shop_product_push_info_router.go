package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ulule/deepcopier"
	"log"
	"net/http"
	"nostr-web-shop/modules/consts"
	"nostr-web-shop/modules/dtos"
	"nostr-web-shop/modules/models"
	"nostr-web-shop/modules/utils"
)

func ShopProductPushInfoGet(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, Result(consts.API_CODE_ERROR, "id can't be null"))
		return
	}

	pubkey := c.GetString(SESSION_PUBKEY)
	dto := &dtos.ProductPushInfoDto{}
	product := models.ProductGet(id)

	if product == nil {
		c.JSON(http.StatusOK, Result(consts.API_NOT_FOUND, "not found"))
		return
	}

	if product.Pubkey != pubkey {
		c.JSON(http.StatusOK, Result(consts.API_CODE_FORBIDDEN, "login need"))
		return
	}

	if product != nil {
		dto.Name = product.Name

		pushInfo := models.ProductPushInfoGet(product.Id)
		if pushInfo != nil {
			deepcopier.Copy(pushInfo).To(dto)
		}
	}

	result := Result(consts.API_CODE_OK, "OK")
	result["data"] = dto

	c.JSON(http.StatusOK, result)
}

func ShopProductPushInfoSave(c *gin.Context) {
	pushDto := &dtos.ProductPushInfoDto{}
	if err := c.ShouldBindJSON(pushDto); err != nil {
		log.Printf("ShopProductPushInfoSave json parse error %v", err)
		c.JSON(http.StatusOK, Result(consts.API_CODE_ERROR, "Arg parse error"))
		return
	}

	if pushDto.Pid == "" || (pushDto.Status > 0 && pushDto.NoticePubkey == "") {
		log.Printf("ShopProductPushInfoSave dto error %v", pushDto)
		c.JSON(http.StatusOK, Result(consts.API_CODE_ERROR, "Arg error"))
		return
	}

	complete := false
	session := models.Begin()
	defer func() {
		if complete {
			session.Commit()
		} else {
			session.Rollback()
		}
	}()

	product := models.ProductGet(pushDto.Pid)
	if product == nil {
		log.Printf("ShopProductPushInfoSave product can't find %d", pushDto.Pid)
		c.JSON(http.StatusOK, Result(consts.API_CODE_ERROR, "Id error"))
		return
	}

	pubkey := c.GetString(SESSION_PUBKEY)
	if product.Pubkey != pubkey {
		log.Printf("ShopProductPushInfoSave product pubkey not equest %v", pushDto)
		c.JSON(http.StatusOK, Result(consts.API_CODE_ERROR, "forbidden"))
		return
	}

	if err := models.ProductPushInfoDel(product.Id, session); err != nil {
		log.Printf("ShopProductPushInfoSave ProductPushInfoDel error %v", err)
		c.JSON(http.StatusOK, Result(consts.API_CODE_ERROR, "handle old data fail"))
		return
	}

	info := &models.ProductPushInfo{}
	deepcopier.Copy(pushDto).To(info)
	info.Id = utils.RandomId()
	if !models.ObjInsert(info, session) {
		c.JSON(http.StatusOK, Result(consts.API_CODE_ERROR, "save fail"))
		return
	}

	complete = true

	result := Result(consts.API_CODE_OK, "OK")
	c.JSON(http.StatusOK, result)
}
