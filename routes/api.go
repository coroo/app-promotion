package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/coroo/go-starter/app/deliveries"
	"github.com/coroo/go-starter/app/usecases"
	"github.com/coroo/go-starter/app/repositories"
	"github.com/coroo/go-starter/app/console"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Api() {
	router := gin.Default()
	// router.Use(middlewares.BasicAuth())

	API_PREFIX := os.Getenv("API_PREFIX")

	router.GET("/", func(c *gin.Context) {
		c.JSON(404, gin.H{
			"message": os.Getenv("MAIN_DESCRIPTION"),
		})
	})
	// PROMOTION
	var (
		promotionRepository repositories.PromotionRepository = repositories.NewPromotionRepository()
		promotionService    usecases.PromotionService = usecases.NewPromotionService(promotionRepository)
		// userController deliveries.UserController   = deliveries.NewUser(userService)
	)
	deliveries.NewPromotionController(router, API_PREFIX, promotionService)

	// PROMOTION FOR PRODUCT
	var (
		promotionForProductRepository repositories.PromotionForProductRepository = repositories.NewPromotionForProductRepository()
		promotionForProductService    usecases.PromotionForProductService = usecases.NewPromotionForProductService(promotionForProductRepository)
		// userController deliveries.UserController   = deliveries.NewUser(userService)
	)
	deliveries.NewPromotionForProductController(router, API_PREFIX, promotionForProductService)

	// PROMOTION Of PRODUCT
	var (
		promotionOfProductRepository repositories.PromotionOfProductRepository = repositories.NewPromotionOfProductRepository()
		promotionOfProductService    usecases.PromotionOfProductService = usecases.NewPromotionOfProductService(promotionOfProductRepository)
		// userController deliveries.UserController   = deliveries.NewUser(userService)
	)
	deliveries.NewPromotionOfProductController(router, API_PREFIX, promotionOfProductService)

	// CALCULATOR PROMOTION
	var (
		calculatorRepository repositories.PromotionOfProductRepository = repositories.NewPromotionOfProductRepository()
		calculatorService    usecases.CalculatorPromotionService = usecases.NewCalculatorPromotionService(calculatorRepository)
		// userController deliveries.UserController   = deliveries.NewUser(userService)
	)
	deliveries.NewCalculatorPromotionController(router, API_PREFIX, calculatorService)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	console.Schedule()
	router.Run(":"+os.Getenv("MAIN_PORT"))
}