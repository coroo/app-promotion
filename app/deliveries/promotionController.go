package deliveries

import (
	"net/http"

	// "github.com/coroo/go-starter/models"
	usecases "github.com/coroo/go-starter/app/usecases"
	entity "github.com/coroo/go-starter/app/entity"
	// "github.com/coroo/go-starter/app/middlewares"
	
	"github.com/gin-gonic/gin"
)

type PromotionController interface {
	GetPromotion(ctx *gin.Context)
	// 	Save(ctx *gin.Context) error
}

type promotionController struct {
	usecases usecases.PromotionService
}

// var validate *validator.Validate

func NewPromotionController(router *gin.Engine, apiPrefix string, promotionService usecases.PromotionService) {
	// validate = validator.New()
	// validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	handlerPromotion := &promotionController{
		usecases: promotionService,
	}
	promotionsGroup := router.Group(apiPrefix + "promotion")
	{
		promotionsGroup.GET("index", handlerPromotion.PromotionsIndex)
		promotionsGroup.GET("detail/:id", handlerPromotion.PromotionsDetail)
		promotionsGroup.POST("create", handlerPromotion.PromotionCreate)
		promotionsGroup.PUT("update", handlerPromotion.PromotionUpdate)
		promotionsGroup.DELETE("delete", handlerPromotion.PromotionDelete)
	}
}

// GetPromotionsIndex godoc
// @Security basicAuth
// @Summary Show all existing Promotions
// @Description Get all existing Promotions
// @Tags Promotions
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.Promotion
// @Failure 401 {object} dto.Response
// @Router /promotion/index [get]
func (deliveries *promotionController) PromotionsIndex(c *gin.Context) {
	promotions := deliveries.usecases.GetAllPromotions()
	c.JSON(http.StatusOK, gin.H{"data": promotions})
}

// GetPromotionsDetail godoc
// @Security basicAuth
// @Summary Show an existing Promotions
// @Description Get detail the existing Promotions
// @Tags Promotions
// @Accept  json
// @Produce  json
// @Param  id path int true "Master Question ID"
// @Success 200 {array} entity.Promotion
// @Failure 401 {object} dto.Response
// @Router /promotion/detail/{id} [get]
func (deliveries *promotionController) PromotionsDetail(c *gin.Context) {
	promotion := deliveries.usecases.GetPromotion(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"data": promotion})
}

// CreatePromotions godoc
// @Security basicAuth
// @Summary Create new Promotions
// @Description Create a new Promotions
// @Tags Promotions
// @Accept  json
// @Produce  json
// @Param promotion body entity.Promotion true "Create promotion"
// @Success 200 {object} entity.Promotion
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /promotion/create [post]
func (deliveries *promotionController) PromotionCreate(c *gin.Context) {
	var promotionEntity entity.Promotion
	c.ShouldBindJSON(&promotionEntity)
	promotionPK, err := deliveries.usecases.SavePromotion(promotionEntity)
	if(err!=nil){
		c.JSON(http.StatusConflict, err)
	} else {
		promotionEntity.ID = promotionPK
		c.JSON(http.StatusOK, promotionEntity)
	}
}

// UpdatePromotions godoc
// @Security basicAuth
// @Summary Update Promotions
// @Description Update a Promotions
// @Tags Promotions
// @Accept  json
// @Produce  json
// @Param promotion body entity.Promotion true "Update promotion"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /promotion/update [put]
func (deliveries *promotionController) PromotionUpdate(c *gin.Context) {
	var promotionEntity entity.Promotion
	c.ShouldBindJSON(&promotionEntity)
	promotion := deliveries.usecases.UpdatePromotion(promotionEntity)
	c.JSON(http.StatusOK, promotion)
}

// DeletePromotions godoc
// @Security basicAuth
// @Summary Delete Promotions
// @Description Delete a Promotions
// @Tags Promotions
// @Accept  json
// @Produce  json
// @Param promotion body entity.Promotion true "Delete promotion"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /promotion/delete [delete]
func (deliveries *promotionController) PromotionDelete(c *gin.Context) {
	var promotionEntity entity.Promotion
	c.ShouldBindJSON(&promotionEntity)
	promotion := deliveries.usecases.DeletePromotion(promotionEntity)
	c.JSON(http.StatusOK, promotion)
}