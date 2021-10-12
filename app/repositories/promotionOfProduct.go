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

type PromotionOfProductRepository interface {
	SavePromotionOfProduct(promotionOfProduct entity.PromotionOfProduct) (int, error)
	UpdatePromotionOfProduct(promotionOfProduct entity.PromotionOfProduct) error
	DeletePromotionOfProduct(promotionOfProduct entity.PromotionOfProduct) error
	GetAllPromotionOfProducts() []entity.PromotionOfProduct
	GetPromotionOfProduct(productCode string) entity.PromotionOfProduct
}

type promotionOfProductDatabase struct {
	connection *gorm.DB
}

func NewPromotionOfProductRepository() PromotionOfProductRepository {
	db, err := config.ConnectDB()
	if err != nil {
		panic("Failed to connect database")
	}
	// db.AutoMigrate(&entity.PromotionOfProductRate{}, &entity.Person{})
	db.AutoMigrate(&entity.PromotionOfProduct{})
	return &promotionOfProductDatabase{
		connection: db,
	}
}

func (db *promotionOfProductDatabase) SavePromotionOfProduct(promotionOfProduct entity.PromotionOfProduct) (int, error) {
	data := &promotionOfProduct
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	err := db.connection.Create(data)
	if(err.Error != nil){
		return 0, err.Error
	}
	return data.ID, nil
}

func (db *promotionOfProductDatabase) UpdatePromotionOfProduct(promotionOfProduct entity.PromotionOfProduct) error {
	data := &promotionOfProduct
	data.UpdatedAt = time.Now()
	db.connection.Save(data)
	return nil
}

func (db *promotionOfProductDatabase) DeletePromotionOfProduct(promotionOfProduct entity.PromotionOfProduct) error {
	db.connection.Delete(&promotionOfProduct)
	return nil
}

func (db *promotionOfProductDatabase) GetAllPromotionOfProducts() []entity.PromotionOfProduct {
	var promotionOfProducts []entity.PromotionOfProduct
	db.connection.Preload(clause.Associations).Find(&promotionOfProducts)
	return promotionOfProducts
}

func (db *promotionOfProductDatabase) GetPromotionOfProduct(productCode string) entity.PromotionOfProduct {
	var promotionOfProduct entity.PromotionOfProduct
	db.connection.Preload("Promotion.PromotionForProduct").Preload(clause.Associations).Where("promotion_of_product_code = ?", productCode).First(&promotionOfProduct)
	return promotionOfProduct
}
