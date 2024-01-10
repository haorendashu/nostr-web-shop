package routers

import (
	"context"
	"fmt"
	"github.com/nbd-wtf/go-nostr"
	"github.com/nbd-wtf/go-nostr/nip04"
	"log"
	"net/http"
	"nostr-web-shop/modules/consts"
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
				go doBackgroundPush(order, orderProduct, pushInfo)
			}
		}
	}
}

func doBackgroundPush(order *models.Order, orderProduct *models.OrderProduct, pushInfo *models.ProductPushInfo) {
	pushDto := genPushInfo(order, orderProduct, pushInfo)

	response := utils.HttpGet(pushDto.PushUrl)
	log.Printf("push %s code %d", pushDto.PushUrl, response.StatusCode)
	if response.StatusCode == http.StatusOK {
		orderProduct.PushCompleted = consts.PUSH_COMPLETED
		models.ObjUpdate(orderProduct.Id, orderProduct)

		doPushWithDM(order, orderProduct, pushInfo)
	}
}

func doPushWithDM(order *models.Order, orderProduct *models.OrderProduct, pushInfo *models.ProductPushInfo) {
	if utils.CONFIG.PrivateKey != "" && pushInfo.NoticePubkey != "" && len(utils.CONFIG.NoticeRelays) > 0 {
		ss, err := nip04.ComputeSharedSecret(pushInfo.NoticePubkey, utils.CONFIG.PrivateKey)
		if err != nil {
			log.Printf("doPushFromDM nip04.ComputeSharedSecret error %v", err)
			return
		}

		plainContent := fmt.Sprintf("Order notice\nProduct: %s \nProduct Code: %s \nNumber: %d\nBuyer: %s \nOrderProductId: %s",
			orderProduct.Name, orderProduct.Code, orderProduct.Num, order.Pubkey, orderProduct.Id)
		encryptContent, err := nip04.Encrypt(plainContent, ss)
		if err != nil {
			log.Printf("doPushFromDM nip04.Encrypt error %v", err)
			return
		}

		pubkey, err := nostr.GetPublicKey(utils.CONFIG.PrivateKey)
		if err != nil {
			log.Printf("doPushFromDM nostr.GetPublicKey error %v", err)
			return
		}

		event := nostr.Event{
			PubKey:    pubkey,
			CreatedAt: nostr.Now(),
			Kind:      nostr.KindEncryptedDirectMessage,
			Tags:      []nostr.Tag{[]string{"p", order.Pubkey}},
			Content:   encryptContent,
		}
		event.Sign(utils.CONFIG.PrivateKey)

		ctx := context.Background()
		for _, relayUrl := range utils.CONFIG.NoticeRelays {
			relay, err := nostr.RelayConnect(ctx, relayUrl)
			if err != nil {
				log.Printf("doPushFromDM nostr.RelayConnect error %v", err)
				continue
			}

			if err := relay.Publish(ctx, event); err != nil {
				log.Printf("doPushFromDM relay.Publish error %v", err)
				continue
			}

			log.Printf("send notice DM message to %s success", relayUrl)
		}
	}
}
