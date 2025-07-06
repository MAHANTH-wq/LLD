package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type emailOtp struct {
}

func (email *emailOtp) genRandomOTP(otpLength int) string {

	minOtp := 1

	for i := 1; i < otpLength; i++ {
		minOtp = minOtp * 10
	}
	maxOtp := minOtp*10 - 1

	randomNumber := rand.Intn(maxOtp-minOtp+1) + minOtp
	fmt.Println("Generated OTP ", randomNumber)
	return strconv.Itoa(randomNumber)

}

func (email *emailOtp) saveOTPToCache(otpValue string) {
	fmt.Println("OTP" + otpValue + " saved to email cache")
}

func (email *emailOtp) getMessage(otpValue string) string {
	return fmt.Sprintf("Please use %s otp to login into the app", otpValue)
}

func (email *emailOtp) sendNotification(message string) error {
	fmt.Println("Sent OTP Notification Message through Email: ", message)
	return nil
}
