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

type PromotionForProductRepository interface {
	SavePromotionForProduct(promotionForProduct entity.PromotionForProduct) (int, error)
	UpdatePromotionForProduct(promotionForProduct entity.PromotionForProduct) error
	DeletePromotionForProduct(promotionForProduct entity.PromotionForProduct) error
	GetAllPromotionForProducts() []entity.PromotionForProduct
	GetPromotionForProduct(id string) []entity.PromotionForProduct
}

type promotionForProductDatabase struct {
	connection *gorm.DB
}

func NewPromotionForProductRepository() PromotionForProductRepository {
	db, err := config.ConnectDB()
	if err != nil {
		panic("Failed to connect database")
	}
	// db.AutoMigrate(&entity.PromotionForProductRate{}, &entity.Person{})
	db.AutoMigrate(&entity.PromotionForProduct{})
	return &promotionForProductDatabase{
		connection: db,
	}
}

func (db *promotionForProductDatabase) SavePromotionForProduct(promotionForProduct entity.PromotionForProduct) (int, error) {
	data := &promotionForProduct
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	err := db.connection.Create(data)
	if(err.Error != nil){
		return 0, err.Error
	}
	return data.ID, nil
}

func (db *promotionForProductDatabase) UpdatePromotionForProduct(promotionForProduct entity.PromotionForProduct) error {
	data := &promotionForProduct
	data.UpdatedAt = time.Now()
	db.connection.Save(data)
	return nil
}

func (db *promotionForProductDatabase) DeletePromotionForProduct(promotionForProduct entity.PromotionForProduct) error {
	db.connection.Delete(&promotionForProduct)
	return nil
}

func (db *promotionForProductDatabase) GetAllPromotionForProducts() []entity.PromotionForProduct {
	var promotionForProducts []entity.PromotionForProduct
	db.connection.Preload(clause.Associations).Find(&promotionForProducts)
	return promotionForProducts
}

func (db *promotionForProductDatabase) GetPromotionForProduct(id string) []entity.PromotionForProduct {
	var promotionForProduct []entity.PromotionForProduct
	db.connection.Preload(clause.Associations).Where("id = ?", id).First(&promotionForProduct)
	return promotionForProduct
}
