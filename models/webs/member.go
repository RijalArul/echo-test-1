package webs

type MemberResponse struct {
	ID        uint   `json:"id" form:"id"`
	Username  string `json:"username" form:"username"`
	Gender    string `json:"gender" form:"gender"`
	SkinType  string `json:"skin_type" form:"skin_type"`
	SkinColor string `json:"skin_color" form:"skin_color"`
}

type MemberDTO struct {
	Username  string `json:"username" form:"username"`
	Gender    string `json:"gender" form:"gender"`
	SkinType  string `json:"skin_type" form:"skin_type"`
	SkinColor string `json:"skin_color" form:"skin_color"`
}
