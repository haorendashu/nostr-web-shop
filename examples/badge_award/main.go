package main

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nbd-wtf/go-nostr"
	"log"
	"net/http"
	"nostr-web-shop/modules/consts"
	"nostr-web-shop/modules/routers"
	"nostr-web-shop/modules/utils"
	"time"
)

func main() {
	//flag.BoolVar(&syncDB, "syncDB", false, "sync object to databases.")
	//flag.Parse()
	ParseConf()
	loadOrders()

	r := gin.Default()

	r.Use(gin.Logger(), gin.Recovery(), cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/orderReceive", OrderReceive)

	for _, f := range utils.CONFIG.RootFiles {
		r.StaticFile(f, f)
	}

	r.Run(utils.CONFIG.Server)
}

func OrderReceive(c *gin.Context) {
	seller := c.Query("nws-seller")
	code := c.Query("nws-code")
	t := c.Query("nws-t")
	sign := c.Query("nws-sign")
	orderProductId := c.Query("orderProductId")
	buyer := c.Query("buyer")
	num := c.Query("num")
	comment := c.Query("comment")
	paidTime := c.Query("paidTime")

	tempStr := fmt.Sprintf("%s%s%s%s%s%s%s%s", seller, code, t, buyer, comment, num, orderProductId, paidTime)
	log.Printf("pushUrl tempStr %s", tempStr)
	localSign := utils.Md5(tempStr + CONFIG.PushKey)

	if sign != localSign {
		c.JSON(http.StatusForbidden, routers.Result(consts.RESULT_CODE_ERROR, "order sign check error"))
		return
	}

	pubkey, err := nostr.GetPublicKey(utils.CONFIG.PrivateKey)
	if err != nil {
		c.JSON(http.StatusForbidden, routers.Result(consts.RESULT_CODE_ERROR, "pubkey gen error"))
		return
	}

	// sign check complete, begin to check order
	if checkOrder(orderProductId) {
		result := routers.Result(consts.RESULT_CODE_OK, "OK")
		c.JSON(http.StatusOK, result)
		return
	}

	addOrder(orderProductId)

	// check complete, send Badge Award event
	awardEvent := nostr.Event{
		PubKey:    pubkey,
		CreatedAt: nostr.Now(),
		Kind:      8,
		Tags:      []nostr.Tag{[]string{"a", CONFIG.BadgeAId}, []string{"p", buyer}},
		Content:   "",
	}
	awardEvent.Sign(utils.CONFIG.PrivateKey)

	ctx := context.Background()
	for _, relayUrl := range CONFIG.NoticeRelays {
		relay, err := nostr.RelayConnect(ctx, relayUrl)
		if err != nil {
			log.Printf("doPushFromDM nostr.RelayConnect error %v", err)
			continue
		}

		if err := relay.Publish(ctx, awardEvent); err != nil {
			log.Printf("doPushFromDM relay.Publish error %v", err)
			continue
		}

		log.Printf("send notice DM message to %s success", relayUrl)
	}

	result := routers.Result(consts.RESULT_CODE_OK, "OK")
	c.JSON(http.StatusOK, result)
}
