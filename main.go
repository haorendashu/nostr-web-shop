package main

import (
	"flag"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"nostr-web-shop/modules/models"
	"nostr-web-shop/modules/routers"
	"nostr-web-shop/modules/utils"
	"time"
)

var syncDB = false

func main() {
	flag.BoolVar(&syncDB, "syncDB", false, "sync object to databases.")
	flag.Parse()
	utils.ParseConf()
	models.Init()
	routers.InitSession()

	if syncDB {
		models.Sync()
		return
	}

	r := gin.Default()

	r.Use(gin.Logger(), gin.Recovery(), cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "token"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))

	apiGroup := r.Group("/api")
	baseGroup := r.Group("/base")
	userGroup := r.Group("/user")
	shopGroup := r.Group("/shop")

	apiGroup.Use(routers.ApiMiddle())
	userGroup.Use(routers.SessionMiddle())
	shopGroup.Use(routers.SessionShopMiddle())

	apiGroup.GET("/order/:orderProductId", routers.ApiOrderGet)
	apiGroup.GET("/order/list", routers.ApiOrderList)
	baseGroup.GET("/login", routers.BaseLogin)
	baseGroup.GET("/product/:id", routers.BaseProductGet)
	baseGroup.GET("/product/list", routers.BaseProductList)
	userGroup.POST("/order/add", routers.UserOrderAdd)
	userGroup.GET("/order/:id", routers.UserOrderGet)
	userGroup.GET("/orderPay/:id", routers.UserPayOrderGet)
	userGroup.GET("/orderPayCheck/:pid", routers.UserPayOrderCheck)
	userGroup.GET("/orderPushInfo/:id", routers.PushInfoGet)
	userGroup.GET("/order/list", routers.UserOrderList)
	shopGroup.POST("/product/", routers.ShopProductAdd)
	shopGroup.GET("/product/:id", routers.ShopProductGet)
	shopGroup.DELETE("/product/:id", routers.ShopProductDel)
	shopGroup.GET("/product/list", routers.ShopProductList)
	shopGroup.GET("/productPushInfo/:id", routers.ShopProductPushInfoGet)
	shopGroup.POST("/productPushInfo/", routers.ShopProductPushInfoSave)

	for _, f := range utils.CONFIG.RootFiles {
		r.StaticFile(f, f)
	}

	r.Run(utils.CONFIG.Server)
}
