package main

import "math"

type ProductOffer interface {
	Discount(Product, int) int
}

type CartOffer interface {
	Discount(Cart) int
}

type BuyXGetYOffer struct {
	toBuy   int
	getFree int
}

type SecondItemDiscountOffer struct {
	discount int
}

type CartDiscount struct {
	minAmount int
	discount  int
}

func (d BuyXGetYOffer) Discount(p Product, qty int) int {

	chargeableQty := math.Ceil(float64(d.toBuy*qty) / float64(d.toBuy+d.getFree))
	discount := (qty - int(chargeableQty)) * p.unitPrice
	return discount
}

func (d SecondItemDiscountOffer) Discount(p Product, qty int) int {

	fullPriceQty := math.Ceil(float64(qty / 2))
	discountedQty := qty - int(fullPriceQty)
	discount := int(float64(discountedQty) * float64(p.unitPrice) * float64(d.discount) / 100)
	return discount
}

func (d CartDiscount) Discount(c Cart) int {
	var discount int
	if c.cartPrice >= d.minAmount {
		discount = int(float64(c.cartPrice) * float64(d.discount) / 100)
	}
	return discount
}
