package repositories

import (
	"echo-test-1/models/entities"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepository interface {
	GetAllProducts() ([]entities.Product, error)
	CreateProduct(product entities.Product) (*entities.Product, error)
	DetailProduct(id uint) (*entities.Product, *entities.ReviewProduct, []entities.Member, int, error)
}

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(DB *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{db: DB}
}

func (r *ProductRepositoryImpl) GetAllProducts() ([]entities.Product, error) {
	products := []entities.Product{}
	err := r.db.Model(&products).Find(&products).Error
	return products, err
}

func (r *ProductRepositoryImpl) CreateProduct(product entities.Product) (*entities.Product, error) {
	err := r.db.Create(&product).Error

	return &product, err
}

type CountLikeReview struct {
	JumlahLikeReview int
}

func (r *ProductRepositoryImpl) DetailProduct(id uint) (*entities.Product, *entities.ReviewProduct, []entities.Member, int, error) {

	var membersLike []entities.Member
	var counts CountLikeReview
	product := entities.Product{}
	reviewProducts := entities.ReviewProduct{}
	err := r.db.Preload(clause.Associations).First(&product).Where("id = ?", id).Error
	err = r.db.Preload(clause.Associations).First(&reviewProducts).Where("product_id = ?", &product.ID).Error
	err = r.db.Raw("SELECT members.id as id, username, gender, skin_type, skin_color, desc_review FROM products RIGHT JOIN review_products ON products.id = review_products.product_id RIGHT JOIN like_reviews ON products.id = like_reviews.review_product_id RIGHT JOIN members ON like_reviews.member_id = members.id WHERE product_id = ?", product.ID).Scan(&membersLike).Error
	err = r.db.Raw("SELECT count(*) as jumlah_like_review FROM products RIGHT JOIN review_products ON products.id = review_products.product_id RIGHT JOIN like_reviews ON products.id = like_reviews.review_product_id RIGHT JOIN members ON like_reviews.member_id = members.id WHERE product_id = ?", product.ID).Scan(&counts).Error
	return &product, &reviewProducts, membersLike, counts.JumlahLikeReview, err
}
