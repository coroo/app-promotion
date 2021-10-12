package seeder

import (
	"fmt"
	"github.com/coroo/go-starter/config"
	entity "github.com/coroo/go-starter/app/entity"
)

var promotion = []entity.Promotion{
	entity.Promotion{
		ID  					: 1,
		PromotionType 			: "half_price",
		Status 		 			: "active",
	},
	entity.Promotion{
		ID  					: 2,
		PromotionType 			: "cheapest_one_free",
		Status 		 			: "active",
	},
}

func SeedPromotions() {
	db, _ := config.ConnectDB()

	_ = db.AutoMigrate(&entity.Promotion{})
	for i, _ := range promotion {
		err := db.Model(&entity.Promotion{}).Create(&promotion[i])
		if(err.Error != nil){
			fmt.Println("cannot seed 'Payment Method' table: ", err)
		}
	}
}