package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ulule/deepcopier"
	"log"
	"net/http"
	"nostr-web-shop/modules/consts"
	"nostr-web-shop/modules/dtos"
	"nostr-web-shop/modules/models"
	"nostr-web-shop/modules/utils"
	"strings"
	"time"
)

// Create a new order
func UserOrderAdd(c *gin.Context) {
	orderDto := &dtos.OrderAddDto{}
	if err := c.ShouldBindJSON(orderDto); err != nil {
		log.Printf("UserOrderAdd json parse error %v", err)
		c.JSON(http.StatusOK, Result(consts.API_CODE_ERROR, "Arg parse error"))
		return
	}

	if orderDto.Skus == nil || len(orderDto.Skus) == 0 {
		c.JSON(http.StatusOK, Result(consts.API_CODE_ERROR, "sku can't be null"))
		return
	}

	for _, sku := range orderDto.Skus {
		if sku.DetailId == "" || sku.Num < 1 {
			c.JSON(http.StatusOK, Result(consts.API_CODE_ERROR, "sku info can't be null"))
			return
		}
	}

	pubkey := c.GetString(SESSION_PUBKEY)

	complete := false
	session := models.Begin()
	defer func() {
		if complete {
			session.Commit()
		} else {
			session.Rollback()
		}
	}()

	total := 0
	lnwallet := ""
	seller := ""
	productMap := make(map[string]*models.Product)
	productDetailMap := make(map[string]*models.ProductDetail)
	orderProducts := make([]*models.OrderProduct, 0)
	for _, sku := range orderDto.Skus {
		productDetail := models.ProductDetailGet(sku.DetailId, session)
		if productDetail == nil {
			c.JSON(http.StatusOK, Result(consts.API_CODE_ERROR, fmt.Sprintf("product detail %s can't find", sku.DetailId)))
			return
		}

		// check stock
		if productDetail.Stock < sku.Num {
			c.JSON(http.StatusOK, Result(consts.API_CODE_ERROR, fmt.Sprintf("product stock %s not enough", sku.DetailId)))
			return
		}

		productDetail.Stock -= sku.Num
		productDetailMap[productDetail.Id] = productDetail

		orderProduct := &models.OrderProduct{}
		deepcopier.Copy(productDetail).To(orderProduct)
		orderProduct.DetailId = productDetail.Id
		orderProduct.Num = sku.Num
		orderProduct.Id = ""

		product := productMap[orderProduct.Pid]
		if product == nil {
			product = models.ProductGet(orderProduct.Pid)
			if product == nil {
				c.JSON(http.StatusOK, Result(consts.API_CODE_ERROR, fmt.Sprintf("product %s can't find", product.Id)))
				return
			}
			productMap[orderProduct.Pid] = product
		}

		if lnwallet == "" {
			lnwallet = product.Lnwallet
		} else if lnwallet != product.Lnwallet {
			c.JSON(http.StatusOK, Result(consts.API_CODE_ERROR, "Lnwallet not the same"))
			return
		}

		if seller == "" {
			seller = product.Pubkey
		} else if seller != product.Pubkey {
			c.JSON(http.StatusOK, Result(consts.API_CODE_ERROR, "Seller not the same"))
			return
		}

		imgs := strings.Split(product.Imgs, ",")
		if imgs != nil && len(imgs) > 0 {
			orderProduct.Img = imgs[0]
		}

		orderProducts = append(orderProducts, orderProduct)
		total += orderProduct.Price * orderProduct.Num
	}

	now := utils.NowInt64()

	order := &models.Order{}
	order.Id = utils.RandomId()
	order.Pubkey = pubkey
	order.UpdatedAt = now
	order.CreatedAt = now
	order.Status = consts.DATA_STATUS_OK
	order.OrderStatus = consts.ORDER_STATUS_ORDERD
	order.PayStatus = consts.PAY_STATUS_UNPAY
	//order.PaiedTime
	order.Price = total
	order.Comment = orderDto.Comment
	order.Lnwallet = lnwallet
	order.Seller = seller

	// begin to save data
	if result := models.ObjInsert(order); !result {
		c.JSON(http.StatusOK, Result(consts.API_CODE_ERROR, "order save error"))
		return
	}
	for _, orderProduct := range orderProducts {
		orderProduct.Id = utils.RandomId()
		orderProduct.OrderId = order.Id
		if result := models.ObjInsert(orderProduct); !result {
			c.JSON(http.StatusOK, Result(consts.API_CODE_ERROR, "order product save error"))
			return
		}
	}
	// update stock
	for _, productDetail := range productDetailMap {
		if err := models.ProductDetailUpdateStock(productDetail.Id, productDetail.Stock); err != nil {
			log.Printf("UserOrderAdd ProductDetailUpdateStock error %v", err)
			c.JSON(http.StatusOK, Result(consts.API_CODE_ERROR, "update stock fail"))
			return
		}
	}

	complete = true

	result := Result(consts.API_CODE_OK, "OK")
	result["oid"] = order.Id
	c.JSON(http.StatusOK, result)
}

// Get user order detail
func UserPayOrderGet(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, Result(consts.API_CODE_ERROR, "id can't be null"))
		return
	}

	order := models.OrderGet(id)
	if order.Status != consts.DATA_STATUS_OK {
		c.JSON(http.StatusOK, Result(consts.API_CODE_ERROR, "order status error"))
		return
	}

	// check order pubkey ????

	payOrder := models.PayOrderGet(order.Id)
	if payOrder == nil {
		// need to create a payOrder
		payInfo := utils.GetAlbyPayInfo(order.Lnwallet, order.Price)
		if payInfo == nil || payInfo.Pr == "" {
			c.JSON(http.StatusOK, Result(consts.API_CODE_ERROR, "get payInfo error"))
			return
		}

		now := utils.NowInt64()

		payOrder = &models.PayOrder{}
		payOrder.Id = utils.RandomId()
		payOrder.Oid = order.Id
		payOrder.Price = order.Price
		payOrder.Pr = payInfo.Pr
		payOrder.VerifyUrl = payInfo.Verify
		payOrder.CreatedAt = now
		payOrder.Status = consts.DATA_STATUS_OK
		payOrder.PayStatus = consts.PAY_STATUS_UNPAY

		invoice, err := utils.LightningInvoiceParse(payInfo.Pr)
		if err != nil {
			log.Printf("UserPayOrderGet getInvoiceTimeout error %v", err)
			c.JSON(http.StatusOK, Result(consts.API_CODE_ERROR, "invoice parse error"))
			return
		}
		beginTime := invoice.Timestamp
		payOrder.ExpireTime = beginTime.Add(invoice.Expiry()).UnixMilli()
		payOrder.CheckedTime = beginTime.Add(time.Minute).UnixMilli()

		if result := models.ObjInsert(payOrder); !result {
			c.JSON(http.StatusOK, Result(consts.API_CODE_ERROR, "payOrder insert fail"))
			return
		}
	}

	// payOrder convert to DTO
	dto := &dtos.PayOrderDto{}
	deepcopier.Copy(payOrder).To(dto)

	result := Result(consts.API_CODE_OK, "OK")
	result["data"] = dto
	c.JSON(http.StatusOK, result)
}

// Order Get
func UserOrderGet(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, Result(consts.API_CODE_ERROR, "id can't be null"))
		return
	}

	order := models.OrderGet(id)
	if order.Status != consts.DATA_STATUS_OK {
		c.JSON(http.StatusOK, Result(consts.API_CODE_ERROR, "order status error"))
		return
	}

	pubkey := c.GetString(SESSION_PUBKEY)
	if order.Pubkey != pubkey {
		c.JSON(http.StatusOK, Result(consts.API_CODE_ERROR, "Seller not the same"))
		return
	}

	orderProducts := models.OrderProductListByOids([]string{order.Id})
	ops := make([]*models.OrderProduct, 0)
	for _, orderProduct := range orderProducts {
		ops = append(ops, orderProduct)
	}

	dto := &dtos.OrderDto{}
	deepcopier.Copy(order).To(dto)
	dto.Skus = ops

	result := Result(consts.API_CODE_OK, "OK")
	result["data"] = dto
	c.JSON(http.StatusOK, result)
}

// list user orders
func UserOrderList(c *gin.Context) {
	pubkey := c.GetString(SESSION_PUBKEY)
	orders := models.OrderListByBuyer(pubkey)
	length := len(orders)

	oids := make([]string, length)
	for index, order := range orders {
		oids[index] = order.Id
	}

	orderProducts := models.OrderProductListByOids(oids)
	orderProductMap := make(map[string][]*models.OrderProduct)
	for _, orderProduct := range orderProducts {
		ops := orderProductMap[orderProduct.OrderId]
		if ops == nil {
			ops = make([]*models.OrderProduct, 0)
		}
		ops = append(ops, orderProduct)

		orderProductMap[orderProduct.OrderId] = ops
	}

	list := make([]*dtos.OrderDto, 0)
	for _, order := range orders {
		dto := &dtos.OrderDto{}
		deepcopier.Copy(order).To(dto)

		ops := orderProductMap[order.Id]
		dto.Skus = ops

		list = append(list, dto)
	}

	result := Result(consts.API_CODE_OK, "OK")
	result["list"] = list
	c.JSON(http.StatusOK, result)
}
