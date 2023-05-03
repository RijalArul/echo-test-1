package webs

type ProductResponse struct {
	ID          uint   `json:"id" form:"id"`
	NameProduct string `json:"name_product" form:"name_product"`
	Price       int    `json:"price" form:"price"`
}

type ProductDTO struct {
	NameProduct string `json:"name_product" form:"name_product"`
	Price       int    `json:"price" form:"price"`
}
