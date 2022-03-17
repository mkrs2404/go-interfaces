package main

import "testing"

var testData = []struct {
	testName    string
	product     []Product
	qty         []int
	onCartOffer bool
	offer       CartOffer
	finalPrice  int
}{
	{
		testName: "AddProductNoOffer",
		product: []Product{
			{
				name:      "Dove",
				unitPrice: 30,
				onOffer:   false,
			},
		},
		qty: []int{
			5,
		},
		finalPrice: 150,
	},
	{
		testName: "AddProductWithOffer",
		product: []Product{
			{
				name:      "Dove",
				unitPrice: 30,
				onOffer:   true,
				offer: BuyXGetYOffer{
					toBuy:   2,
					getFree: 1,
				},
			},
			{
				name:      "Axe Deo",
				unitPrice: 100,
				onOffer:   false,
			},
		},
		qty: []int{
			3,
			2,
		},
		finalPrice: 260,
	},
	{
		testName: "AddProductWithDiscount",
		product: []Product{
			{
				name:      "Dove",
				unitPrice: 30,
				onOffer:   true,
				offer: SecondItemDiscountOffer{
					discount: 50,
				},
			},
		},
		qty: []int{
			2,
		},
		finalPrice: 45,
	},
	{
		testName: "AddProductWithCartDiscount",
		product: []Product{
			{
				name:      "Dove",
				unitPrice: 30,
				onOffer:   false,
			},
			{
				name:      "Axe Deo",
				unitPrice: 100,
				onOffer:   false,
			},
		},
		qty: []int{
			5,
			4,
		},
		onCartOffer: true,
		offer: CartDiscount{
			minAmount: 500,
			discount:  20,
		},
		finalPrice: 440,
	},
}

func TestAddProduct(t *testing.T) {

	for _, test := range testData {
		t.Run(test.testName, func(t *testing.T) {
			var c Cart
			itemMap := make(map[Product]int)
			c.onOffer = test.onCartOffer
			c.offer = test.offer
			c.items = itemMap
			for i, prod := range test.product {
				c.addProduct(prod, test.qty[i])
			}
			if c.cartPrice != test.finalPrice {
				t.Errorf("Expected %d Got %d", test.finalPrice, c.cartPrice)
			}
		})

	}
}
