package biz

import "math"

/*
@Time : 2021/8/22 13:48
@Author : Firewine
@File : money
@Software: GoLand
@Description:
*/


type Money struct {
	CurrencyCode string
	Units        int64
	Nanos        int32
}


func Carry(amount Money) Money {
	fractionSize := math.Pow(10, 9)

	amount.Nanos += int32(float64(amount.Units%1) * fractionSize)
	amount.Units = int64(math.Floor(float64(amount.Units)) + math.Floor(float64(amount.Nanos)/fractionSize))
	amount.Nanos = (amount.Nanos) % int32(fractionSize)
	return amount
}