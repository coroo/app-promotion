package entity

import (
	"time"
)

type Status string

const (
	Inactive Status = "inactive"
	Active Status = "active"
)

type Promotion struct {
	ID               		int            	`gorm:"type:BIGINT UNSIGNED NOT NULL AUTO_INCREMENT" json:"id"`
	PromotionType 			string    		`gorm:"type:varchar(191) NULL default NULL;default:null" json:"promotion_type"`
	PromotionOfProduct      []PromotionOfProduct
	PromotionForProduct     []PromotionForProduct
	Status 		 			string    		`gorm:"type:enum('active','inactive') NOT NULL DEFAULT 'inactive';default:'inactive'" json:"status"`
	CreatedAt        		time.Time      	`gorm:"type:TIMESTAMP DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        		time.Time      	`gorm:"type:TIMESTAMP DEFAULT CURRENT_TIMESTAMP" json:"updated_at"`
}
