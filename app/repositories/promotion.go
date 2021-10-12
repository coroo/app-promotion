package repositories

import (
	"time"
	"gorm.io/gorm/clause"
	"github.com/coroo/go-starter/config"
	entity "github.com/coroo/go-starter/app/entity"

	// "github.com/gin-gonic/gin"
	"gorm.io/gorm"
	_ "gorm.io/driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

type PromotionRepository interface {
	SavePromotion(promotion entity.Promotion) (int, error)
	UpdatePromotion(promotion entity.Promotion) error
	DeletePromotion(promotion entity.Promotion) error
	GetAllPromotions() []entity.Promotion
	GetPromotion(id string) []entity.Promotion
}

type promotionDatabase struct {
	connection *gorm.DB
}

func NewPromotionRepository() PromotionRepository {
	db, err := config.ConnectDB()
	if err != nil {
		panic("Failed to connect database")
	}
	// db.AutoMigrate(&entity.Promotion{}, &entity.Person{})
	db.AutoMigrate(&entity.Promotion{})
	return &promotionDatabase{
		connection: db,
	}
}

func (db *promotionDatabase) SavePromotion(promotion entity.Promotion) (int, error) {
	data := &promotion
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	err := db.connection.Create(data)
	if(err.Error != nil){
		return 0, err.Error
	}
	return data.ID, nil
}

func (db *promotionDatabase) UpdatePromotion(promotion entity.Promotion) error {
	data := &promotion
	data.UpdatedAt = time.Now()
	db.connection.Save(data)
	return nil
}

func (db *promotionDatabase) DeletePromotion(promotion entity.Promotion) error {
	db.connection.Delete(&promotion)
	return nil
}

func (db *promotionDatabase) GetAllPromotions() []entity.Promotion {
	var promotions []entity.Promotion
	// db.connection.Preload(clause.Associations).Find(&promotions)
	db.connection.Set("gorm:auto_preload", true).Select("*").Joins("left join promotion_rule on promotion_rule.promotion_id = promotion.id").Find(&promotions)
	return promotions
}

func (db *promotionDatabase) GetPromotion(id string) []entity.Promotion {
	var promotion []entity.Promotion
	db.connection.Preload(clause.Associations).Where("id = ?", id).First(&promotion)
	return promotion
}
