package routers

import (
	"bytes"
	"github.com/gin-gonic/gin"
	lru "github.com/hashicorp/golang-lru/v2"
	"io"
	"log"
	"net/http"
	"nostr-web-shop/modules/consts"
	"nostr-web-shop/modules/models"
	"nostr-web-shop/modules/utils"
)

func Result(code int, msg string) map[string]any {
	return gin.H{"code": code, "msg": msg}
}

const (
	TOKEN_KEY      = "Token"
	SESSION_PUBKEY = "pubkey"
)

var session *lru.Cache[string, string]

func InitSession() {
	session, _ = lru.New[string, string](2048)
}

//func CorsMiddle() gin.HandlerFunc {
//
//	return func(c *gin.Context) {
//		origin := c.Request.Header.Get("Origin")
//
//		header := c.Writer.Header()
//		if origin != "" {
//			header.Set("Access-Control-Allow-Origin", origin)
//		} else {
//			header.Set("Access-Control-Allow-Origin", "*")
//		}
//		header.Set("Access-Control-Allow-Credentials", "true")
//		header.Set("Access-Control-Expose-Headers", "Origin, Content-Length, Content-Type, Authorization, token")
//		header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, HEAD, OPTIONS")
//		header.Set("Access-Control-Max-Age", "1728000")
//
//		if c.Request.Method == "OPTIONS" {
//			c.AbortWithStatus(http.StatusOK)
//			return
//		}
//
//		c.Next()
//	}
//}

func ApiMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		// check header and check sign
		// set seller, code to context

		seller := c.GetHeader(consts.API_SELLER)
		code := c.GetHeader(consts.API_CODE)
		t := c.GetHeader(consts.API_T)
		sign := c.GetHeader(consts.API_SIGN)

		body := ""
		if c.Request.Method == "POST" {
			bodyData, err := io.ReadAll(c.Request.Body)
			if err == nil {
				body = string(bodyData)
			}

			c.Request.Body.Close()
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyData))
		}

		productPushInfo := models.ProductPushInfoGetByCode(seller, code)
		if productPushInfo == nil {
			c.JSON(http.StatusOK, Result(consts.RESULT_CODE_FORBIDDEN, "product info error"))
			c.Abort()
			return
		}

		str := seller + code + t + c.Request.Method + c.Request.RequestURI + body
		log.Printf("api temp str gen: %s", str)
		localSign := utils.Md5(str + productPushInfo.PushKey)

		if localSign != sign {
			c.JSON(http.StatusOK, Result(consts.RESULT_CODE_FORBIDDEN, "sign check error"))
			c.Abort()
			return
		}

		c.Set(consts.API_SELLER, seller)
		c.Set(consts.API_CODE, code)
		c.Set(consts.PUSH_INFO, productPushInfo)

		c.Next()
	}
}

func SessionMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := GetToken(c)
		pubkey, existSession := session.Get(token)
		if !existSession || pubkey == "" {
			c.JSON(http.StatusOK, Result(consts.RESULT_CODE_FORBIDDEN, "Login need"))
			c.Abort()
			return
		}

		c.Set(SESSION_PUBKEY, pubkey)
		c.Next()
	}
}

func SessionShopMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := GetToken(c)
		pubkey, existSession := session.Get(token)
		if !existSession || pubkey == "" {
			c.JSON(http.StatusOK, Result(consts.RESULT_CODE_FORBIDDEN, "Login need"))
			c.Abort()
			return
		}

		exist := false
		for _, spk := range utils.CONFIG.ShopPubkeys {
			if spk == pubkey {
				exist = true
				break
			}
		}

		if exist {
			c.Set(SESSION_PUBKEY, pubkey)
			c.Next()
		} else {
			c.JSON(http.StatusOK, Result(consts.RESULT_CODE_FORBIDDEN, "Login need"))
			c.Abort()
			return
		}
	}
}

func GetToken(c *gin.Context) string {
	token := c.GetHeader(TOKEN_KEY)
	if token != "" {
		return token
	}
	token, _ = c.GetQuery(TOKEN_KEY)
	if token != "" {
		return token
	}
	token, _ = c.Cookie(TOKEN_KEY)
	if token != "" {
		return token
	}

	return ""
}
