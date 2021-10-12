package entity

type CalculatorPromotion struct {
	// ID               		int            	`gorm:"type:BIGINT UNSIGNED NOT NULL AUTO_INCREMENT" json:"id"`
	ProductCode  			string    		`gorm:"type:varchar(191) NULL default NULL;default:null" json:"product_code"`
	Price  					int	    		`gorm:"type:INT UNSIGNED NULL" json:"price"`
	// CreatedAt        		time.Time      	`gorm:"type:TIMESTAMP DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	// UpdatedAt        		time.Time      	`gorm:"type:TIMESTAMP DEFAULT CURRENT_TIMESTAMP" json:"updated_at"`
}
