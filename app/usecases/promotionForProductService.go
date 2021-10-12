package usecases

import (
	entity "github.com/coroo/go-starter/app/entity"
	repositories "github.com/coroo/go-starter/app/repositories"

	// "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type PromotionForProductService interface {
	SavePromotionForProduct(entity.PromotionForProduct) (int, error)
	UpdatePromotionForProduct(entity.PromotionForProduct) error
	DeletePromotionForProduct(entity.PromotionForProduct) error
	GetAllPromotionForProducts() []entity.PromotionForProduct
	GetPromotionForProduct(id string) []entity.PromotionForProduct
}

type promotionForProductService struct {
	repositories repositories.PromotionForProductRepository
}

func NewPromotionForProductService(repository repositories.PromotionForProductRepository) PromotionForProductService {
	return &promotionForProductService{
		repositories: repository,
	}
}

func (usecases *promotionForProductService) GetAllPromotionForProducts() []entity.PromotionForProduct {
	return usecases.repositories.GetAllPromotionForProducts()
}

func (usecases *promotionForProductService) GetPromotionForProduct(id string) []entity.PromotionForProduct {
	return usecases.repositories.GetPromotionForProduct(id)
}

func (usecases *promotionForProductService) SavePromotionForProduct(promotionForProduct entity.PromotionForProduct) (int, error) {
	return usecases.repositories.SavePromotionForProduct(promotionForProduct)
}

func (usecases *promotionForProductService) UpdatePromotionForProduct(promotionForProduct entity.PromotionForProduct) error {
	return usecases.repositories.UpdatePromotionForProduct(promotionForProduct)
}

func (usecases *promotionForProductService) DeletePromotionForProduct(promotionForProduct entity.PromotionForProduct) error {
	return usecases.repositories.DeletePromotionForProduct(promotionForProduct)
}