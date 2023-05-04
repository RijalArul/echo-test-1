package webs

type LikeReviewResp struct {
	ID       uint `json:"id" form:"id"`
	ReviewID uint `json:"review_product_id" form:"review_product_id"`
	MemberID uint `json:"member_id" form:"member_id"`
}
