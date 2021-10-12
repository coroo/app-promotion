package deliveries

import (
	"net/http"

	usecases "github.com/coroo/go-starter/app/usecases"
	entity "github.com/coroo/go-starter/app/entity"
	// "github.com/coroo/go-starter/app/middlewares"
	
	"github.com/gin-gonic/gin"
)

type PromotionForProductController interface {
	GetPromotionForProduct(ctx *gin.Context)
	// 	Save(ctx *gin.Context) error
}

type promotionForProductController struct {
	usecases usecases.PromotionForProductService
}

// var validate *validator.Validate

func NewPromotionForProductController(router *gin.Engine, apiPrefix string, promotionForProductService usecases.PromotionForProductService) {
	// validate = validator.New()
	// validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	handlerPromotionForProduct := &promotionForProductController{
		usecases: promotionForProductService,
	}
	promotionForProductsGroup := router.Group(apiPrefix + "promotionForProduct")
	{
		promotionForProductsGroup.GET("index", handlerPromotionForProduct.PromotionForProductsIndex)
		promotionForProductsGroup.GET("detail/:id", handlerPromotionForProduct.PromotionForProductsDetail)
		promotionForProductsGroup.POST("create", handlerPromotionForProduct.PromotionForProductCreate)
		promotionForProductsGroup.PUT("update", handlerPromotionForProduct.PromotionForProductUpdate)
		promotionForProductsGroup.DELETE("delete", handlerPromotionForProduct.PromotionForProductDelete)
	}
}

// GetPromotionForProductsIndex godoc
// @Security basicAuth
// @Summary Show all existing PromotionForProducts
// @Description Get all existing PromotionForProducts
// @Tags PromotionForProducts
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.PromotionForProduct
// @Failure 401 {object} dto.Response
// @Router /promotionForProduct/index [get]
func (deliveries *promotionForProductController) PromotionForProductsIndex(c *gin.Context) {
	promotionForProducts := deliveries.usecases.GetAllPromotionForProducts()
	c.JSON(http.StatusOK, gin.H{"data": promotionForProducts})
}

// GetPromotionForProductsDetail godoc
// @Security basicAuth
// @Summary Show an existing PromotionForProducts
// @Description Get detail the existing PromotionForProducts
// @Tags PromotionForProducts
// @Accept  json
// @Produce  json
// @Param  id path int true "Master Question ID"
// @Success 200 {array} entity.PromotionForProduct
// @Failure 401 {object} dto.Response
// @Router /promotionForProduct/detail/{id} [get]
func (deliveries *promotionForProductController) PromotionForProductsDetail(c *gin.Context) {
	promotionForProduct := deliveries.usecases.GetPromotionForProduct(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"data": promotionForProduct})
}

// CreatePromotionForProducts godoc
// @Security basicAuth
// @Summary Create new PromotionForProducts
// @Description Create a new PromotionForProducts
// @Tags PromotionForProducts
// @Accept  json
// @Produce  json
// @Param promotionForProduct body entity.PromotionForProduct true "Create promotionForProduct"
// @Success 200 {object} entity.PromotionForProduct
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /promotionForProduct/create [post]
func (deliveries *promotionForProductController) PromotionForProductCreate(c *gin.Context) {
	var promotionForProductEntity entity.PromotionForProduct
	c.ShouldBindJSON(&promotionForProductEntity)
	promotionForProductPK, err := deliveries.usecases.SavePromotionForProduct(promotionForProductEntity)
	if(err!=nil){
		c.JSON(http.StatusConflict, err)
	} else {
		promotionForProductEntity.ID = promotionForProductPK
		c.JSON(http.StatusOK, promotionForProductEntity)
	}
}

// UpdatePromotionForProducts godoc
// @Security basicAuth
// @Summary Update PromotionForProducts
// @Description Update a PromotionForProducts
// @Tags PromotionForProducts
// @Accept  json
// @Produce  json
// @Param promotionForProduct body entity.PromotionForProduct true "Update promotionForProduct"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /promotionForProduct/update [put]
func (deliveries *promotionForProductController) PromotionForProductUpdate(c *gin.Context) {
	var promotionForProductEntity entity.PromotionForProduct
	c.ShouldBindJSON(&promotionForProductEntity)
	promotionForProduct := deliveries.usecases.UpdatePromotionForProduct(promotionForProductEntity)
	c.JSON(http.StatusOK, promotionForProduct)
}

// DeletePromotionForProducts godoc
// @Security basicAuth
// @Summary Delete PromotionForProducts
// @Description Delete a PromotionForProducts
// @Tags PromotionForProducts
// @Accept  json
// @Produce  json
// @Param promotionForProduct body entity.PromotionForProduct true "Delete promotionForProduct"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /promotionForProduct/delete [delete]
func (deliveries *promotionForProductController) PromotionForProductDelete(c *gin.Context) {
	var promotionForProductEntity entity.PromotionForProduct
	c.ShouldBindJSON(&promotionForProductEntity)
	promotionForProduct := deliveries.usecases.DeletePromotionForProduct(promotionForProductEntity)
	c.JSON(http.StatusOK, promotionForProduct)
}