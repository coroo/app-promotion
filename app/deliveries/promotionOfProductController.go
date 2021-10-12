package deliveries

import (
	"net/http"

	usecases "github.com/coroo/go-starter/app/usecases"
	entity "github.com/coroo/go-starter/app/entity"
	// "github.com/coroo/go-starter/app/middlewares"
	
	"github.com/gin-gonic/gin"
)

type PromotionOfProductController interface {
	GetPromotionOfProduct(ctx *gin.Context)
	// 	Save(ctx *gin.Context) error
}

type promotionOfProductController struct {
	usecases usecases.PromotionOfProductService
}

// var validate *validator.Validate

func NewPromotionOfProductController(router *gin.Engine, apiPrefix string, promotionOfProductService usecases.PromotionOfProductService) {
	// validate = validator.New()
	// validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	handlerPromotionOfProduct := &promotionOfProductController{
		usecases: promotionOfProductService,
	}
	promotionOfProductsGroup := router.Group(apiPrefix + "promotionOfProduct")
	{
		promotionOfProductsGroup.GET("index", handlerPromotionOfProduct.PromotionOfProductsIndex)
		promotionOfProductsGroup.GET("detail/:id", handlerPromotionOfProduct.PromotionOfProductsDetail)
		promotionOfProductsGroup.POST("create", handlerPromotionOfProduct.PromotionOfProductCreate)
		promotionOfProductsGroup.PUT("update", handlerPromotionOfProduct.PromotionOfProductUpdate)
		promotionOfProductsGroup.DELETE("delete", handlerPromotionOfProduct.PromotionOfProductDelete)
	}
}

// GetPromotionOfProductsIndex godoc
// @Security basicAuth
// @Summary Show all existing PromotionOfProducts
// @Description Get all existing PromotionOfProducts
// @Tags PromotionOfProducts
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.PromotionOfProduct
// @Failure 401 {object} dto.Response
// @Router /promotionOfProduct/index [get]
func (deliveries *promotionOfProductController) PromotionOfProductsIndex(c *gin.Context) {
	promotionOfProducts := deliveries.usecases.GetAllPromotionOfProducts()
	c.JSON(http.StatusOK, gin.H{"data": promotionOfProducts})
}

// GetPromotionOfProductsDetail godoc
// @Security basicAuth
// @Summary Show an existing PromotionOfProducts
// @Description Get detail the existing PromotionOfProducts
// @Tags PromotionOfProducts
// @Accept  json
// @Produce  json
// @Param  id path int true "Master Question ID"
// @Success 200 {array} entity.PromotionOfProduct
// @Failure 401 {object} dto.Response
// @Router /promotionOfProduct/detail/{id} [get]
func (deliveries *promotionOfProductController) PromotionOfProductsDetail(c *gin.Context) {
	promotionOfProduct := deliveries.usecases.GetPromotionOfProduct(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"data": promotionOfProduct})
}

// CreatePromotionOfProducts godoc
// @Security basicAuth
// @Summary Create new PromotionOfProducts
// @Description Create a new PromotionOfProducts
// @Tags PromotionOfProducts
// @Accept  json
// @Produce  json
// @Param promotionOfProduct body entity.PromotionOfProduct true "Create promotionOfProduct"
// @Success 200 {object} entity.PromotionOfProduct
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /promotionOfProduct/create [post]
func (deliveries *promotionOfProductController) PromotionOfProductCreate(c *gin.Context) {
	var promotionOfProductEntity entity.PromotionOfProduct
	c.ShouldBindJSON(&promotionOfProductEntity)
	promotionOfProductPK, err := deliveries.usecases.SavePromotionOfProduct(promotionOfProductEntity)
	if(err!=nil){
		c.JSON(http.StatusConflict, err)
	} else {
		promotionOfProductEntity.ID = promotionOfProductPK
		c.JSON(http.StatusOK, promotionOfProductEntity)
	}
}

// UpdatePromotionOfProducts godoc
// @Security basicAuth
// @Summary Update PromotionOfProducts
// @Description Update a PromotionOfProducts
// @Tags PromotionOfProducts
// @Accept  json
// @Produce  json
// @Param promotionOfProduct body entity.PromotionOfProduct true "Update promotionOfProduct"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /promotionOfProduct/update [put]
func (deliveries *promotionOfProductController) PromotionOfProductUpdate(c *gin.Context) {
	var promotionOfProductEntity entity.PromotionOfProduct
	c.ShouldBindJSON(&promotionOfProductEntity)
	promotionOfProduct := deliveries.usecases.UpdatePromotionOfProduct(promotionOfProductEntity)
	c.JSON(http.StatusOK, promotionOfProduct)
}

// DeletePromotionOfProducts godoc
// @Security basicAuth
// @Summary Delete PromotionOfProducts
// @Description Delete a PromotionOfProducts
// @Tags PromotionOfProducts
// @Accept  json
// @Produce  json
// @Param promotionOfProduct body entity.PromotionOfProduct true "Delete promotionOfProduct"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /promotionOfProduct/delete [delete]
func (deliveries *promotionOfProductController) PromotionOfProductDelete(c *gin.Context) {
	var promotionOfProductEntity entity.PromotionOfProduct
	c.ShouldBindJSON(&promotionOfProductEntity)
	promotionOfProduct := deliveries.usecases.DeletePromotionOfProduct(promotionOfProductEntity)
	c.JSON(http.StatusOK, promotionOfProduct)
}