package main

import (
	seeds "github.com/coroo/go-starter/database/seeders"
)

func main() {
	seeds.SeedPromotions()
	seeds.SeedPromotionForProducts()
	seeds.SeedPromotionOfProducts()
}