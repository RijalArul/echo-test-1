package webs

import "echo-test-1/models/entities"

type ProductResponse struct {
	ID          uint   `json:"id" form:"id"`
	NameProduct string `json:"name_product" form:"name_product"`
	Price       int    `json:"price" form:"price"`
}

type ProductReviewResponse struct {
	ID            uint   `json:"id" form:"id"`
	NameProduct   string `json:"name_product" form:"name_product"`
	Price         int    `json:"price" form:"price"`
	ReviewProduct entities.ReviewProduct
	Member        []entities.Member
	CountLike     int
}

type ProductDTO struct {
	NameProduct string `json:"name_product" form:"name_product"`
	Price       int    `json:"price" form:"price"`
}
