package routers

import (
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nbd-wtf/go-nostr"
	"log"
	"net/http"
	"nostr-web-shop/modules/consts"
	"nostr-web-shop/modules/utils"
	"strings"
	"time"
)

func BaseLogin(c *gin.Context) {
	authorizationStr := c.GetHeader("Authorization")
	authorizationStr = strings.Replace(authorizationStr, "Nostr ", "", 1)
	authEventData, err := base64.RawStdEncoding.DecodeString(authorizationStr)
	if err != nil {
		log.Printf("base64 decode error %v", err)
		c.JSON(http.StatusOK, Result(consts.RESULT_CODE_ERROR, "base64 decode error"))
		return
	}

	event := &nostr.Event{}
	err = json.Unmarshal(authEventData, event)
	if err != nil {
		log.Printf("event json.Unmarshal error %v", err)
		c.JSON(http.StatusOK, Result(consts.RESULT_CODE_ERROR, "event json.Unmarshal error"))
		return
	}

	if time.Now().UnixMilli()-event.CreatedAt.Time().UnixMilli() > 1000*60*5 {
		// CreatedAt is too long
		log.Println("event sign timeout")
		c.JSON(http.StatusOK, Result(consts.RESULT_CODE_ERROR, "event sign timeout"))
		return
	}

	loginEventU := ""
	for _, tag := range event.Tags {
		if tag.Key() == "u" {
			loginEventU = tag.Value()
		}
	}
	u := utils.CONFIG.LoginEventU
	if u != loginEventU {
		log.Println("login fail, u not match")
		c.JSON(http.StatusOK, Result(consts.RESULT_CODE_ERROR, "login fail u must be "+u))
		return
	}
	if result, err := event.CheckSignature(); !result || err != nil {
		log.Println("login fail sign check fail")
		c.JSON(http.StatusOK, Result(consts.RESULT_CODE_ERROR, "login fail sign check fail"))
		return
	}

	// set token to cache
	token := uuid.New().String()
	session.Add(token, event.PubKey)

	// return token
	result := Result(consts.RESULT_CODE_OK, "OK")
	result["token"] = token
	c.JSON(http.StatusOK, result)

	log.Println(result)
}
