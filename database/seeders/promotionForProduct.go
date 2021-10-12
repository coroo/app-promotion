package seeder

import (
	"fmt"

	"github.com/coroo/go-starter/config"
	entity "github.com/coroo/go-starter/app/entity"
)

var promotionForProduct = []entity.PromotionForProduct{
	entity.PromotionForProduct{
		PromotionForProductCode 		: "P0001",
		PromotionID  					: 1,
		PromotionCapacity				: 1,
	},
	entity.PromotionForProduct{
		PromotionForProductCode 		: "P0002",
		PromotionID  					: 1,
		PromotionCapacity				: 1,
	},
	entity.PromotionForProduct{
		PromotionForProductCode 		: "P0003",
		PromotionID  					: 2,
		PromotionCapacity				: 1,
	},
	entity.PromotionForProduct{
		PromotionForProductCode 		: "P0004",
		PromotionID  					: 2,
		PromotionCapacity				: 1,
	},
	entity.PromotionForProduct{
		PromotionForProductCode 		: "P0005",
		PromotionID  					: 2,
		PromotionCapacity				: 1,
	},
	entity.PromotionForProduct{
		PromotionForProductCode 		: "P0006",
		PromotionID  					: 2,
		PromotionCapacity				: 1,
	},
}

func SeedPromotionForProducts() {
	db, _ := config.ConnectDB()

	_ = db.AutoMigrate(&entity.PromotionForProduct{})
	for i, _ := range promotionForProduct {
		err := db.Model(&entity.PromotionForProduct{}).Create(&promotionForProduct[i])
		if(err.Error != nil){
			fmt.Println("cannot seed 'Payment Method' table: ", err)
		}
	}
}