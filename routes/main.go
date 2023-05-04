package routes

import (
	"echo-test-1/databases"
	"echo-test-1/handlers"
	"echo-test-1/repositories"
	"echo-test-1/services"

	"github.com/labstack/echo/v4"
)

func Routes() *echo.Echo {
	r := echo.New()
	db := databases.GetDB()

	memberRepository := repositories.NewMemberRepository(db)
	memberService := services.NewMemberService(memberRepository)
	memberHandler := handlers.NewMemberHandler(memberService)
	memberRouter := r.Group("/members")

	{
		memberRouter.GET("", memberHandler.Members)
		memberRouter.POST("/register", memberHandler.Create)
		memberRouter.PUT("/update/:id", memberHandler.Update)
		memberRouter.DELETE("/delete/:id", memberHandler.Delete)
	}

	productRepository := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepository)
	productHandler := handlers.NewProductHandler(productService)
	productRouter := r.Group("/products")

	{
		productRouter.GET("", productHandler.Products)
		productRouter.POST("/register", productHandler.Create)
	}

	reviewRepository := repositories.NewReviewProductRepository(db)
	reviewService := services.NewReviewProductService(reviewRepository)
	reviewHandler := handlers.NewReviewProductHandler(reviewService)
	reviewRouter := r.Group("/reviews")

	{
		reviewRouter.POST("/create/:product_id", reviewHandler.Create)
		reviewRouter.GET("/:product_id", reviewHandler.Reviews)
	}

	likeReviewRepository := repositories.NewLikeReviewRepository(db)
	likeReviewService := services.NewLikeReviewService(likeReviewRepository)
	likeReviewHandler := handlers.NewLikeReviewHandler(likeReviewService)
	likeReviewRouter := r.Group("/like_reviews")

	{
		likeReviewRouter.POST("/:review_id", likeReviewHandler.Reaction)
	}

	r.Start(":9000")

	return r
}
