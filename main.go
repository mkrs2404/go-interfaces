package main

import "math"

type Offer interface{}

type BuyXGetYOffer struct {
	toBuy   int
	getFree int
}

type SecondItemDiscountOffer struct {
	discount int
}

type CartOffer struct{
	minAmount int
	discount int
}

type Product struct {
	name      string
	unitPrice int
	onOffer   bool
	offer     Offer
}

// Cart : An item in map represents a product with its quantity
type Cart struct {
	items     map[Product]int
	cartPrice int
	onOffer bool
	offer Offer
}

func main() {

}

func (c *Cart) addProduct(product Product, qty int) {
	c.items[product] = qty
	c.updateCartPrice()
}

func (c *Cart) updateCartPrice() {
	var totalPrice int
	for prod, qty := range c.items {
		if prod.onOffer {
			switch prod.offer.(type) {
			case BuyXGetYOffer:
				qty = prod.getChargeableQty(qty)
				totalPrice += prod.unitPrice * qty
			case SecondItemDiscountOffer:
				fullPriceQty := qty / 2
				discountedQty := qty - fullPriceQty
				totalPrice += int(float64(discountedQty) * float64(prod.unitPrice) * float64(prod.offer.(SecondItemDiscountOffer).discount) / 100)
				totalPrice += fullPriceQty * prod.unitPrice
			}

		} else {
			totalPrice += prod.unitPrice * qty
		}
	}
	c.cartPrice = totalPrice

	if c.onOffer{
		offer := c.offer.(CartOffer)
		if c.cartPrice >= offer.minAmount{
			c.cartPrice = int(float64(c.cartPrice) - float64(c.cartPrice) * (float64(offer.discount) / 100))
		}
	}

}

func (p Product) getChargeableQty(qty int) int {
	chargeableQty := math.Ceil(float64(p.offer.(BuyXGetYOffer).toBuy*qty) / float64(p.offer.(BuyXGetYOffer).toBuy+p.offer.(BuyXGetYOffer).getFree))
	return int(chargeableQty)
}
