package webs

type ReviewProductResp struct {
	MemberID   uint   `json:"member_id" form:"member_id"`
	ProductID  uint   `json:"product_id" form:"product_id"`
	DescReview string `json:"desc_review" form:"desc_review"`
}

type ReviewProductDTO struct {
	DescReview string `json:"desc_review" form:"desc_review"`
}
