package dto

type CreateTodoDTO struct {
	Text   string `json:"text" form:"text" binding:"required"`
	Status string `json:"status" form:"status" binding:"required"`
}

type UpdateTodoDTO struct {
	ID     uint64 `json:"id" form:"id" binding:"required"`
	Text   string `json:"text" form:"text" binding:"required"`
	Status string `json:"status" form:"status" binding:"required"`
}
