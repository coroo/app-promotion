package usecases

import (
	entity "github.com/coroo/go-starter/app/entity"
	repositories "github.com/coroo/go-starter/app/repositories"

	// "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type PromotionOfProductService interface {
	SavePromotionOfProduct(entity.PromotionOfProduct) (int, error)
	UpdatePromotionOfProduct(entity.PromotionOfProduct) error
	DeletePromotionOfProduct(entity.PromotionOfProduct) error
	GetAllPromotionOfProducts() []entity.PromotionOfProduct
	GetPromotionOfProduct(id string) entity.PromotionOfProduct
}

type promotionOfProductService struct {
	repositories repositories.PromotionOfProductRepository
}

func NewPromotionOfProductService(repository repositories.PromotionOfProductRepository) PromotionOfProductService {
	return &promotionOfProductService{
		repositories: repository,
	}
}

func (usecases *promotionOfProductService) GetAllPromotionOfProducts() []entity.PromotionOfProduct {
	return usecases.repositories.GetAllPromotionOfProducts()
}

func (usecases *promotionOfProductService) GetPromotionOfProduct(productCode string) entity.PromotionOfProduct {
	return usecases.repositories.GetPromotionOfProduct(productCode)
}

func (usecases *promotionOfProductService) SavePromotionOfProduct(promotionOfProduct entity.PromotionOfProduct) (int, error) {
	return usecases.repositories.SavePromotionOfProduct(promotionOfProduct)
}

func (usecases *promotionOfProductService) UpdatePromotionOfProduct(promotionOfProduct entity.PromotionOfProduct) error {
	return usecases.repositories.UpdatePromotionOfProduct(promotionOfProduct)
}

func (usecases *promotionOfProductService) DeletePromotionOfProduct(promotionOfProduct entity.PromotionOfProduct) error {
	return usecases.repositories.DeletePromotionOfProduct(promotionOfProduct)
}