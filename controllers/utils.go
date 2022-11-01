package controllers

import (
	"fmt"

	"github.com/stepann0/balance-microservice/models"
	"gorm.io/gorm"
)

// Получает аккаунт из БД по ID
func GetAccountByID(id, dest any) error {
	if err := models.DB.Where("id = ?", id).First(dest).Error; err != nil {
		return fmt.Errorf("record not found")
	}
	return nil
}

// Проверяет наличие средств на главном счете пользователя
func CheckEnoughBalance(account *models.Account, amount float64) error {
	if account.Balance < amount {
		return fmt.Errorf("unsufficient funds in the account")
	}
	return nil
}

func CheckEnoughReservation(account *models.ReservationAccount, amount float64) error {
	if account.Amount < amount {
		return fmt.Errorf("unsufficient funds in the reserve account")
	}
	return nil
}

// Списывает указанную сумму с главного аккаунта польщователя на аккаунт резервации
func ReserveAmount(main_account *models.Account, reserve_account *models.ReservationAccount, input ReserveBalanceInput) error {
	if err := CheckEnoughBalance(main_account, input.Amount); err != nil {
		return err
	}
	models.DB.Model(&main_account).Update("balance", gorm.Expr("balance - ?", input.Amount))
	models.DB.Model(&reserve_account).Update("amount", gorm.Expr("amount + ?", input.Amount))
	models.DB.Model(&reserve_account).Update("service_id", input.ServiceID)
	return nil
}

// Списывает всю сумму с аккаунта резервации обратно на главный счет пользователя
func UnreserveAmount(main_account *models.Account, reserve_account *models.ReservationAccount, input ReserveBalanceInput) error {
	if err := CheckEnoughReservation(reserve_account, input.Amount); err != nil {
		return err
	}
	models.DB.Model(&reserve_account).Update("amount", gorm.Expr("amount - ?", input.Amount))
	models.DB.Model(&main_account).Update("balance", gorm.Expr("balance + ?", input.Amount))
	models.DB.Model(&reserve_account).Update("service_id", 0)
	return nil
}

// Списывает всю сумму с аккаунта резервации в пользу компании
func ApplyPayment(reserve_account *models.ReservationAccount, input ReserveBalanceInput) error {
	if err := CheckEnoughReservation(reserve_account, input.Amount); err != nil {
		return fmt.Errorf("unsufficient funds in the reserve account")
	}
	payment := models.Payment{
		ServiceID: input.ServiceID,
		Amount:    input.Amount,
	}
	models.DB.Create(&payment)
	models.DB.Model(&reserve_account).Update("amount", gorm.Expr("amount - ?", input.Amount))
	models.DB.Model(&reserve_account).Update("service_id", 0)
	return nil
}
