package controllers

type CreateAccountInput struct {
	Name    string  `json:"name" binding:"required"`
	Balance float64 `json:"balance"`
}

type UpdateBalance struct {
	ID     uint    `json:"id" binding:"required"`
	Amount float64 `json:"amount" binding:"required"`
}

type ReserveBalanceInput struct {
	UserID    uint    `json:"user_id" binding:"required"`
	ServiceID uint    `json:"service_id" binding:"required"`
	Amount    float64 `json:"amount" binding:"required"`
}
