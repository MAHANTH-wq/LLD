package main

func main() {

	newBookProduct := newBook(54, "Zero To One", 1000)
	productAfterApplyingCouponA := applyCouponA(newBookProduct)
	productAfterApplyingCouponA.printDetails()
}
