package entity

import (
	"time"
)

type PromotionForProduct struct {
	ID               		int            	`gorm:"type:BIGINT UNSIGNED NOT NULL AUTO_INCREMENT" json:"id"`
	PromotionID 			int    			`gorm:"type:BIGINT UNSIGNED NOT NULL" json:"promotion_id"`
	PromotionForProductCode	string    		`gorm:"type:varchar(191) NULL default NULL;default:null" json:"promotion_for_product_code"`
	PromotionCapacity  		int	    		`gorm:"type:INT UNSIGNED NULL" json:"promotion_capacity"`
	CreatedAt        		time.Time      	`gorm:"type:TIMESTAMP DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        		time.Time      	`gorm:"type:TIMESTAMP DEFAULT CURRENT_TIMESTAMP" json:"updated_at"`
}
