package main

import "fmt"

type product interface {
	getPrice() int
	printDetails()
}

type bookProduct struct {
	bookId   int
	bookName string
	price    int
}

func newBook(id int, name string, price int) *bookProduct {

	return &bookProduct{
		bookId:   id,
		bookName: name,
		price:    price,
	}
}

func (b *bookProduct) getPrice() int {
	return b.price
}

func (b *bookProduct) printDetails() {
	fmt.Println("Book Name ", b.bookName, "Book Price ", b.price)
}

type couponA struct {
	product    product
	couponCode int
	couponName string
	price      int
}

func applyCouponA(p product) *couponA {

	newPrice := p.getPrice()

	//For Products above 100 get 20 off
	if p.getPrice() > 100 {
		newPrice = newPrice - 20
	}

	return &couponA{
		product:    p,
		couponCode: 0,
		couponName: "Coupon A",
		price:      newPrice,
	}
}

func (c *couponA) getPrice() int {
	return c.price
}

func (c *couponA) printDetails() {
	c.product.printDetails()
	fmt.Println("Coupon A Discount 20 off on products above 100")
	fmt.Println("After Coupon A discount price ", c.price)
}
