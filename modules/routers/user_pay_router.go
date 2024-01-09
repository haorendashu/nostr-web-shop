package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ulule/deepcopier"
	"log"
	"net/http"
	"net/url"
	"nostr-web-shop/modules/consts"
	"nostr-web-shop/modules/dtos"
	"nostr-web-shop/modules/models"
	"nostr-web-shop/modules/utils"
	"strings"
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

func UserPayOrderCheck(c *gin.Context) {
	pid := c.Param("pid")
	if pid == "" {
		c.JSON(http.StatusOK, Result(consts.RESULT_CODE_ERROR, "pid can't be null"))
		return
	}

	checkResult := CheckLightningResult(pid)

	result := Result(consts.RESULT_CODE_OK, "OK")
	result["data"] = checkResult
	c.JSON(http.StatusOK, result)
}

func PushInfoGet(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, Result(consts.RESULT_CODE_ERROR, "id can't be null"))
		return
	}
	pubkey := c.GetString(SESSION_PUBKEY)

	order, list := GetPushInfo(id)
	if order == nil || len(list) == 0 {
		c.JSON(http.StatusOK, Result(consts.RESULT_CODE_ERROR, "gen push info error"))
		return
	}

	if order.Pubkey != pubkey {
		c.JSON(http.StatusOK, Result(consts.RESULT_CODE_ERROR, "haven't permission"))
		return
	}

	result := Result(consts.RESULT_CODE_OK, "OK")
	result["list"] = list
	c.JSON(http.StatusOK, result)
}

func GetPushInfo(oid string) (*models.Order, []*dtos.OrderPushInfoDto) {
	order := models.OrderGet(oid)
	if order == nil {
		log.Println("Order not found!")
		return nil, nil
	}

	if order.PayStatus != consts.PAY_STATUS_PAIED {
		log.Println("Order not paid!")
		return nil, nil
	}

	pids := make([]string, 0)
	orderProducts := models.OrderProductListByOids([]string{oid})
	for _, orderProduct := range orderProducts {
		pids = append(pids, orderProduct.Pid)
	}

	pushInfos := models.ProductPushInfoListByPids(pids)
	pushInfoMap := make(map[string]*models.ProductPushInfo)
	for _, pushInfo := range pushInfos {
		pushInfoMap[pushInfo.Pid] = pushInfo
	}

	list := make([]*dtos.OrderPushInfoDto, 0)
	for _, orderProduct := range orderProducts {
		pushInfo := pushInfoMap[orderProduct.Pid]

		if pushInfo == nil {
			log.Printf("product %s can't find product push info", orderProduct.Pid)
			continue
		}

		//address := pushInfo.PushAddress
		//
		//seller := orderProduct.Seller
		//code := orderProduct.Code
		//t := now
		//
		//orderProductId := orderProduct.Id
		//buyer := order.Pubkey
		//num := orderProduct.Num
		//comment := url.QueryEscape(order.Comment)
		//paidTime := order.PaidTime
		//
		//tempStr := fmt.Sprintf("%s%s%d%s%s%d%s%d", seller, code, t, buyer, comment, num, orderProductId, paidTime)
		//log.Printf("pushUrl tempStr %s", tempStr)
		//
		//sign := utils.Md5(tempStr + pushInfo.PushKey)
		//
		//pushUrl := address
		//if strings.Index(pushUrl, "?") == -1 {
		//	pushUrl += "?"
		//}
		//
		//pushUrl += "nws-seller=" + seller
		//pushUrl += "&nws-code=" + code
		//pushUrl += "&nws-t=" + fmt.Sprintf("%d", t)
		//pushUrl += "&nws-sign=" + sign
		//pushUrl += "&orderProductId=" + orderProductId
		//pushUrl += "&buyer=" + buyer
		//pushUrl += "&num=" + fmt.Sprintf("%d", num)
		//pushUrl += "&comment=" + comment
		//pushUrl += "&paidTime=" + fmt.Sprintf("%d", paidTime)
		//
		//dto := &dtos.OrderPushInfoDto{
		//	OrderProductId: orderProductId,
		//	PushType:       pushInfo.PushType,
		//	PushUrl:        pushUrl,
		//}

		dto := genPushInfo(order, orderProduct, pushInfo)

		if orderProduct.PushCompleted != consts.PUSH_COMPLETED && dto.PushType == consts.PUSH_TYPE_WEB {
			orderProduct.PushCompleted = consts.PUSH_COMPLETED
			models.ObjUpdate(orderProduct.Id, orderProduct)
		}

		list = append(list, dto)
	}

	return order, list
}

func genPushInfo(order *models.Order, orderProduct *models.OrderProduct, pushInfo *models.ProductPushInfo) *dtos.OrderPushInfoDto {
	address := pushInfo.PushAddress

	seller := orderProduct.Seller
	code := orderProduct.Code
	t := utils.NowInt64()

	orderProductId := orderProduct.Id
	buyer := order.Pubkey
	num := orderProduct.Num
	comment := url.QueryEscape(order.Comment)
	paidTime := order.PaidTime

	tempStr := fmt.Sprintf("%s%s%d%s%s%d%s%d", seller, code, t, buyer, comment, num, orderProductId, paidTime)
	log.Printf("pushUrl tempStr %s", tempStr)

	sign := utils.Md5(tempStr + pushInfo.PushKey)

	pushUrl := address
	if strings.Index(pushUrl, "?") == -1 {
		pushUrl += "?"
	}

	pushUrl += "nws-seller=" + seller
	pushUrl += "&nws-code=" + code
	pushUrl += "&nws-t=" + fmt.Sprintf("%d", t)
	pushUrl += "&nws-sign=" + sign
	pushUrl += "&orderProductId=" + orderProductId
	pushUrl += "&buyer=" + buyer
	pushUrl += "&num=" + fmt.Sprintf("%d", num)
	pushUrl += "&comment=" + comment
	pushUrl += "&paidTime=" + fmt.Sprintf("%d", paidTime)

	dto := &dtos.OrderPushInfoDto{
		OrderProductId: orderProductId,
		PushType:       pushInfo.PushType,
		PushUrl:        pushUrl,
	}

	return dto
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
		po.PaidTime = now
		if !models.ObjUpdate(po.Id, po, session) {
			log.Println("checkLightningResult payOrder update fail")
			return false
		}

		o.UpdatedAt = now
		o.PayStatus = consts.PAY_STATUS_PAIED
		o.OrderStatus = consts.ORDER_STATUS_PAIED
		o.PaidTime = now
		if !models.ObjUpdate(o.Id, o, session) {
			log.Println("checkLightningResult order update fail")
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
