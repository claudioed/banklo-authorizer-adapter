package domain

import "time"

type Transaction struct {
	Id         string
	Amount     float64
	LocalTime  time.Time
	CardId     string
	MerchantId string
}
