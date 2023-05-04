package services

import (
	"echo-test-1/models/entities"
	"echo-test-1/models/webs"
	"echo-test-1/repositories"
)

type ProductService interface {
	GetAllProducts() ([]webs.ProductResponse, error)
	CreateProduct(productDTO webs.ProductDTO) (webs.ProductResponse, error)
	DetailProduct(id uint) (webs.ProductReviewResponse, error)
}

type ProductServiceImpl struct {
	productRepository repositories.ProductRepository
}

func NewProductService(ProductRepo repositories.ProductRepository) ProductService {
	return &ProductServiceImpl{productRepository: ProductRepo}
}

func convertBodyProductResp(product entities.Product) webs.ProductResponse {
	return webs.ProductResponse{
		ID:          product.ID,
		NameProduct: product.NameProduct,
		Price:       product.Price,
	}
}

func convertBodyProductReviewResp(product entities.Product, reviewProduct entities.ReviewProduct, membersLike []entities.Member, countLike int) webs.ProductReviewResponse {
	return webs.ProductReviewResponse{
		ID:            product.ID,
		NameProduct:   product.NameProduct,
		Price:         product.Price,
		ReviewProduct: reviewProduct,
		Member:        membersLike,
		CountLike:     countLike,
	}
}

func (s *ProductServiceImpl) GetAllProducts() ([]webs.ProductResponse, error) {
	products, err := s.productRepository.GetAllProducts()
	productsResp := []webs.ProductResponse{}

	for i := 0; i < len(products); i++ {
		productBody := convertBodyProductResp(products[i])
		productsResp = append(productsResp, productBody)
	}

	return productsResp, err
}

func (s *ProductServiceImpl) CreateProduct(productDTO webs.ProductDTO) (webs.ProductResponse, error) {
	products := entities.Product{
		NameProduct: productDTO.NameProduct,
		Price:       productDTO.Price,
	}

	newProduct, err := s.productRepository.CreateProduct(products)
	ProductResp := convertBodyProductResp(*newProduct)
	return ProductResp, err
}

func (s *ProductServiceImpl) DetailProduct(id uint) (webs.ProductReviewResponse, error) {
	product, reviewProduct, membersLike, countLike, err := s.productRepository.DetailProduct(id)
	productResp := convertBodyProductReviewResp(*product, *reviewProduct, membersLike, countLike)
	return productResp, err
}
