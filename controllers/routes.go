package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stepann0/balance-microservice/models"
	"github.com/stepann0/balance-microservice/utils"
	"gorm.io/gorm"
)

// PUT /
// Создает новый и резервный аккаунты согласно с json телом запроса. ID задается автоматически
func CreateAccount(c *gin.Context) {
	var input CreateAccountInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account := models.Account{Name: input.Name, Balance: input.Balance}
	models.DB.Create(&account)
	reserved_account := models.ReservationAccount{ID: account.ID}
	models.DB.Create(&reserved_account)

	c.JSON(http.StatusOK, gin.H{"account": account})
}

// GET /:id
// Достает аккаунт из БД
func GetAccount(c *gin.Context) {
	var account models.Account
	if err := GetAccountByID(c.Param("id"), &account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"account": account})
}

// Удаляет аккаунт из БД
func DeleteAccount(c *gin.Context) {
	var account models.Account
	if err := GetAccountByID(c.Param("id"), &account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Delete(&account)
	c.Status(http.StatusOK)
}

// POST /increase/:id
// Начисляет сумму на главный счет
func IncreaseBalance(c *gin.Context) {
	// Validate input
	var input UpdateBalance
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var account models.Account
	if err := GetAccountByID(input.ID, &account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Model(&account).Update("balance", gorm.Expr("balance + ?", input.Amount))

	GetAccountByID(input.ID, &account)
	c.JSON(http.StatusOK, gin.H{"status": "OK",
		"amount":  input.Amount,
		"balance": account.Balance,
	})
}

// POST /reserve
// Резервирует сумму в отдельном аккаунте
func ReserveBalance(c *gin.Context) {
	var input ReserveBalanceInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var main_account models.Account
	if err := GetAccountByID(input.UserID, &main_account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var reserve_account models.ReservationAccount
	if err := GetAccountByID(input.UserID, &reserve_account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ReserveAmount(&main_account, &reserve_account, input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	GetAccountByID(input.UserID, &main_account)
	GetAccountByID(input.UserID, &reserve_account)
	c.JSON(http.StatusOK, gin.H{
		"user_balance":   main_account.Balance,
		"reserve_amount": reserve_account.Amount,
	})
}

// PUT /accept/:id
// Одобряет платеж. Списывает зарезервированную сумму
func AcceptPayment(c *gin.Context) {
	var input ReserveBalanceInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var main_account models.Account
	if err := GetAccountByID(input.UserID, &main_account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var reserve_account models.ReservationAccount
	if err := GetAccountByID(input.UserID, &reserve_account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ApplyPayment(&reserve_account, input)
	c.Status(http.StatusOK)
}

// Отклоняет платеж
func DeclinePayment(c *gin.Context) {
	var input ReserveBalanceInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var main_account models.Account
	if err := GetAccountByID(input.UserID, &main_account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var reserve_account models.ReservationAccount
	if err := GetAccountByID(input.UserID, &reserve_account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := UnreserveAmount(&main_account, &reserve_account, input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

// GET /report/:month/:year
// Создает отчет за :mounth.:year
func Report(c *gin.Context) {
	report, err := utils.MonthReport(c.Param("month"), c.Param("year"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	c.String(http.StatusOK, report)
}
