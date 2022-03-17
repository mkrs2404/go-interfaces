package main

type Product struct {
	name      string
	unitPrice int
	onOffer   bool
	offer     ProductOffer
}

// Cart : An item in map represents a product with its quantity
type Cart struct {
	items     map[Product]int
	cartPrice int
	onOffer   bool
	offer     CartOffer
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
		var discount int
		if prod.onOffer {
			discount = prod.offer.Discount(prod, qty)
		}
		totalPrice += (prod.unitPrice * qty) - discount
	}
	c.cartPrice = totalPrice

	if c.onOffer {
		c.cartPrice = c.cartPrice - c.offer.Discount(*c)
	}
}
