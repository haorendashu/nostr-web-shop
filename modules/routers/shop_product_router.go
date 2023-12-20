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

// add a product
func ShopProductAdd(c *gin.Context) {
	productDto := &dtos.ProductDto{}
	if err := c.ShouldBindJSON(productDto); err != nil {
		log.Printf("ShopProductAdd json parse error %v", err)
		c.JSON(http.StatusOK, Result(consts.API_CODE_ERROR, "Arg parse error"))
		return
	}

	if productDto.Skus == nil || len(productDto.Skus) == 0 {
		c.JSON(http.StatusOK, Result(consts.API_CODE_ERROR, "sku can't be null"))
		return
	}

	pubkey := c.GetString(SESSION_PUBKEY)
	product := &models.Product{}
	productDetails := make([]*models.ProductDetail, len(productDto.Skus))

	deepcopier.Copy(productDto).To(product)
	if product.Name == "" || product.Imgs == "" || product.Content == "" || product.Price < 1 {
		c.JSON(http.StatusOK, Result(consts.API_CODE_ERROR, "product args error"))
		return
	}
	minPrice := product.Price
	for i, sku := range productDto.Skus {
		productDetail := &models.ProductDetail{}
		deepcopier.Copy(sku).To(productDetail)

		if productDetail.Code == "" || productDetail.Stock <= 0 || productDetail.Price < 1 {
			c.JSON(http.StatusOK, Result(consts.API_CODE_ERROR, "product args error"))
			return
		}

		if productDetail.Name == "" {
			productDetail.Name = product.Name
		}
		if productDetail.Price < minPrice {
			minPrice = productDetail.Price
		}
		productDetails[i] = productDetail
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

	now := utils.NowInt64()
	product.UpdatedAt = now
	product.Status = consts.DATA_STATUS_OK
	product.Price = minPrice
	product.Pubkey = pubkey

	if product.Id == "" {
		// this is a new action
		product.Id = utils.RandomId()
		product.CreatedAt = now

		if result := models.ObjInsert(product, session); !result {
			c.JSON(http.StatusOK, Result(consts.API_CODE_ERROR, "obj save error"))
			return
		}
		for _, pd := range productDetails {
			pd.Id = utils.RandomId()
			pd.Pid = product.Id
			pd.Status = consts.DATA_STATUS_OK
			if result := models.ObjInsert(pd, session); !result {
				c.JSON(http.StatusOK, Result(consts.API_CODE_ERROR, "obj save error"))
				return
			}
		}
	} else {
		// update
		// try to find old product
		oldProduct := models.ProductGet(product.Id, session)
		if oldProduct == nil {
			c.JSON(http.StatusOK, Result(consts.API_CODE_ERROR, "product can't find"))
			return
		}

		// check pubkey
		if oldProduct.Pubkey != product.Pubkey {
			c.JSON(http.StatusOK, Result(consts.API_CODE_ERROR, "can't has permission"))
			return
		}
		// TODO check product if is forbiden

		product.CreatedAt = oldProduct.CreatedAt
		if !models.ObjUpdate(product.Id, product, session) {
			c.JSON(http.StatusOK, Result(consts.API_CODE_ERROR, "product update fail"))
			return
		}

		if err := models.ProductDetailDel(product.Id, session); err != nil {
			log.Printf("ShopProductAdd ProductDetailDel error %v", err)
			c.JSON(http.StatusOK, Result(consts.API_CODE_ERROR, "handle old data fail"))
			return
		}
		for _, pd := range productDetails {
			pd.Id = utils.RandomId()
			pd.Pid = product.Id
			pd.Status = consts.DATA_STATUS_OK
			if !models.ObjInsert(pd, session) {
				c.JSON(http.StatusOK, Result(consts.API_CODE_ERROR, "productDetail insert fail"))
				return
			}
		}
	}

	complete = true

	result := Result(consts.API_CODE_OK, "OK")
	result["pid"] = product.Id
	c.JSON(http.StatusOK, result)
}

// delete a product
func ShopProductDel(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, Result(consts.API_CODE_ERROR, "id can't be null"))
		return
	}

	product := models.ProductGet(id)
	if product == nil {
		c.JSON(http.StatusOK, Result(consts.API_CODE_OK, "OK"))
		return
	}
	pubkey := c.GetString(SESSION_PUBKEY)

	if product.Pubkey != pubkey {
		c.JSON(http.StatusOK, Result(consts.API_CODE_OK, "pubkey not equal"))
		return
	}

	models.ProductDel(id)
	models.ProductDetailDel(id)
	models.ProductPushInfoDel(id)

	c.JSON(http.StatusOK, Result(consts.API_CODE_OK, "OK"))
}

// get a product
func ShopProductGet(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, Result(consts.API_CODE_ERROR, "id can't be null"))
		return
	}

	pubkey := c.GetString(SESSION_PUBKEY)
	productDto := productDtoGet(id)

	if productDto.Pubkey != pubkey {
		c.JSON(http.StatusOK, Result(consts.API_CODE_FORBIDDEN, "login need"))
		return
	}

	result := Result(consts.API_CODE_OK, "OK")
	result["data"] = productDto

	c.JSON(http.StatusOK, result)
}

func productDtoGet(id string) *dtos.ProductDto {
	productDto := &dtos.ProductDto{}
	if id != "" {
		product := models.ProductGet(id)
		if product != nil {
			deepcopier.Copy(product).To(productDto)
		}

		productDetails := models.ProductDetailList(id)
		pdds := make([]*dtos.ProductDetailDto, len(productDetails))
		for index, pd := range productDetails {
			pdd := &dtos.ProductDetailDto{}
			deepcopier.Copy(pd).To(pdd)
			pdds[index] = pdd
		}
		productDto.Skus = pdds
	}

	return productDto
}

// product list
func ShopProductList(c *gin.Context) {
	pubkey := c.GetString(SESSION_PUBKEY)
	args := &models.ProductQueryArgs{}
	args.Pubkey = pubkey

	list := models.ProductList(args)
	detailListMap := make(map[string][]*dtos.ProductDetailDto)

	pids := make([]string, 0)
	for _, product := range list {
		pids = append(pids, product.Id)
	}

	if len(pids) > 0 {
		detailList := models.ProductDetailListByPids(pids)
		for _, detail := range detailList {
			subDetailList := detailListMap[detail.Pid]
			if subDetailList == nil {
				subDetailList = make([]*dtos.ProductDetailDto, 0)
			}

			pdd := &dtos.ProductDetailDto{}
			deepcopier.Copy(detail).To(pdd)

			subDetailList = append(subDetailList, pdd)
			detailListMap[detail.Pid] = subDetailList
		}
	}

	products := make([]*dtos.ProductDto, 0)
	for _, product := range list {
		dto := &dtos.ProductDto{}
		deepcopier.Copy(product).To(dto)
		products = append(products, dto)

		dto.Skus = detailListMap[dto.Id]
	}

	result := Result(consts.API_CODE_OK, "OK")
	result["list"] = products
	c.JSON(http.StatusOK, result)
}
