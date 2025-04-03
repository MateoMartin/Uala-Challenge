package handler

type followUserDTO struct {
	UserID         string `json:"user_id" binding:"required" example:"0f089136-3f38-4757-840c-d0c954782457"`
	UserIDToFollow string `json:"user_id_to_follow" binding:"required" example:"409ebe62-52ad-4c79-b129-78d363114f8b"`
}
