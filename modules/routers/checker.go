package routers

import (
	"log"
	"net/http"
	"nostr-web-shop/modules/consts"
	"nostr-web-shop/modules/dtos"
	"nostr-web-shop/modules/models"
	"nostr-web-shop/modules/utils"
	"time"
)

func BeginCheck() {
	for true {
		doCheck()

		time.Sleep(time.Minute * 30)
	}
}

func doCheck() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("find err when docheck %v", r)
		}
	}()

	checkPayOrders()
	checkOrderPush()
}

func checkPayOrders() {
	payOrders := models.PayOrdersNeedChecked()
	for _, payOrder := range payOrders {
		CheckLightningResult(payOrder.Id)
	}
}

func checkOrderPush() {
	orderProducts := models.OrderProductsNeededPush()
	for _, orderProduct := range orderProducts {
		order := models.OrderGet(orderProduct.OrderId)
		pushInfo := models.ProductPushInfoGet(orderProduct.Pid)

		if order != nil && pushInfo != nil {
			if pushInfo.PushType == consts.PUSH_TYPE_API {
				pushDto := genPushInfo(order, orderProduct, pushInfo)

				go doBackgroundPush(orderProduct, pushDto)
			}
		}
	}
}

func doBackgroundPush(orderProduct *models.OrderProduct, pushDto *dtos.OrderPushInfoDto) {
	response := utils.HttpGet(pushDto.PushUrl)
	log.Printf("push %s code %d", pushDto.PushUrl, response.StatusCode)
	if response.StatusCode == http.StatusOK {
		orderProduct.PushCompleted = consts.PUSH_COMPLETED
		models.ObjUpdate(orderProduct.Id, orderProduct)
	}
}
