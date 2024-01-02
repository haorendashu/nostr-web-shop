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
	"time"
)

// Get user order detail
func UserPayOrderGet(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, Result(consts.RESULT_CODE_ERROR, "id can't be null"))
		return
	}

	order := models.OrderGet(id)
	if order.Status != consts.DATA_STATUS_OK {
		c.JSON(http.StatusOK, Result(consts.RESULT_CODE_ERROR, "order status error"))
		return
	}

	// check order pubkey ????

	payOrder := models.PayOrderGetByOid(order.Id)
	if payOrder == nil {
		// need to create a payOrder
		payInfo := utils.GetAlbyPayInfo(order.Lnwallet, order.Price)
		if payInfo == nil || payInfo.Pr == "" {
			c.JSON(http.StatusOK, Result(consts.RESULT_CODE_ERROR, "get payInfo error"))
			return
		}

		now := utils.NowInt64()

		payOrder = &models.PayOrder{}
		payOrder.Id = utils.RandomId()
		payOrder.Oid = order.Id
		payOrder.Price = order.Price
		payOrder.Pr = payInfo.Pr
		payOrder.VerifyUrl = payInfo.Verify
		payOrder.UpdatedAt = now
		payOrder.CreatedAt = now
		payOrder.Status = consts.DATA_STATUS_OK
		payOrder.PayStatus = consts.PAY_STATUS_UNPAY

		invoice, err := utils.LightningInvoiceParse(payInfo.Pr)
		if err != nil {
			log.Printf("UserPayOrderGet getInvoiceTimeout error %v", err)
			c.JSON(http.StatusOK, Result(consts.RESULT_CODE_ERROR, "invoice parse error"))
			return
		}
		beginTime := invoice.Timestamp
		payOrder.ExpireTime = beginTime.Add(invoice.Expiry()).UnixMilli()
		payOrder.CheckedTime = beginTime.Add(time.Minute).UnixMilli()

		if result := models.ObjInsert(payOrder); !result {
			c.JSON(http.StatusOK, Result(consts.RESULT_CODE_ERROR, "payOrder insert fail"))
			return
		}
	}

	// payOrder convert to DTO
	dto := &dtos.PayOrderDto{}
	deepcopier.Copy(payOrder).To(dto)

	result := Result(consts.RESULT_CODE_OK, "OK")
	result["data"] = dto
	c.JSON(http.StatusOK, result)
}

// check and update payOrder, order
func CheckLightningResult(pid string) bool {
	po := models.PayOrderGet(pid)
	if po == nil {
		log.Printf("payOrder not found!")
		return false
	}

	if po.PayStatus == consts.PAY_STATUS_PAIED {
		return true
	}

	now := utils.NowInt64()

	result := utils.GetAlbyPrResult(po.VerifyUrl)
	if result == nil {
		log.Printf("utils.GetAlbyPrResult result is null!")
		return false
	}
	log.Printf("checkLightningResult payOrder %s pr check result %v", pid, result.Settled)
	if result.Status == "OK" && result.Settled {
		// pr complete

		complete := false
		session := models.Begin()
		defer func() {
			if complete {
				session.Commit()
			} else {
				session.Rollback()
			}
		}()

		o := models.OrderGet(po.Oid)
		if o == nil {
			log.Println("checkLightningResult order not found")
			return false
		}

		po.PayStatus = consts.PAY_STATUS_PAIED
		po.UpdatedAt = now
		if !models.ObjUpdate(po.Id, po, session) {
			log.Println("checkLightningResult payOrder update fail")
			return false
		}

		o.UpdatedAt = now
		o.PayStatus = consts.PAY_STATUS_PAIED
		o.OrderStatus = consts.ORDER_STATUS_PAIED
		o.PaidTime = now
		if !models.ObjUpdate(po.Id, po, session) {
			log.Println("checkLightningResult payOrder update fail")
			return false
		}

		complete = true
		return true
	}

	// the pr haven't paied, handle next check time
	po.CheckedTime = po.CreatedAt + (now-po.CreatedAt)*3
	po.UpdatedAt = now
	models.ObjUpdate(po.Id, po)

	return false
}
