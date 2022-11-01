package utils

import (
	"encoding/csv"
	"log"
	"strconv"
	"strings"

	"github.com/stepann0/balance-microservice/models"
)

type Income struct {
	Name string
	Sum  float64
}

func MonthReport(month, year string) (string, error) {
	var incomes []Income
	err := models.DB.Raw(
		`SELECT services.name, SUM(payments.amount) AS sum
		FROM payments
		LEFT JOIN services ON payments.service_id=services.id
		WHERE MONTH(created_at)=? AND YEAR(created_at)=?
		GROUP BY payments.service_id`, month, year).Scan(&incomes).Error
	if err != nil {
		return "", err
	}

	return WriteCSV(incomes), nil
}

// Записывает отчет в CSV
func WriteCSV(data []Income) string {
	var buff strings.Builder
	w := csv.NewWriter(&buff)
	w.Comma = ';'

	for _, d := range data {
		if err := w.Write(StrIncome(d)); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
	return buff.String()
}

func StrIncome(i Income) []string {
	return []string{i.Name, strconv.FormatFloat(i.Sum, 'f', 3, 64)}
}
