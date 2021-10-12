package usecases

import (
	entity "github.com/coroo/go-starter/app/entity"
	repositories "github.com/coroo/go-starter/app/repositories"

	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type CalculatorPromotionService interface {
	CalculatorPromotion([]entity.CalculatorPromotion) []entity.CalculatorPromotion
}

type calculatorService struct {
	repositories repositories.PromotionOfProductRepository
}

func NewCalculatorPromotionService(repository repositories.PromotionOfProductRepository) CalculatorPromotionService {
	return &calculatorService{
		repositories: repository,
	}
}

func (usecases *calculatorService) CalculatorPromotion(promotion []entity.CalculatorPromotion) []entity.CalculatorPromotion {
	var i = 0
	var promoCapacity = 0
	var promotType = ""
	var cheapestProductCode = ""
	var cheapestPrice = 0
	promotionForProduct := make(map[int]string)

	// GET PROMOTION PRODUCTS
    for _, element := range promotion {
		checkPromotion := usecases.repositories.GetPromotionOfProduct(element.ProductCode)
		if checkPromotion.PromotionOfProductCode != "" {
			promotType = checkPromotion.Promotion.PromotionType
			for _, forElement := range checkPromotion.Promotion.PromotionForProduct {
				promotionForProduct[i] = forElement.PromotionForProductCode
				promoCapacity = forElement.PromotionCapacity
			}
			if promotType == "cheapest_one_free" {
				// SET CHEAPEST PRICE
				if element.Price < cheapestPrice || cheapestPrice == 0{
					cheapestProductCode = element.ProductCode
					cheapestPrice = element.Price
				}
			}
		}
		fmt.Println(cheapestProductCode)
		fmt.Println(cheapestPrice)
    }

	// SET CALCULATION FOR PROMO
	for key, element := range promotion {
		// HALF OF PRICE
		if promotType == "half_price" {
			for _, forElement := range promotionForProduct {
				if element.ProductCode == forElement && promoCapacity>0 {
					promotion[key].Price = element.Price/2
					// SET HALF OF PRICE FOR PROMOTION PRODUCT
					promoCapacity--
				}
			}
		}

		// CHEAPEST FREE
		if promotType == "cheapest_one_free" {
			if element.ProductCode == cheapestProductCode {
				promotion[key].Price = 0
				// SET HALF OF PRICE FOR PROMOTION PRODUCT
				promoCapacity--
			}
		}
	}

	return promotion
	// return usecases.repositories.SavePromotion(promotion)
}