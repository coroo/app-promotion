package usecases

import (
	entity "github.com/coroo/go-starter/app/entity"
	repositories "github.com/coroo/go-starter/app/repositories"

	// "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type PromotionService interface {
	SavePromotion(entity.Promotion) (int, error)
	UpdatePromotion(entity.Promotion) error
	DeletePromotion(entity.Promotion) error
	GetAllPromotions() []entity.Promotion
	GetPromotion(id string) []entity.Promotion
}

type promotionService struct {
	repositories repositories.PromotionRepository
}

func NewPromotionService(repository repositories.PromotionRepository) PromotionService {
	return &promotionService{
		repositories: repository,
	}
}

func (usecases *promotionService) GetAllPromotions() []entity.Promotion {
	return usecases.repositories.GetAllPromotions()
}

func (usecases *promotionService) GetPromotion(id string) []entity.Promotion {
	return usecases.repositories.GetPromotion(id)
}

func (usecases *promotionService) SavePromotion(promotion entity.Promotion) (int, error) {
	return usecases.repositories.SavePromotion(promotion)
}

func (usecases *promotionService) UpdatePromotion(promotion entity.Promotion) error {
	return usecases.repositories.UpdatePromotion(promotion)
}

func (usecases *promotionService) DeletePromotion(promotion entity.Promotion) error {
	return usecases.repositories.DeletePromotion(promotion)
}