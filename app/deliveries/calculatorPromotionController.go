package deliveries

import (
	"net/http"

	// "github.com/coroo/go-starter/models"
	usecases "github.com/coroo/go-starter/app/usecases"
	entity "github.com/coroo/go-starter/app/entity"
	// "github.com/coroo/go-starter/app/middlewares"
	
	"github.com/gin-gonic/gin"
)

type CalculatorController interface {
	GetCalculator(ctx *gin.Context)
	// 	Save(ctx *gin.Context) error
}

type calculatorController struct {
	usecases usecases.CalculatorPromotionService
}

// var validate *validator.Validate

func NewCalculatorPromotionController(router *gin.Engine, apiPrefix string, calculatorService usecases.CalculatorPromotionService) {
	// validate = validator.New()
	// validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	handlerCalculator := &calculatorController{
		usecases: calculatorService,
	}
	promotionsGroup := router.Group(apiPrefix + "promotion")
	{
		promotionsGroup.POST("calculator", handlerCalculator.PromotionCalculate)
		// promotionsGroup.GET("index", handlerCalculator.PromotionsIndex)
		// promotionsGroup.GET("detail/:id", handlerCalculator.PromotionsDetail)
		// promotionsGroup.POST("create", handlerCalculator.PromotionCreate)
		// promotionsGroup.PUT("update", handlerCalculator.PromotionUpdate)
		// promotionsGroup.DELETE("delete", handlerCalculator.PromotionDelete)
	}
}

// CalculatePromotions godoc
// @Security basicAuth
// @Summary Calculate new CalculatorPromotions
// @Description Calculate a new CalculatorPromotions
// @Tags CalculatorPromotions
// @Accept  json
// @Produce  json
// @Param promotion body []entity.CalculatorPromotion true "Calculate promotion"
// @Success 200 {object} []entity.CalculatorPromotion
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /promotion/calculator [post]
func (deliveries *calculatorController) PromotionCalculate(c *gin.Context) {
	var promotionEntity []entity.CalculatorPromotion
	c.ShouldBindJSON(&promotionEntity)
	promotionRes := deliveries.usecases.CalculatorPromotion(promotionEntity)
	// if(err!=nil){
	// 	c.JSON(http.StatusConflict, err)
	// } else {
	// 	promotionEntity.ID = promotionPK
		// c.JSON(http.StatusOK, promotionEntity)
	// }
	c.JSON(http.StatusOK, promotionRes)
}