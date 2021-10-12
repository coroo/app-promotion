package seeder

import (
	"fmt"

	"github.com/coroo/go-starter/config"
	entity "github.com/coroo/go-starter/app/entity"
)

var promotionOfProduct = []entity.PromotionOfProduct{
	entity.PromotionOfProduct{
		PromotionOfProductCode 		: "P0001",
		PromotionID  				: 1,
	},
	entity.PromotionOfProduct{
		PromotionOfProductCode 		: "P0002",
		PromotionID  				: 1,
	},
	entity.PromotionOfProduct{
		PromotionOfProductCode 		: "P0003",
		PromotionID  				: 2,
	},
	entity.PromotionOfProduct{
		PromotionOfProductCode 		: "P0004",
		PromotionID  				: 2,
	},
	entity.PromotionOfProduct{
		PromotionOfProductCode 		: "P0005",
		PromotionID  				: 2,
	},
	entity.PromotionOfProduct{
		PromotionOfProductCode 		: "P0006",
		PromotionID  				: 2,
	},
}

func SeedPromotionOfProducts() {
	db, _ := config.ConnectDB()

	_ = db.AutoMigrate(&entity.PromotionOfProduct{})
	for i, _ := range promotionOfProduct {
		err := db.Model(&entity.PromotionOfProduct{}).Create(&promotionOfProduct[i])
		if(err.Error != nil){
			fmt.Println("cannot seed 'Payment Method' table: ", err)
		}
	}
}