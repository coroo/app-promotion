package entity

import (
	"time"
)

type PromotionOfProduct struct {
	ID               		int            	`gorm:"type:BIGINT UNSIGNED NOT NULL AUTO_INCREMENT" json:"id"`
	PromotionID 			int    			`gorm:"type:BIGINT UNSIGNED NOT NULL" json:"promotion_id"`
	PromotionOfProductCode	string    		`gorm:"type:varchar(191) NULL default NULL;default:null" json:"promotion_of_product_code"`
	Promotion      			Promotion
	CreatedAt        		time.Time      	`gorm:"type:TIMESTAMP DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        		time.Time      	`gorm:"type:TIMESTAMP DEFAULT CURRENT_TIMESTAMP" json:"updated_at"`
}
